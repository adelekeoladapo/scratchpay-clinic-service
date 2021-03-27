package dto

type DentalClinicClinicResponse struct {
	Name			string				`json:"name"`
	State			string				`json:"stateName"`
	Availability	AvailabilityDto		`json:"availability"`
}



