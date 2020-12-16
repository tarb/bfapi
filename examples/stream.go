package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/tarb/bfapi"
)

var (
	// login vars

	username = ""
	password = ""
	appkey   = ""

	// stream market filters
	heartbeatMs = bfapi.Duration(4000)
	conflateMs  = bfapi.Duration(2000)
	countries   = []string{"AU"}        // GB, US, NZ
	eventTypes  = []string{"7", "4339"} // 7=horse 4339=greyhound
	marketTypes = []string{"WIN"}       // WIN, PLACE, LINE
	venues      = []string{}

	// stream data filters
	ladderLvls = 3
	fields     = []string{"EX_BEST_OFFERS_DISP", "EX_TRADED", "EX_TRADED_VOL", "EX_MARKET_DEF"}
)

func main() {

	// create a client
	bfClient := bfapi.NewClient(appkey, nil, nil)

	// init and login to bf
	if _, err := bfClient.Login(username, password); err != nil {
		log.Fatalf("Failed to login: %v", err)
	}

	// setup a stream handler and create a stream
	var sh StreamHandler
	stream := bfClient.NewTCPStream(&sh, bfapi.SubsMessage{
		Op: bfapi.MarketSub,
		ID: 1,
		// SegmentationEnabled: false,
		HeartbeatMs: heartbeatMs,
		// ConflateMs:  conflateMs,
		MarketFilter: &bfapi.MarketStreamFilter{
			CountryCodes: countries,
			EventTypeIds: eventTypes,
			MarketTypes:  marketTypes,
		},
		MarketDataFilter: &bfapi.MarketDataFilter{
			LadderLevels: ladderLvls,
			Fields:       fields,
		},
	})

	// stream := bfClient.NewTCPStream(&sh, bfapi.SubsMessage{
	// 	Op:                  bfapi.OrderSub,
	// 	ID:                  1,
	// 	SegmentationEnabled: false,
	// 	HeartbeatMs:         heartbeatMs,
	// 	ConflateMs:          conflateMs,
	// 	OrderFilter: &bfapi.OrderFilter{
	// 		IncludeOverallPosition:        true,
	// 		PartitionMatchedByStrategyRef: true,
	// 		CustomerStrategyRefs:          []string{"New Rule"},
	// 	},
	// })

	// listen to the stream
	go stream.Listen()

	time.Sleep(time.Hour)
	stream.Close()
}

// StreamHandler implements bfapi.StreamHandler
type StreamHandler struct {
}

// OnConnect called on new stream connections, with bfapi ConnectionMessage
func (s *StreamHandler) OnConnect(cm bfapi.ConnectionMessage) {
	bs, err := json.Marshal(cm)
	if err != nil {
		fmt.Println("OnConnect", err)
	}
	fmt.Println(string(bs))
}

// OnStatus called on stream connection changes, with bfapi StatusMessage
func (s *StreamHandler) OnStatus(sm bfapi.StatusMessage) {
	bs, err := json.Marshal(sm)
	if err != nil {
		fmt.Println("OnStatus", err)
	}
	fmt.Println(string(bs))
}

// OnChange called with market or order changes, or HEARTBEAT msg at the specified
// interval to show stream connectivity if no other changes were sent
func (s *StreamHandler) OnChange(cm bfapi.ChangeMessage) {
	bs, err := json.Marshal(cm)
	if err != nil {
		fmt.Println("OnChange", err)
	}
	fmt.Println(string(bs))
}

// OnClose called when stream disconects/ends or when close is called
func (s *StreamHandler) OnClose(err error) {
	fmt.Println("OnClose", err)
	// bs, _ := json.Marshal(cm)
	// fmt.Println(string(bs))
}
