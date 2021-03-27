package impl

import (
	"fmt"
	"scratchpay.com/clinics/dto"
)

type ClinicServiceImpl struct {}

func (clinicServiceImpl ClinicServiceImpl) Search(clinicName, state, from, to string) ([]dto.ClinicDto, error)  {

	fmt.Printf("Clinic name: %s", clinicName)

	return []dto.ClinicDto{}, nil
}