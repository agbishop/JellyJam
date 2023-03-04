package handlers

import (
	"embed"
	"github.com/agbishop/JellyJam/pkgs/jellyfish"
	"github.com/labstack/echo/v5"
	"log"
)

type (
	Service struct {
		echo *echo.Echo
		jf   *jellyfish.Client
		ui   *embed.FS
	}
)

func (svc *Service) Start() {
	svc.Routes()
	if err := svc.echo.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}

func (svc *Service) Routes() {
	svc.echo.GET("/api/pattern", svc.ListPatterns)
	svc.echo.GET("/api/pattern/:category/:name", svc.PatternData)
	svc.echo.GET("/api/zones", svc.Zones)
	h := UIHandler{fs: svc.ui}
	svc.echo.GET("/*", echo.WrapHandler(h))
}

func NewService(url string, ui *embed.FS) (*Service, func()) {
	jf, cleanupJF, err := jellyfish.New(url)
	if err != nil {
		log.Fatal(err)
	}
	svc := Service{
		echo: echo.New(),
		jf:   jf,
		ui:   ui,
	}

	return &svc, func() {
		cleanupJF()
	}
}
