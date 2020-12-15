package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/tarb/bfapi"
)

var (
	// login vars
	username = ""
	password = ""
	appkey   = ""

	// market catalogue filters
	maxResults  = 50
	countries   = []string{"AU"}  // GB, US, NZ
	eventTypes  = []string{"7"}   // 7=horse 4339=greyhound
	marketTypes = []string{"WIN"} // WIN, PLACE, LINE
	projection  = []string{"MARKET_START_TIME", "MARKET_DESCRIPTION", "EVENT", "EVENT_TYPE", "RUNNER_DESCRIPTION", "RUNNER_METADATA"}
)

func main() {

	// create a client
	bfClient := bfapi.NewClient(appkey, nil, nil)

	// init and login to bf
	if _, err := bfClient.Login(username, password); err != nil {
		log.Fatalf("Failed to login: %v", err)
	}

	// request a catalogue of markets meeting the specified filter/projection
	mcat, err := bfClient.ListMarketCatalogue(bfapi.ListMarketCatalogueArg{
		Filter: bfapi.MarketListFilter{
			Countries:  countries,
			EventTypes: eventTypes,
			TypeCodes:  marketTypes,
		},
		MaxResults:       maxResults,
		MarketProjection: projection,
	})
	if err != nil {
		log.Fatalf("Could not retrieve catalogue: %v", err)
	}

	// range over the catalogue and print basic details in form
	// <num> <MarketID> <time as HH:MM> <Venue>-<CountryCode> <MarketName> (<EventTypeName>)
	// 42. 1.143886544 11:52 Cranbourne-AU        R10 311m Gr5    (Greyhound Racing)
	for _, c := range mcat {
		bs, err := json.Marshal(c)
		if err == nil {
			fmt.Println(string(bs))
		}
	}
}
