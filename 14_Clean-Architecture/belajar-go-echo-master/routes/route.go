package routes

import (
	"belajar-go-echo/config"
	"belajar-go-echo/controller"
	"belajar-go-echo/konstan"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = config.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	app := echo.New()
	appJWT := app.Group("")
	appJWT.Use(mid.JWT([]byte(konstan.SECRET_JWT)))
	appJWT.GET("/users", controller.GetAllUsers(db))
	appJWT.POST("/users", controller.CreateUser(db))

	return app
}
