package impl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"scratchpay.com/clinics/dto"
	"time"
)

type DentalClinicSearchService struct {}

func (dental DentalClinicSearchService) search(clinicName, state, from, to string) (dto.ClinicDto, error) {
	url := os.Getenv("DENTAL_CLINIC_URL")
	httpClient := http.Client{Timeout: 5 * time.Second}
	res, err := httpClient.Get(url)
	if err != nil {
		log.Panic("An error occurred.", err.Error())
		return dto.ClinicDto{}, err
	}
	if res.StatusCode >= 200 && res.StatusCode < 300 {
		var data map[string] string
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, &data)
		fmt.Println(data);
	} else {
		return dto.ClinicDto{}, errors.New("An error occurred.")
	}

	return dto.ClinicDto{}, nil
}
