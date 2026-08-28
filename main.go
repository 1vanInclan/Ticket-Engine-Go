package main

import (
	"ticket-engine/infrastructure/cache"
	"ticket-engine/infrastructure/datastore"
	"ticket-engine/infrastructure/router"
	"ticket-engine/interface/controller"
	"ticket-engine/interface/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Services
	db := datastore.NewDB()
	_ = cache.NewRedisClient()

	// Inicializar repositorios
	repository.Init(db)

	// Inicializar controllers
	appController := controller.AppController{
		Auth: controller.AuthCtrl,
	}

	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	router.Router(e, appController)

	e.Logger.Fatal(e.Start((":8080")))

}
