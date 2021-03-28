package impl

import (
	"fmt"
	"scratchpay.com/clinics/dto"
	"scratchpay.com/clinics/service"
	"scratchpay.com/clinics/utils"
	"scratchpay.com/clinics/utils/message"
	"sort"
	"strings"
)

type ClinicServiceImpl struct {}

var dental service.ClinicListService = DentalClinicListService{}
var vet service.ClinicListService = VetClinicListService{}

func (clinicServiceImpl ClinicServiceImpl) Search(name, state, from, to string, page, limit int) (dto.ListResponse, error)  {
	ch := make(chan []dto.ClinicDto, 0)
	go initClinicList(dental.GetClinicList, ch)
	go initClinicList(vet.GetClinicList, ch)
	list := append(<-ch, <-ch...)
	return sortFilterAndPaginate(list, name, state, from, to, page, limit), nil
}

func initClinicList(f func() ([]dto.ClinicDto, error), c chan []dto.ClinicDto) {
	if list, err := f(); err != nil {
		fmt.Println(message.GENERAL_ERROR_MESSAGE, err.Error())
	} else {
		c <- list
	}
}

func sortFilterAndPaginate(list []dto.ClinicDto, name, stateCode, from, to string, page, limit int) dto.ListResponse {
	response := dto.ListResponse{Page: page, Limit: limit}
	sort.Sort(dto.ByName(list))
	filtered := make([]dto.ClinicDto, 0)
	for _, value := range list {
		if (strings.Contains(strings.ToLower(value.Name), strings.ToLower(name))) && (stateCode == "" || strings.EqualFold(utils.GetStateNameFromCode(stateCode), value.State)) && ((from == "" || from >= value.Availability.From) && (from == "" || from <= value.Availability.To)) && (to == "" || to <= value.Availability.To) {
			filtered = append(filtered, dto.ClinicDto{Name: value.Name, State: value.State, Availability: value.Availability})
		}
	}
	response.Total = len(filtered)
	if limit > response.Total { limit = response.Total }
	filtered = filtered[(page-1)*limit:(page*limit)]
	response.Data = filtered
	return response
}