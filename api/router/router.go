package router

import (
	"golang-api/api/handler"

	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo) {
    api := e.Group("/api")

    api.POST("/user", handler.CreateItem)
    api.GET("/user/:id", handler.GetItem)
    api.PUT("/user/:id", handler.UpdateItem)
    api.DELETE("/user/:id", handler.DeleteItem)

	// Login
	api.POST("/login", handler.Login)
	// [TODO] Singup
	

    api.GET("/hello", handler.Greeting)
}