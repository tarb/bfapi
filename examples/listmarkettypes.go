package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"

// 	"github.com/tarb/bfapi"
// )

// var (
// 	// login vars
// 	username = ""
// 	password = ""
// 	appkey   = ""
// )

// func main() {

// 	// create a client
// 	bfClient := bfapi.NewClient(appkey, nil, nil)

// 	// init and login to bf
// 	if _, err := bfClient.Login(username, password); err != nil {
// 		log.Fatalf("Failed to login: %v", err)
// 	}

// 	// types, err := bfClient.ListCompetitions(bfapi.ListArg{
// 	// 	Filter: bfapi.MarketListFilter{
// 	// 		// Countries: []string{"AU"},
// 	// 	},
// 	// })

// 	rates, err := bfClient.ListCurrencyRates()
// 	if err != nil {
// 		log.Fatalf("Could not retrieve catalogue: %v", err)
// 	}

// 	bs, err := json.Marshal(rates)
// 	if err == nil {
// 		fmt.Println(string(bs))
// 	} else {
// 		fmt.Println(err)
// 	}

// }
