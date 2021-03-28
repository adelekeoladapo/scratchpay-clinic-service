package service

import (
	"scratchpay.com/clinics/dto"
)

type ClinicService interface {

	Search(clinicName, state, from, to string, page, limit int) ([]dto.ClinicDto, error)

}