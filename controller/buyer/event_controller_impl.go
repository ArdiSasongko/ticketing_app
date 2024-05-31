package buyer

import (
	"net/http"

	"github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/model"
	"github.com/ArdiSasongko/ticketing_app/service/buyer"
	"github.com/labstack/echo"
)

type EventControllerImpl struct {
	eventService buyer.EventService
}

func NewEventController(service buyer.EventService) *EventControllerImpl {
	return &EventControllerImpl{
		eventService: service,
	}
}

func (controller *EventControllerImpl) GetEventList(c echo.Context) error {
	filters, sort, limit, page := helper.ExtractFilterSort(c.QueryParams())
	events, err := controller.eventService.GetEventList(filters, sort, limit, page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "success", events))
}