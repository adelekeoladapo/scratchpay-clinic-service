package service

import "scratchpay.com/clinics/dto"

type ClinicSearchService interface {

	search(clinicName, state, from, to string) (dto.ClinicDto, error)

}