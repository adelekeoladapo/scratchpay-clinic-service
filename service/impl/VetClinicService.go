package impl

import "scratchpay.com/clinics/dto"

type VetClinicService struct {}

func (vet VetClinicService) search(clinicName, state, from, to string) (dto.ClinicDto, error) {
	return dto.ClinicDto{}, nil
}