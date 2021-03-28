package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"scratchpay.com/clinics/dto"
	"scratchpay.com/clinics/service"
	"scratchpay.com/clinics/service/impl"
	"scratchpay.com/clinics/utils/message"
	"scratchpay.com/clinics/utils/responseStatus"
	"strconv"
)

var clinicService service.ClinicService = impl.ClinicServiceImpl{}

func Search(c echo.Context) error {
	name := c.QueryParam("name")
	state := c.QueryParam("state")
	from := c.QueryParam("from")
	to := c.QueryParam("to")
	page, err := strconv.Atoi(c.QueryParam("page")); if err != nil {
		log.Println("Page not set. ",err.Error())
		page = 1
	}
	limit, err := strconv.Atoi(c.QueryParam("limit")); if err != nil {
		log.Println("Limit not set.", err.Error())
		limit = 10
	}
	log.Printf("name:%s, state: %s, from: %s, to: %s, page: %d, limit: %d", name, state, from, to, page, limit)
	response, err := clinicService.Search(name, state, from, to, page, limit)
	if err != nil {
		fmt.Println("An error occurred.", err.Error())
		return c.JSON(http.StatusOK, dto.ServiceResponse{Status: responseStatus.ERROR, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.ServiceResponse{Status: responseStatus.SUCCESS, Message: message.GENERAL_SUCCESS_MESSAGE, Data: response})
}
