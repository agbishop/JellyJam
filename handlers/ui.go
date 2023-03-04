package handlers

import (
	"embed"
	"errors"
	"fmt"
	"github.com/labstack/echo/v5"
	"io"
	"mime"
	"net/http"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type UIHandler struct {
	fs *embed.FS
	http.Handler
}

func readFromUIFS(fs *embed.FS, prefix, requestedPath string, w http.ResponseWriter) error {
	f, err := fs.Open(path.Join(prefix, requestedPath))
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	stat, _ := f.Stat()
	if stat.IsDir() {
		return errors.New("no files in path")
	}

	contentType := mime.TypeByExtension(filepath.Ext(requestedPath))
	cacheTime := 5 * time.Minute

	if strings.Contains(contentType, "image/") {
		cacheTime = time.Hour
	}
	if strings.Contains(contentType, "html") {
		cacheTime = 0
	}

	w.Header().Add("Cache-Control", fmt.Sprintf("max-age=%.0f", cacheTime.Seconds()))
	w.Header().Add(echo.HeaderContentType, contentType)
	_, err = io.Copy(w, f)

	return err
}

const uiExportPath = "client/out"

func (ui UIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := readFromUIFS(ui.fs, uiExportPath, r.URL.Path, w); err == nil {
		return
	}
	if err := readFromUIFS(ui.fs, uiExportPath, r.URL.Path+".html", w); err == nil {
		return
	}
	if err := readFromUIFS(ui.fs, uiExportPath, "index.html", w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

}
