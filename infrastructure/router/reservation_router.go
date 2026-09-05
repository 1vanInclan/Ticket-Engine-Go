package router

import (
	resCtrl "ticket-engine/interface/controller"

	"github.com/labstack/echo/v4"
)

func ReservationRouter(g *echo.Group, ctrl *resCtrl.ReservationController) {
	group := g.Group("/reservation")

	group.POST("", ctrl.Create)
	group.GET("", ctrl.FindByUserID)
	group.GET("/:id", ctrl.FindByID)
}
