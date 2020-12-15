package bfapi

// This is an incomplete list of the types
// see more from this Rob Hare email
// ---------------------------------------------------
// Following on from our catch up last week; here is the detail to be able to access Exchange Read Only (ERO):

// https://www.betfair.com.au/www/sports/exchange/readonly/v1/bymarket?alt=json&currencyCode=AUD&locale=en& rollupModel=STAKE&types=MARKET_STATE,MARKET_RATES,MARKET_DESCRIPTION,EVENT,RUNNER_DESCRIPTION,RUNNER_STATE,RUNNER_METADATA,RUNNER_SP&marketIds=1.152396018, 1.152396022,1.152396044

// If you paste the above URL in to a browser, you will see a market from 12/12/2018 (should disappear in about 6 days as 3 month retention).
// You’ll find that you don’t need X-Application or X-Authentication to access this and call it as a straight GET.
// You’ll also notice multiple marketIds, you can send one or a few (not sure on limit) so batch retrieval.
// The response will be delayed. This is preferred approach for now just to stay in line with commercials.

// The alternative approach (and better for server to server) is to change the /www/ in to /api/ and specific X-Application (your vendor key) in the header. The /www/ version SSOID from cookie not header.
// https://www.betfair.com.au/api/sports/exchange/readonly/v1/bymarket?alt=json&currencyCode=AUD&locale=en& rollupModel=STAKE&types=MARKET_STATE,MARKET_RATES,MARKET_DESCRIPTION,EVENT,RUNNER_DESCRIPTION,RUNNER_STATE,RUNNER_METADATA,RUNNER_SP&marketIds=1.152396018, 1.152396022,1.152396044
// If you want live pricing, you can specific valid X-Authenticaion (but you can get live price from public API). To use this approach, you’ll need to work through the vendor relationship that Kerrin has been working with you on.

//
type ROResult struct {
	Events []ROEventType `json:"eventTypes"`
}

//
type ROEventType struct {
	EventTypeID uint64        `json:"eventTypeId"`
	EventNodes  []ROEventNode `json:"eventNodes"`
}

//
type ROEventNode struct {
	EventID     uint64         `json:"eventId"`
	MarketNodes []ROMarketNode `json:"marketNodes"`
	Event       struct {
		EventName   string `json:"eventName"`
		CountryCode string `json:"countryCode"`
		Timezone    string `json:"timezone"`
		Venue       string `json:"venue"`
		OpenDate    Time   `json:"openDate"`
	} `json:"event"`
}

//
type ROMarketNode struct {
	MarketID    MarketID            `json:"marketId"`
	Runners     []RORunner          `json:"runners"`
	Description ROMarketDescription `json:"description"`
	Rates       struct {
		MarketBaseRate  float64 `json:"marketBaseRate"`
		DiscountAllowed bool    `json:"discountAllowed"`
	} `json:"rates"`
}

//
type ROMarketDescription struct {
	MarketName             string  `json:"marketName"`
	MarketTime             Time    `json:"marketTime"`
	SuspendTime            Time    `json:"suspendTime"`
	SettleTime             Time    `json:"settleTime"`
	MarketType             string  `json:"marketType"`
	Regulator              string  `json:"regulator"`
	EachWayDivisor         float64 `json:"eachWayDivisor"`
	PriceLadderDescription struct {
		Type string `json:"type"`
	} `json:"priceLadderDescription"`
	BettingType        string `json:"bettingType"`
	PersistenceEnabled bool   `json:"persistenceEnabled"`
	BspMarket          bool   `json:"bspMarket"`
	TurnInPlayEnabled  bool   `json:"turnInPlayEnabled"`
}

//
type RORunner struct {
	SelectionID SelectionID `json:"selectionId"`
	Handicap    float64     `json:"handicap"`
	Description struct {
		RunnerName string `json:"runnerName"`
		Metadata   struct {
			RunnerID uint64 `json:"runnerId,string"`
		} `json:"metadata"`
	} `json:"description"`
	State struct {
		SortPriority int    `json:"sortPriority"`
		Status       string `json:"status"`
	} `json:"state"`
	Sp struct {
		ActualStartingPrice PriceVol `json:"actualStartingPrice"`
	} `json:"sp"`
}
