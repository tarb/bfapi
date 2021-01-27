# BFAPI - Betfair Api

Implementation of Betfair's Exchange, Account and Stream Api


#### Normal non-vendor login flow

```go
// create a client
// can pass in a http.Client if you want to specifiy custom timeouts/proxy/options
// can pass in a certificate if you want to login with certlogin
client := bfapi.NewClient("APPKEY", nil, nil)

// call login
token, err := client.Login("USERNAME", "PASSWORD")
// login will call either interactive login or certLogin depending on whether a certificate was supplied
```

The login expires every 8 hours but KeepAlive can be called to extend this time (by another 8 hours). 
```go
go func() {
    ticker := time.NewTicker(7 * time.Hour)
    for range ticker.C {
        if err := client.KeepAlive(); err != nil {
            // ...
        }
    }    
}
```

#### Methods

ListMarketCatalogue can be used to retrieve a broad overview of a selection of markets

Example of ListMarketCatalogu
```go
	// request a catalogue of markets meeting the specified filter/projection
	cat, err := client.ListMarketCatalogue(bfapi.ListMarketCatalogueArg{
		Filter: bfapi.MarketListFilter{
			Countries:  []string{"AU", "GB"},
			EventTypes: []string{"7"},
			TypeCodes:  []string{"WIN",
		},
		MaxResults:       200,
		MarketProjection: []string{"MARKET_START_TIME", "MARKET_DESCRIPTION", "EVENT", "EVENT_TYPE", "RUNNER_DESCRIPTION", "RUNNER_METADATA"},
	})

```

#### Stream

```go
    // setup a stream handler and create a stream
    var sh StreamHandler
    // connecting to a market stream
	stream := client.NewTCPStream(&sh, bfapi.SubsMessage{
		Op: bfapi.MarketSub,
		ID: 1,
		HeartbeatMs: heartbeatMs,
		MarketFilter: &bfapi.MarketStreamFilter{
			CountryCodes: []string{"AU", "GB"},,
			EventTypeIds: []string{"7"},,
			MarketTypes:  []string{"WIN",,
		},
		MarketDataFilter: &bfapi.MarketDataFilter{
			LadderLevels: 3,
			Fields:       []string{"EX_BEST_OFFERS_DISP", "EX_TRADED", "EX_TRADED_VOL", "EX_MARKET_DEF"},
		},
	})

    // connecting to an order stream
	// stream := bfClient.NewTCPStream(&sh, bfapi.SubsMessage{
	// 	Op:                  bfapi.OrderSub,
	// 	ID:                  1,
	// 	SegmentationEnabled: false,
	// 	HeartbeatMs:         heartbeatMs,
	// 	ConflateMs:          conflateMs,
	// 	OrderFilter: &bfapi.OrderFilter{
	// 		IncludeOverallPosition:        true,
	// 		PartitionMatchedByStrategyRef: true,
	// 		CustomerStrategyRefs:          []string{"My Cool Strategy"},
	// 	},
	// })

	// listen to the stream for an hour
	go stream.Listen()
	time.Sleep(time.Hour)
	stream.Close()
}

// StreamHandler implements bfapi.StreamHandler
type StreamHandler struct {
}

// OnConnect called on new stream connections, with bfapi ConnectionMessage
func (s *StreamHandler) OnConnect(cm bfapi.ConnectionMessage) {
    fmt.Println(cm)
}

// OnStatus called on stream connection changes, with bfapi StatusMessage
func (s *StreamHandler) OnStatus(sm bfapi.StatusMessage) {
    fmt.Println(sm)
}

// OnChange called with market or order changes, or HEARTBEAT msg at the specified
// interval to show stream connectivity if no other changes were sent
func (s *StreamHandler) OnChange(cm bfapi.ChangeMessage) {
    fmt.Println(cm)
}

// OnClose called when stream disconects/ends or when close is called
func (s *StreamHandler) OnClose(err error) {
    fmt.Println(err)
}

```


### TODO

* Documentation - Function/TypeDef comments
* Documentation - list the other methods
* Add more Betfair methods and types 
