package router

import (
	"golang-api/api/handler"

	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo) {
    api := e.Group("/api")

    api.POST("/items", handler.CreateItem)
    api.GET("/items/:id", handler.GetItem)
    api.PUT("/items/:id", handler.UpdateItem)
    api.DELETE("/items/:id", handler.DeleteItem)

    api.GET("/hello", handler.Greeting)
}