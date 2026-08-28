package router

import (
	"ticket-engine/interface/controller"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo, appController controller.AppController) {
	// 1. Grupo Público
	public := e.Group("/public")
	AuthRouter(public, appController.Auth)

	// 2. Grupo Privado (Protegido con JWT)
	// private := e.Group("/private")
	// private.Use(middleware.AuthMiddleware())

	// Aquí irán los módulos protegidos
	// EventRouter(private, appController.Event)
	// ReservationRouter(private, appController.Reservation)
}
