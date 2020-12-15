package main

// import (
// 	"flag"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/tarb/bfapi"
// )

// var (
// 	// login vars
// 	user   = ""
// 	pass   = ""
// 	appkey = ""

// 	// stream market filters
// 	heartbeatMs = bfapi.Duration(2000)
// 	conflateMs  = bfapi.Duration(0)

// 	// stream data filters
// 	ladderLvls = 3
// 	fields     = []string{"EX_BEST_OFFERS", "EX_TRADED", "EX_TRADED_VOL", "EX_MARKET_DEF"}
// )

// func main() {

// 	mid := flag.String("mid", "", "market id")
// 	flag.Parse()

// 	if *mid == "" {
// 		log.Fatalf("Need to specify market id")
// 	}

// 	// create a client
// 	bfClient := bfapi.NewClient(appkey, nil, nil)

// 	// init and login to bf
// 	if _, err := bfClient.Login(user, pass, appkey); err != nil {
// 		log.Fatalf("Failed to login: %v", err)
// 	}

// 	// setup a stream handler and create a stream
// 	var sh StreamHandler
// 	stream := bfClient.NewTCPStream(&sh, bfapi.SubsMessage{
// 		Op:                  bfapi.MarketSub,
// 		ID:                  1,
// 		SegmentationEnabled: false,
// 		HeartbeatMs:         heartbeatMs,
// 		ConflateMs:          conflateMs,
// 		MarketFilter: &bfapi.MarketStreamFilter{
// 			MarketIds: []string{*mid},
// 		},
// 		MarketDataFilter: &bfapi.MarketDataFilter{
// 			LadderLevels: ladderLvls,
// 			Fields:       fields,
// 		},
// 	})
// 	defer stream.Close()

// 	// listen to the stream
// 	go stream.ListenReadIn()
// 	go stream.ListenWriteOut()

// 	var matched float64

// 	// every 3 seconds call listMarketBook
// 	for range time.NewTicker(100 * time.Millisecond).C {
// 		go func() {
// 			arg := bfapi.ListMarketBookArg{
// 				MarketIds: []string{*mid},
// 				PriceProjection: bfapi.PriceProjection{
// 					PriceData:  []string{"EX_BEST_OFFERS", "EX_TRADED", "SP_TRADED"},
// 					Virtualise: false,
// 				},
// 				CurrencyCode: "GBP",
// 			}

// 			book, err := bfClient.ListMarketBook(arg)
// 			if err != nil {
// 				log.Printf("Error retrieving marketBook: %v", err)
// 				return
// 			}

// 			for i := range book {
// 				if tv := book[i].TotalMatched; tv != matched && tv > matched {
// 					fmt.Printf("lmb,%d,%.2f\n", time.Now().UnixNano()/1e6, tv)
// 					matched = tv
// 				}
// 			}
// 		}()
// 	}
// }

// // StreamHandler implements bfapi.StreamHandler
// type StreamHandler struct {
// 	tv bfapi.PriceVol
// }

// // OnChange called with market or order changes, or HEARTBEAT msg at the specified
// // interval to show stream connectivity if no other changes were sent
// func (s *StreamHandler) OnChange(cm bfapi.ChangeMessage) {
// 	if cm.Ct == "HEARTBEAT" {
// 		return
// 	}

// 	for i := range cm.Mc {
// 		change := &cm.Mc[i]
// 		if change.Tv != s.tv && change.Tv != 0 {
// 			fmt.Printf("stream,%d,%.2f\n", time.Now().UnixNano()/1e6, change.Tv)
// 			s.tv = change.Tv
// 		}
// 	}
// }

// // OnConnect called on new stream connections, with bfapi ConnectionMessage
// func (s *StreamHandler) OnConnect(cm bfapi.ConnectionMessage) {}

// // OnStatus called on stream connection changes, with bfapi StatusMessage
// func (s *StreamHandler) OnStatus(sm bfapi.StatusMessage) {}

// // OnClose called when stream disconects/ends or when close is called
// func (s *StreamHandler) OnClose() {}
