package router

import (
	authCtrl "ticket-engine/interface/controller"

	"github.com/labstack/echo/v4"
)

func AuthRouter(g *echo.Group, ctrl *authCtrl.AuthController) {
	group := g.Group("/auth")

	group.POST("/register", ctrl.Register)
	group.POST("/login", ctrl.Login)
}
