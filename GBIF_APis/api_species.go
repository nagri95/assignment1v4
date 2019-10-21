package GBIF_APis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type SpeciesAnswer struct {
	Key            string `json:"key"`            //"<species key>",
	Kingdom        string `json:"kingdom"`        //"<kingdom>",
	Phylum         string `json:"phylum"`         //"<phylum>",
	Order          string `json:"order"`          //"<order>",
	Family         string `json:"family"`         //"<family>",
	Genus          string `json:"genus"`          //"<genus>",
	ScientificName string `json:"scientificName"` //"<scientific name>",
	CanonicalName  string `json:"canonicalName"`  //"<canonical name>",
	Year           string `json:"year"`           //<four-letter year>"
}

var speciesAnswer = SpeciesAnswer{}

func HandlerSpecies(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) != 5 || parts[3] != "species" {
		http.Error(w, "Malformed URL", http.StatusBadRequest)
		return
	}

	// get the species json
	var getArgument = fmt.Sprintf("http://api.gbif.org/v1/species/%s", parts[4])

	resp, err := http.Get(getArgument)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var speciesJsonResponse interface{}

	err = json.Unmarshal(body, &speciesJsonResponse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var MapOfSpecies = speciesJsonResponse.(map[string]interface{})

	// put country name and flag in our answer

	speciesAnswer.Key = parts[4]

	if kingdom, checker := MapOfSpecies["kingdom"].(string); checker {
		speciesAnswer.Kingdom = kingdom
	}

	if phylum, checker := MapOfSpecies["phylum"].(string); checker {
		speciesAnswer.Phylum = phylum
	}

	if order, checker := MapOfSpecies["order"].(string); checker {
		speciesAnswer.Order = order
	}

	if family, checker := MapOfSpecies["family"].(string); checker {
		speciesAnswer.Family = family
	}

	if genus, checker := MapOfSpecies["genus"].(string); checker {
		speciesAnswer.Genus = genus
	}

	if scientificName, checker := MapOfSpecies["scientificName"].(string); checker {
		speciesAnswer.ScientificName = scientificName
	}

	if canonicalName, checker := MapOfSpecies["canonicalName"].(string); checker {
		speciesAnswer.CanonicalName = canonicalName
	}
	yearS := strings.Split(MapOfSpecies["lastInterpreted"].(string), "-")

	speciesAnswer.Year = yearS[0]

	json.NewEncoder(w).Encode(speciesAnswer)
	if err != nil {
		fmt.Println("error : we have a problem", err)
	}
}
