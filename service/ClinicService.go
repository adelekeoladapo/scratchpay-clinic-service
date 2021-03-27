package service

import (
	"scratchpay.com/clinics/dto"
)

type ClinicService interface {

	Search(clinicName, state, from, to string) ([]dto.ClinicDto, error)

}