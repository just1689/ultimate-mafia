package api

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/game/${id}", getGameHandler)
	e.POST("/game", newGameHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
