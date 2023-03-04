package handlers

import (
	"github.com/agbishop/JellyJam/pkgs/jellyfish/models"
	"github.com/labstack/echo/v5"
	"net/http"
)

func (svc *Service) ListPatterns(c echo.Context) error {
	res, err := svc.jf.PatternFileList()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, PatternCategories(res))
}

func PatternCategories(l *models.PatterFileList) map[string][]*models.PatterFileListInfo {
	cats := map[string][]*models.PatterFileListInfo{}
	for _, pat := range l.PatternFileList {
		pat := pat
		if pat.Name == "" {
			continue // empty is a cat name
		}
		items, ok := cats[pat.Folders]
		if !ok {
			cats[pat.Folders] = []*models.PatterFileListInfo{&pat}
		} else {
			cats[pat.Folders] = append(items, &pat)
		}
	}
	return cats
}

func (svc *Service) PatternData(c echo.Context) error {
	cat := c.PathParam("category")
	name := c.PathParam("name")
	res, err := svc.jf.PatternFile(cat, name)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
