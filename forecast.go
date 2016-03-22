package forecast

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Data struct {
	Timezone string `json:"timezone"`
	Current struct {
		Summary 		string `json:"summary"`
		Icon 				string `json:"icon"`
		Temperature float64 `json:"temperature"`
		DewPoint 		float64 `json:"dewPoint"`
		Humidity		float64 `json:"humidity"`
		WindSpeed 	float64 `json:"windSpeed"`
	} `json:"currently"`
}

func Get(api string, latitude float64, longitude float64) (*Data, error) {
	url := endpoint(city)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}

	var data Data
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func endpoint(api string, latitude float64, longitude float64) string {
	baseURI := "http://api.forecast.io"
	return fmt.Sprintf("%s/%s/%s,%s", baseURI, apiKey, latitude, longitude)
}
