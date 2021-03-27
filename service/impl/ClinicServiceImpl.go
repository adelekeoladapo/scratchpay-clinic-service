package impl

import (
	"fmt"
	"scratchpay.com/clinics/dto"
	"scratchpay.com/clinics/service"
	"scratchpay.com/clinics/utils"
	"scratchpay.com/clinics/utils/message"
	"strings"
)

type ClinicServiceImpl struct {}

var dental service.ClinicSearchService = DentalClinicSearchService{}
var vet service.ClinicSearchService = VetClinicSearchService{}

func (clinicServiceImpl ClinicServiceImpl) Search(name, state, from, to string) ([]dto.ClinicDto, error)  {
	list := make([]dto.ClinicDto, 0)
	fmt.Printf("GetClinicList for name:%s, state: %s, from: %s, to: %s", name, state, from, to)
	fmt.Println()

	if dentalClinicList, err := dental.GetClinicList(); err != nil {
		fmt.Println(message.GENERAL_ERROR_MESSAGE, err.Error())
	} else {
		list = append(list, dentalClinicList...)
	}
	if vetClinicList, err := vet.GetClinicList(); err != nil {
		fmt.Println(message.GENERAL_ERROR_MESSAGE, err.Error())
	} else {
		list = append(list, vetClinicList...)
	}
	return filter(list, name, state, from, to), nil
}



func filter(list []dto.ClinicDto, name, stateCode, from, to string) []dto.ClinicDto {
	filtered := make([]dto.ClinicDto, 0)
	for _, value := range list {
		if (strings.Contains(strings.ToLower(value.Name), strings.ToLower(name))) && (stateCode == "" || strings.EqualFold(stateCode, value.State) || strings.EqualFold(utils.GetStateNameFromCode(stateCode), value.State)) && ((from == "" || from >= value.Availability.From) && (from == "" || from <= value.Availability.To)) && (to == "" || to <= value.Availability.To) {
			filtered = append(filtered, dto.ClinicDto{Name: value.Name, State: value.State, Availability: value.Availability})
		}
	}
	return filtered
}