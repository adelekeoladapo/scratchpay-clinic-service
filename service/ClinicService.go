package service

import "scratchpay.com/clinics/dto"

type ClinicService interface {

	search(clinicName, state, from, to string) (dto.ClinicDto, error)

}