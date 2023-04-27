package api

import (
	"github.com/labstack/echo/v4"
)

func getGameHandler(c echo.Context) error {
	_ = c.String(200, "hi")
	return nil
}
func newGameHandler(c echo.Context) error {
	_ = c.String(200, "hi")
	return nil
}
