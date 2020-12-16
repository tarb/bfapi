package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/tarb/bfapi"
)

var (
	// login vars
	user   = ""
	pass   = ""
	appkey = ""

	// market catalogue filters
	maxResults  = 100
	countries   = []string{"AU"}           // GB, US, NZ
	eventTypes  = []string{"7"}            // 7=horse 4339=greyhound
	marketTypes = []string{"WIN", "PLACE"} // WIN, PLACE, LINE
	projection  = []string{"MARKET_START_TIME", "MARKET_DESCRIPTION", "EVENT", "EVENT_TYPE", "RUNNER_DESCRIPTION"}
)

func main() {

	// create a client
	client := bfapi.NewClient(appkey, nil, nil)

	// init and login to bf
	if _, err := client.Login(user, pass); err != nil {
		log.Fatalf("Failed to login: %v", err)
	}

	// request a catalogue of markets meeting the specified filter/projection
	mcat, err := client.ListMarketCatalogue(bfapi.ListMarketCatalogueArg{
		Filter: bfapi.MarketListFilter{
			Countries:       countries,
			EventTypeIds:    eventTypes,
			MarketTypeCodes: marketTypes,
		},
		MaxResults:       maxResults,
		MarketProjection: projection,
	})
	if err != nil {
		log.Fatalf("Could not retrieve catalogue: %v", err)
	}

	for _, cat := range mcat {
		cos, err := client.ListCurrentOrders(bfapi.ListCurrentOrdersArgs{
			MarketIDs:       []bfapi.MarketID{cat.MarketID},
			OrderProjection: "EXECUTION_COMPLETE",
			SortDir:         "EARLIEST_TO_LATEST",
		})

		if err != nil {
			log.Println(fmt.Errorf("listCurrentOrders failed: %v", err))
			return
		}

		bs, _ := json.Marshal(cos)
		fmt.Println(string(bs))
	}

}
