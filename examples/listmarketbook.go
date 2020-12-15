package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"time"

// 	"github.com/tarb/bfapi"
// )

// var (
// 	// login vars
// 	user   = ""
// 	pass   = ""
// 	appkey = ""

// 	// market catalogue filters
// 	maxResults  = 200
// 	countries   = []string{"GB"}   // GB, US, NZ
// 	eventTypes  = []string{"4339"} // 7=horse 4339=greyhound
// 	marketTypes = []string{"WIN"}  // WIN, PLACE, LINE
// 	projection  = []string{"MARKET_START_TIME", "MARKET_DESCRIPTION", "EVENT", "EVENT_TYPE", "RUNNER_DESCRIPTION"}
// )

// func main() {

// 	// create a client
// 	bfClient := bfapi.NewClient(appkey, nil, nil)

// 	arg := bfapi.ListMarketBookArg{
// 		MarketIds: []string{"1.1234567890", "1.0987654321"},
// 		PriceProjection: bfapi.PriceProjection{
// 			PriceData:  []string{"EX_BEST_OFFERS", "EX_TRADED", "SP_TRADED"},
// 			Virtualise: true,
// 		},
// 		OrderProjection:        "ALL",
// 		IncludeOverallPosition: true,
// 		CurrencyCode:           "AUD",
// 	}

// 	bs, _ := json.Marshal(arg)
// 	fmt.Println(string(bs))

// 	// init and login to bf
// 	if _, err := bfClient.Login(user, pass); err != nil {
// 		log.Fatalf("Failed to login: %v", err)
// 	}

// 	// request a catalogue of markets meeting the specified filter/projection
// 	mcat, err := bfClient.ListMarketCatalogue(bfapi.ListMarketCatalogueArg{
// 		Filter: bfapi.MarketListFilter{
// 			Countries:  countries,
// 			EventTypes: eventTypes,
// 			TypeCodes:  marketTypes,
// 		},
// 		MaxResults:       maxResults,
// 		MarketProjection: projection,
// 	})
// 	if err != nil {
// 		log.Fatalf("Could not retrieve catalogue: %v", err)
// 	}

// 	if len(mcat) > 0 {
// 		// pick a random market from the first 1/4 of the catalogue
// 		c := mcat[rand.Intn(len(mcat)/4)]

// 		// every 3 seconds call listMarketBook
// 		for range time.NewTicker(3 * time.Second).C {

// 			arg := bfapi.ListMarketBookArg{
// 				MarketIds: []string{c.MarketID},
// 				PriceProjection: bfapi.PriceProjection{
// 					PriceData:  []string{"EX_BEST_OFFERS", "EX_TRADED", "SP_TRADED"},
// 					Virtualise: true,
// 				},
// 				OrderProjection:        "ALL",
// 				IncludeOverallPosition: true,
// 				CurrencyCode:           "AUD",
// 			}

// 			bs, _ := json.Marshal(arg)
// 			fmt.Println(string(bs))

// 			book, err := bfClient.ListMarketBook(arg)
// 			if err != nil {
// 				log.Printf("Error retrieving marketBook: %v", err)
// 				continue
// 			}

// 			// print off the runner name and its ladder
// 			for i := range c.Runners {
// 				fmt.Printf("%-20s %v %v\n", c.Runners[i].RunnerName, book[0].Runners[i].Ex.AvailableToBack, book[0].Runners[i].Ex.AvailableToLay)
// 			}
// 			fmt.Println()
// 		}
// 	}
// }
