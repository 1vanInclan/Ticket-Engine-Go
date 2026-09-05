package controller

import (
	"net/http"
	"strconv"
	"ticket-engine/usecase/dto"
	reservationInt "ticket-engine/usecase/interactor/reservation"

	"github.com/labstack/echo/v4"
)

type ReservationController struct {
	reservationInteractor reservationInt.ReservationInteractor
}

var ReservationCtrl = &ReservationController{
	reservationInteractor: reservationInt.ReservationInt,
}

func (c *ReservationController) Create(ctx echo.Context) error {
	userID, ok := ctx.Get("userID").(uint)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	var req dto.CreateReservationInput
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "payload invalido"})
	}

	res, err := c.reservationInteractor.Create(ctx.Request().Context(), userID, req)
	if err != nil {
		if err.Error() == "insufficient stock for this event" {
			return ctx.JSON(http.StatusConflict, map[string]string{"error": err.Error()})
		}
		if err.Error() == "event not found" {
			return ctx.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (c *ReservationController) FindByUserID(ctx echo.Context) error {
	userID, ok := ctx.Get("userID").(uint)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	res, err := c.reservationInteractor.FindByUserID(ctx.Request().Context(), userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *ReservationController) FindByID(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "id de reservacion invalido"})
	}

	res, err := c.reservationInteractor.FindByID(ctx.Request().Context(), uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if res == nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "reservacion no encontrada"})
	}

	return ctx.JSON(http.StatusOK, res)
}
