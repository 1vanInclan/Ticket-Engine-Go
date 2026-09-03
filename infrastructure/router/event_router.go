package router

import (
	eventCtrl "ticket-engine/interface/controller"

	"github.com/labstack/echo/v4"
)

func EventRouter(public *echo.Group, private *echo.Group, ctrl *eventCtrl.EventController) {

	pubGroup := public.Group("/event")
	pubGroup.GET("", ctrl.FindAll)
	pubGroup.GET("/:id", ctrl.FindByID)

	privGroup := private.Group("/event")
	privGroup.POST("", ctrl.Create)

}
