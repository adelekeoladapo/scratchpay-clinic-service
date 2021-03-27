package impl

import "scratchpay.com/clinics/dto"

type DentalClinicService struct {}

func (dental DentalClinicService) search(clinicName, state, from, to string) (dto.ClinicDto, error) {
	return dto.ClinicDto{}, nil
}
