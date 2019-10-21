package main

import (
	"assignment1v4/GBIF_APis"
	"fmt"
	"log"
	"net/http"
	"os"
)

func handlerRequests() {
	http.HandleFunc("/conservation/v1/country/", GBIF_APis.CountryHandler)
	http.HandleFunc("/conservation/v1/species/", GBIF_APis.HandlerSpecies)
	http.HandleFunc("/conservation/v1/diag/", GBIF_APis.HandlerDiag)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {

	fmt.Println("heey")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	handlerRequests()

}
