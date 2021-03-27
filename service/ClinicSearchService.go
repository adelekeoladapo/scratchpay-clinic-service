package service

import "scratchpay.com/clinics/dto"

type ClinicSearchService interface {

	GetClinicList() ([]dto.ClinicDto, error)

}