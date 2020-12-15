package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/tarb/bfapi"
)

var (
	username = ""
	password = ""
	appkey   = ""
)

func main() {
	var list, gen, add bool
	var days int
	var ref, token string

	flag.BoolVar(&list, "list", false, "list the betfair tokens")
	flag.BoolVar(&gen, "gen", false, "generate a new token")
	flag.BoolVar(&add, "add", false, "add time to token")
	flag.StringVar(&token, "vcid", "", "vcid to add time too")
	flag.StringVar(&ref, "ref", "", "client ref for new token")
	flag.IntVar(&days, "len", 0, "length (in days) for new token")
	flag.Parse()

	// if list == gen || add == list {
	// 	fmt.Println("select gen or list")
	// }

	// create a client
	bfc := bfapi.NewClient(appkey, nil, nil)

	// init and login to bf
	_, err := bfc.Login(username, password)
	if err != nil {
		log.Fatalf("Failed to login: %v", err)
	}

	if list {
		printKeys(bfc)
	}

	if gen {
		genKey(bfc, days, ref)
	}

	if add {
		addTime(bfc, token, days)
	}

}

func printKeys(bfc *bfapi.Client) {
	keys, err := bfc.ListApplicationSubscriptionTokens("ALL")
	if err != nil {
		log.Fatalf("Failed to fetch keys: %v", err)
	}

	fmt.Println("Token,SubscriptionStatus,ClientReference,VendorClientId,ExpiryDateTime,ExpiredDateTime,CreatedDateTime,ActivationDateTime,CancellationDateTime")

	for _, k := range keys {
		var expiryStr, expiredStr, createdStr, activationStr, cancellationStr string
		if k.ExpiryDateTime != 0 {
			expiryStr = k.ExpiryDateTime.ToStdTime().Format("02-01-2006 15:04:05")
		}
		if k.ExpiredDateTime != 0 {
			expiredStr = k.ExpiredDateTime.ToStdTime().Format("02-01-2006 15:04:05")
		}
		if k.CreatedDateTime != 0 {
			createdStr = k.CreatedDateTime.ToStdTime().Format("02-01-2006 15:04:05")
		}
		if k.ActivationDateTime != 0 {
			activationStr = k.ActivationDateTime.ToStdTime().Format("02-01-2006 15:04:05")
		}
		if k.CancellationDateTime != 0 {
			cancellationStr = k.CancellationDateTime.ToStdTime().Format("02-01-2006 15:04:05")
		}

		fmt.Printf("%s,%s,%s,%s,%s,%s,%s,%s,%s\n",
			k.SubscriptionToken,
			k.SubscriptionStatus,
			k.ClientReference,
			k.VendorClientID,
			expiryStr,
			expiredStr,
			createdStr,
			activationStr,
			cancellationStr,
		)
	}
}

func genKey(bfc *bfapi.Client, days int, ref string) {
	key, err := bfc.GetApplicationSubscriptionToken(days, ref)
	if err != nil {
		log.Fatalf("Failed to gen keys: %v", err)
	}

	fmt.Println(key)
}

func addTime(bfc *bfapi.Client, vcid string, days int) {
	t, err := bfc.UpdateApplicationSubscription(vcid, days)
	if err != nil {
		log.Fatalf("Failed to add time: %v", err)
	}

	fmt.Println(t)
}
