package dto

type ClinicDto struct {
	Name			string			`json:"name"`
	State 			string 			`json:"state"`
	Availability	Availability	`son:"availability"`
}

type Availability struct {
	From 		string 			`json:"from"`
	To 			string 			`json:"to"`
}

