package impl

import "scratchpay.com/clinics/dto"

type VetClinicSearchService struct {}

func (vet VetClinicSearchService) search(clinicName, state, from, to string) (dto.ClinicDto, error) {
	return dto.ClinicDto{}, nil
}