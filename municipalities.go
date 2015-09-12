package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

var gMunicipalities Municipalities

//Load gMunicipalities with info from Municipal Court Website File

func getMunicipalityData() ([][]string, error) {
	csvfile, err := os.Open("municipal_courts.csv")

	if err != nil {
		return nil, fmt.Errorf("Failed to Open Municipal Court Information")
	}

	defer csvfile.Close()
	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = 13

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	return rawCSVdata, nil

}
func init() {
	rawCSVdata, err := getMunicipalityData()
	if err != nil {
		panic(err)
	}

	gMunicipalities = make(Municipalities, len(rawCSVdata))
	var errString = ""
	for i, line := range rawCSVdata {
		//skip first line header
		if i == 0 {
			continue
		}
		m := Municipality{}
		m.Id = i
		m.Name = line[0]                    // Municipality
		switch strings.TrimSpace(line[1]) { // Municipal Court
		case "X":
			m.MunicipalCourt = HAS_MUNICIPAL_COURT
		case "Normandy":
			m.MunicipalCourt = NORMANDY_MUNICIPAL_COURT
		case "NONE":
			m.MunicipalCourt = NO_MUNICIPAL_COURT
		case "NONE (StLCo)":
			m.MunicipalCourt = NO_MUNICIPAL_COURT_STLCO
		case "Frontenac":
			m.MunicipalCourt = FRONTENAC_MUNICIPAL_COURT
		default:
			errString += fmt.Sprintf("Invalid Mancipal Court in line %d", i)
		}
		m.MunicipalWebsite = line[2]                                     //Municipal Website
		m.MunicipalCourtWebsite = line[3]                                //Municipal Court Website
		m.CourtDocketListed = strings.TrimSpace(line[4]) != "not listed" //Court Docket
		//m.CourtDocket TODO: Find way to retreive actual court docket
		m.CourtClerkPhoneListed = strings.TrimSpace(line[5]) != "not listed" //Court Clerk Phone #
		m.CourtClerkPhone = line[5]
		m.InfoSource = line[6]
		m.OnlinePaymentProvider = line[7]
		m.HasDressCode = strings.TrimSpace(line[8]) == "yes"
		m.FineScheduleListed = strings.TrimSpace(line[9]) == "yes"
		m.MunicipalCodes = line[10]
		m.OrderSeen = line[11]
		m.DirectionsListed = strings.TrimSpace(line[12]) != "no"

		//skipped a line at the beginning
		gMunicipalities[i-1] = m
	}

	log.Printf(errString)

}
