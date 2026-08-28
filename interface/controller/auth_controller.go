package controller

import (
	"net/http"
	"ticket-engine/usecase/dto"
	authInt "ticket-engine/usecase/interactor/auth"

	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthController struct {
	authInteractor authInt.AuthInteractor
}

var AuthCtrl = &AuthController{
	authInteractor: authInt.AuthInt,
}

func (c *AuthController) Register(ctx echo.Context) error {
	var req RegisterRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "payload inválido"})
	}

	if req.Email == "" || req.Password == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "email y contraseña requeridos"})
	}

	input := dto.RegisterInput{
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := c.authInteractor.Register(ctx.Request().Context(), input)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (c *AuthController) Login(ctx echo.Context) error {
	var req LoginRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "payload inválido"})
	}

	if req.Email == "" || req.Password == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "email y contraseña requeridos"})
	}

	input := dto.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := c.authInteractor.Login(ctx.Request().Context(), input)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, res)
}
