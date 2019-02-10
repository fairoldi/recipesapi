package main

import (
	"github.com/fairoldi/recipesapi/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	url = "localhost:32769"
)


func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	router.InitRoutes(e, url)

	e.Logger.Fatal(e.Start(":8080"))
}
