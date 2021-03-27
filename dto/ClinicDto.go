package dto

type ClinicDto struct {
	Name			string			`json:"name"`
	State 			string 			`json:"state"`
	Availability	AvailabilityDto	`json:"availability"`
}

