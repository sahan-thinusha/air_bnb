package api

import (
	"air_bnb/pkg/controller"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetSingleReview(c echo.Context) error {

	data, err := controller.GetSingleAirBNB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)

}
