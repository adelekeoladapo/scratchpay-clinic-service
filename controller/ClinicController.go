package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"scratchpay.com/clinics/dto"
	"scratchpay.com/clinics/service"
	"scratchpay.com/clinics/service/impl"
	"scratchpay.com/clinics/utils/message"
	"scratchpay.com/clinics/utils/responseStatus"
)

var clinicService service.ClinicService = impl.ClinicServiceImpl{}

func Search(c echo.Context) error {
	name := c.QueryParam("name")
	state := c.QueryParam("state")
	from := c.QueryParam("from")
	to := c.QueryParam("to")
	response, err := clinicService.Search(name, state, from, to)
	if err != nil {
		fmt.Println("An error occurred.", err.Error())
		return c.JSON(http.StatusOK, dto.ServiceResponse{Status: responseStatus.ERROR, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.ServiceResponse{Status: responseStatus.SUCCESS, Message: message.GENERAL_SUCCESS_MESSAGE, Data: response})
}
