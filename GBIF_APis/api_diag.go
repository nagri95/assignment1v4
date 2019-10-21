package GBIF_APis

import (
	"encoding/json"
	"fmt"
	"strconv"

	"net/http"

	"time"
)

type DiagAnswer struct {
	Gbif          string  `json:"gbif"`          //: "<http status code for GBIF API>",
	Restcountries string  `json:"restcountries"` //: "<http status code for restcountries API>"
	Version       string  `json:"version"`
	Uptime        float64 `json:"uptime"` //: <time in seconds from the last service restart>
}

var startTime = time.Now()

var diagAnswer = DiagAnswer{}

func HandlerDiag(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")

	// get the restcountries json

	respRCountries, err := http.Get("https://restcountries.eu/rest/v2")
	defer respRCountries.Body.Close()

	diagAnswer.Restcountries = strconv.Itoa(respRCountries.StatusCode)

	respGBIF, err := http.Get("http://api.gbif.org/v1/species")
	defer respGBIF.Body.Close()

	diagAnswer.Gbif = strconv.Itoa(respGBIF.StatusCode)

	diagAnswer.Version = "v1"

	diagAnswer.Uptime = time.Since(startTime).Seconds()

	json.NewEncoder(w).Encode(diagAnswer)
	if err != nil {
		fmt.Println("error : we have a problem", err)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
