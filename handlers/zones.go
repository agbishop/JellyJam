package handlers

import (
	"github.com/labstack/echo/v5"
	"net/http"
)

func (svc *Service) Zones(c echo.Context) error {
	res, err := svc.jf.Zones()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
