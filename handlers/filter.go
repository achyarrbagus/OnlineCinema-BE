package handlers

import (
	dto "backEnd/dto/result"
	"backEnd/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handlerFilter struct {
	FilterRepository repositories.FilterRepository
}

func HandlerFilter(FilterRepository repositories.FilterRepository) *handlerFilter {
	return &handlerFilter{FilterRepository}
}

func (h *handlerFilter) MultiFilter(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")

	params := c.Request().URL.Query()

	films, err := h.FilterRepository.MultiFilter(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := dto.SuccessResult{Code: http.StatusOK, Data: films}
	return c.JSON(http.StatusOK, response)
}
