package router

import (
	"ticket-engine/interface/controller"
	"ticket-engine/interface/middleware"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo, appController controller.AppController) {
	// 1. Grupo Público
	public := e.Group("/public")
	AuthRouter(public, appController.Auth)

	// 2. Grupo Privado (Protegido con JWT)
	private := e.Group("/private")
	private.Use(middleware.AuthMiddleware())

	// 3. Módulos compartidos entre público y privado
	EventRouter(public, private, appController.Event)

	// 4. Módulos 100% protegidos
	// ReservationRouter(private, appController.Reservation)
}
