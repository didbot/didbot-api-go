package main

import 	(
	"github.com/labstack/echo"
	"github.com/didbot/didbot-api-go/controllers"
)

func LoadRoutes(e *echo.Echo) {
	// Did routes
	e.GET("/dids", controllers.GetDids)
	e.GET("/dids/:id", controllers.ShowDid)
	e.POST("/dids", controllers.CreateDid)
	e.DELETE("/dids/:id", controllers.DeleteDid)
}
