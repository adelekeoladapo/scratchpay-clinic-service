package dto

type ClinicDto struct {
	Name			string			`json:"name"`
	State 			string 			`json:"state"`
	Availability	AvailabilityDto	`json:"availability"`
}

type ByName []ClinicDto

func (a ByName) Len() int {
	return len(a)
}

func (a ByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
