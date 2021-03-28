package impl

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"scratchpay.com/clinics/dto"
	"time"
)

type DentalClinicListService struct {}

func (dental DentalClinicListService) GetClinicList() ([]dto.ClinicDto, error) {
	list := make([]dto.ClinicDto, 0)
	url := os.Getenv("DENTAL_CLINIC_URL")
	httpClient := http.Client{Timeout: 5 * time.Second}
	res, err := httpClient.Get(url)
	if err != nil {
		log.Panic("An error occurred.", err.Error())
		return []dto.ClinicDto{}, err
	}
	if res.StatusCode >= 200 && res.StatusCode < 300 {
		var data []dto.DentalClinicClinicResponse
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, &data)
		for _, v := range data {
			list = append(list, dto.ClinicDto{Name: v.Name, State: v.State, Availability: v.Availability})
		}
		return list, nil
	} else {
		return []dto.ClinicDto{}, errors.New("An error occurred.")
	}
	return []dto.ClinicDto{}, nil
}
