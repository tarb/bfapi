package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/tarb/tarbot/lib/bfapi"
// )

// var (
// 	// login vars
// 	username = ""
// 	password = ""
// 	appkey   = ""

// 	// market catalogue filters
// 	maxResults = 100
// 	projection = []string{"MARKET_START_TIME", "MARKET_DESCRIPTION", "EVENT", "EVENT_TYPE", "RUNNER_METADATA"}
// )

// func main() {

// 	// create a client
// 	bfClient := bfapi.NewClient(appkey, nil, nil)

// 	// init and login to bf
// 	if _, err := bfClient.Login(username, password); err != nil {
// 		log.Fatalf("Failed to login: %v", err)
// 	}

// 	// request a catalogue of markets meeting the specified filter/projection
// 	mcat, err := bfClient.ListMarketCatalogue(bfapi.ListMarketCatalogueArg{
// 		Filter: bfapi.MarketListFilter{
// 			MarketIds: []bfapi.MarketID{
// 				bfapi.NewMarketID([]byte("1.128151441")), //next president
// 				bfapi.NewMarketID([]byte("1.128988348")), //winning party (repu/demo)
// 				bfapi.NewMarketID([]byte("1.131485565")), //next gender
// 				bfapi.NewMarketID([]byte("1.153132662")), //nominee forcast - pick both nominees
// 				bfapi.NewMarketID([]byte("1.128999265")), //repub nominee
// 				bfapi.NewMarketID([]byte("1.128161111")), //demo nominee
// 			},
// 		},
// 		MaxResults:       maxResults,
// 		MarketProjection: projection,
// 	})
// 	if err != nil {
// 		log.Fatalf("Could not retrieve catalogue: %v", err)
// 	}

// 	fmt.Printf("mid,mname,selid,selname,selmat,ltp,atb0p,atl0p\n")

// 	// range over the catalogue and print basic details in form
// 	// <num> <MarketID> <time as HH:MM> <Venue>-<CountryCode> <MarketName> (<EventTypeName>)
// 	// 42. 1.143886544 11:52 Cranbourne-AU        R10 311m Gr5    (Greyhound Racing)
// 	for _, c := range mcat {

// 		arg := bfapi.ListMarketBookArg{
// 			MarketIds: []bfapi.MarketID{c.MarketID},
// 			PriceProjection: bfapi.PriceProjection{
// 				PriceData:  []string{"EX_BEST_OFFERS", "EX_TRADED", "SP_TRADED"},
// 				Virtualise: true,
// 			},
// 			CurrencyCode: "AUD",
// 		}
// 		mbook, err := bfClient.ListMarketBook(arg)
// 		if err != nil {
// 			log.Fatalf("Could not retrieve %v book: %v", c.MarketID, err)
// 		}

// 		book := mbook[0]

// 		for _, r := range book.Runners {

// 			var rc *bfapi.RunnerCatalogue
// 			for i := range c.Runners {
// 				if r.SelectionID == c.Runners[i].SelectionID {
// 					rc = &c.Runners[i]
// 				}
// 			}
// 			if rc == nil {
// 				log.Fatal("could not find runner")
// 			}

// 			if r.Status == "ACTIVE" {
// 				b0, l0 := 0.0, 0.0
// 				if len(r.Ex.AvailableToBack) > 0 {
// 					b0 = r.Ex.AvailableToBack[0].Price
// 				}
// 				if len(r.Ex.AvailableToLay) > 0 {
// 					l0 = r.Ex.AvailableToLay[0].Price
// 				}

// 				fmt.Printf("%s,%s,%d,%s,%.2f,%.2f,%.2f,%.2f\n",
// 					string(c.MarketID.Encode()), c.MarketName, r.SelectionID, rc.RunnerName,
// 					r.TotalMatched, r.LastPriceTraded, b0, l0,
// 				)
// 			} else {
// 				fmt.Printf("%s,%s,%d,%s\n", string(c.MarketID.Encode()), c.MarketName, r.SelectionID, rc.RunnerName)
// 			}

// 		}

// 		fmt.Println()
// 	}
// }
