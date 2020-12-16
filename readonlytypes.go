package bfapi

// https://www.betfair.com.au/www/sports/exchange/readonly/v1/bymarket?alt=json&currencyCode=AUD&locale=en&rollupModel=STAKE&types=MARKET_STATE,MARKET_RATES,MARKET_DESCRIPTION,EVENT,RUNNER_DESCRIPTION,RUNNER_STATE,RUNNER_METADATA,RUNNER_SP&marketIds=1.152396018,1.152396022,1.152396044
// https://www.betfair.com.au/api/sports/exchange/readonly/v1/bymarket?alt=json&currencyCode=AUD&locale=en&rollupModel=STAKE&types=MARKET_STATE,MARKET_RATES,MARKET_DESCRIPTION,EVENT,RUNNER_DESCRIPTION,RUNNER_STATE,RUNNER_METADATA,RUNNER_SP&marketIds=1.152396018,1.152396022,1.152396044

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
