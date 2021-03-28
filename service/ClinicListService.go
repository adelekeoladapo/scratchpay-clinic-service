package service

import "scratchpay.com/clinics/dto"

type ClinicListService interface {

	GetClinicList() ([]dto.ClinicDto, error)

}