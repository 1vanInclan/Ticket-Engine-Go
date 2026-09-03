package controller

import (
	"net/http"
	"strconv"
	"ticket-engine/usecase/dto"
	eventInt "ticket-engine/usecase/interactor/event"

	"github.com/labstack/echo/v4"
)

type EventController struct {
	eventInteractor eventInt.EventInteractor
}

var EventCtrl = &EventController{
	eventInteractor: eventInt.EventInt,
}

func (c *EventController) Create(ctx echo.Context) error {
	var req dto.CreateEventInput

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "payload invalido"})
	}

	if req.Name == "" || req.TotalCapacity <= 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "nombre y capacidad total son requeridos"})
	}

	res, err := c.eventInteractor.Create(ctx.Request().Context(), req)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, res)

}

func (c *EventController) FindByID(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID de evento invalido"})
	}

	res, err := c.eventInteractor.FindByID(ctx.Request().Context(), uint(id))
	if err != nil {
		if err.Error() == "event not found" {
			return ctx.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *EventController) FindAll(ctx echo.Context) error {
	res, err := c.eventInteractor.FindAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, res)
}
