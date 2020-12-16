package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/tarb/bfapi"
)

// go run placeorders.go

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
	projection  = []string{"MARKET_START_TIME", "MARKET_DESCRIPTION", "EVENT", "EVENT_TYPE", "RUNNER_METADATA"}
)

func main() {

	// create a client
	bfClient := bfapi.NewClient(appkey, nil, nil)

	// init and login to bf
	if _, err := bfClient.Login(user, pass); err != nil {
		log.Fatalf("Failed to login: %v", err)
	}

	// request a catalogue of markets meeting the specified filter/projection
	mcat, err := bfClient.ListMarketCatalogue(bfapi.ListMarketCatalogueArg{
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
	} else if len(mcat) == 0 {
		log.Fatal("No markets returned in catalogue")
	}

	m := mcat[0]

	// place bets on first market, first selection
	per, err := bfClient.PlaceOrders(bfapi.PlaceOrderArg{
		MarketID: m.MarketID,
		Instructions: []bfapi.PlaceInstruction{
			{ // the bet to place
				OrderType:   bfapi.OrderTypeLimit,
				SelectionID: m.Runners[0].SelectionID,
				Side:        bfapi.SideBack,
				LimitOrder: &bfapi.LimitOrder{
					TimeInForce: bfapi.TimeInForceFillOrKill,
					Price:       1000,
					Size:        0.10,
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("Could not retrieve catalogue: %v", err)
	}

	cosr, err := bfClient.ListCurrentOrders(bfapi.ListCurrentOrdersArgs{
		BetIDs:          []string{per.InstructionReports[0].BetID},
		OrderProjection: "ALL",
	})
	if err != nil {
		log.Fatalf("Could not retrieve catalogue: %v", err)
	}

	// print report as json
	bs, _ := json.Marshal(per)
	fmt.Println(string(bs))

	// print report as json
	bs, _ = json.Marshal(cosr)
	fmt.Println(string(bs))
}
