package dto

type VetClinicClinicResponse struct {
	Name			string				`json:"clinicName"`
	State			string				`json:"stateCode"`
	Availability	AvailabilityDto		`json:"opening"`
}



