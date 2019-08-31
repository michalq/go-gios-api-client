package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/michalq/go-gios-api-client/client"
	"github.com/michalq/go-gios-api-client/client/stations"
)

func main() {
	transport := &client.TransportConfig{
		Host:     "api.gios.gov.pl",
		BasePath: "/pjp-api/rest",
		Schemes:  []string{"http"},
	}

	apiClient := client.NewHTTPClientWithConfig(strfmt.NewFormats(), transport)
	params := &stations.AllStationsParams{}
	params.WithTimeout(5 * time.Second)
	allStations, err := apiClient.Stations.AllStations(params)
	if err != nil {
		log.Fatal(err)
	}

	for _, station := range allStations.Payload {
		if station.City != nil {
			fmt.Println(station.City.Name)
		}
	}
}
