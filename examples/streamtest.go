package main

// import (
// 	"bufio"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/buger/jsonparser"
// 	"github.com/tarb/bfapi"
// )

// // This is a test file to look at the need to implement large json stream decoding.
// // Results still unclear, with benefits on either side.
// // Non stream - faster
// // Streamed can start processing marketchanges whilst still processing

// func main() {

// 	t1 := time.Now()
// 	streamed()
// 	fmt.Println(time.Now().Sub(t1))

// 	t2 := time.Now()
// 	buffered()
// 	fmt.Println(time.Now().Sub(t2))

// }

// func streamed() {
// 	f, err := os.Open("testjson2.json")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	dec := json.NewDecoder(f)
// 	for {

// 		t, err := dec.Token()
// 		if err == io.EOF {
// 			return
// 		}

// 		switch t {
// 		case json.Delim('{'):

// 			_, err := dec.Token()
// 			if err == io.EOF {
// 				log.Fatal(err)
// 			}
// 			// if typeKey != "op" {
// 			// 	log.Fatal("not leading json with op")
// 			// }

// 			var op string
// 			dec.Decode(&op)

// 			switch op {
// 			case bfapi.Mcm, bfapi.Ocm:
// 				msg := bfapi.ChangeMessage{Op: op}
// 				for dec.More() {
// 					key, _ := dec.Token()
// 					switch key {
// 					case "id":
// 						dec.Decode(&msg.ID)
// 					case "clk":
// 						dec.Decode(&msg.Clk)
// 					case "pt":
// 						dec.Decode(&msg.Pt)
// 					case "ct":
// 						dec.Decode(&msg.Ct)
// 					case "mc":
// 						arrayT, _ := dec.Token()
// 						if arrayT == json.Delim('[') {
// 							for dec.More() {
// 								var mc bfapi.MarketChange
// 								dec.Decode(&mc)
// 							}
// 						}
// 						dec.Token() // closing ']'
// 					case "oc":
// 						arrayT, _ := dec.Token()
// 						if arrayT == json.Delim('[') {
// 							for dec.More() {
// 								var oc bfapi.OrderChange
// 								dec.Decode(&oc)
// 								// fmt.Println(oc)
// 							}
// 						}
// 						dec.Token() // closing ']'
// 					case "conflateMs":
// 						dec.Decode(&msg.ConflateMs)
// 					case "heartbeatMs":
// 						dec.Decode(&msg.HeartbeatMs)
// 					case "initialClk":
// 						dec.Decode(&msg.InitialClk)
// 					}
// 				}

// 			case bfapi.Conn:
// 				msg := bfapi.ConnectionMessage{Op: op}
// 				for dec.More() {
// 					key, _ := dec.Token()
// 					switch key {
// 					case "id":
// 						dec.Decode(&msg.ID)
// 					case "connectionId":
// 						dec.Decode(&msg.ConnectionID)
// 					}
// 				}

// 			case bfapi.Status:
// 				msg := bfapi.StatusMessage{Op: op}
// 				for dec.More() {
// 					key, _ := dec.Token()
// 					switch key {
// 					case "id":
// 						dec.Decode(&msg.ID)
// 					case "errorMessage":
// 						dec.Decode(&msg.ErrorMessage)
// 					case "errorCode":
// 						dec.Decode(&msg.ErrorCode)
// 					case "connectionId":
// 						dec.Decode(&msg.ConnectionID)
// 					case "statusCode":
// 						dec.Decode(&msg.StatusCode)
// 					case "connectionClosed":
// 						dec.Decode(&msg.ConnectionClosed)
// 					}
// 				}

// 			}

// 		case json.Delim('}'):
// 		}
// 	}
// }

// func buffered() {
// 	f, err := os.Open("testjson2.json")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	in := bufio.NewReaderSize(f, 2<<20)

// 	for {
// 		jsonBytes, err := in.ReadSlice('\n')
// 		if err != nil {
// 			fmt.Println(err)
// 			// sc.disconected()
// 			return
// 		}

// 		op, err := jsonparser.GetUnsafeString(jsonBytes, "op")
// 		if err != nil {
// 			fmt.Println(err)
// 			// sc.disconected()
// 			return
// 		}

// 		// op := string(re.FindSubmatch(jsonBytes)[1])

// 		switch op {
// 		case bfapi.Conn:
// 			var cm bfapi.ConnectionMessage
// 			err = json.Unmarshal(jsonBytes, &cm)

// 			// err = easyjson.Unmarshal(jsonBytes, &cm)
// 			if err == nil {
// 				// sc.handler.OnConnect(cm)
// 			}

// 			// go sc.listenFailure()

// 		case bfapi.Status:
// 			var sm bfapi.StatusMessage
// 			err = json.Unmarshal(jsonBytes, &sm)

// 			// err = easyjson.Unmarshal(jsonBytes, &sm)
// 			if err == nil {
// 				// sc.handler.OnStatus(sm)
// 			}

// 		case bfapi.Mcm, bfapi.Ocm:
// 			var mc bfapi.ChangeMessage
// 			if err = json.Unmarshal(jsonBytes, &mc); err != nil {
// 				log.Fatal(err)
// 			}
// 			// if err = easyjson.Unmarshal(jsonBytes, &mc); err != nil {
// 			// 	log.Fatal(err)
// 			// }

// 			// sc.updateDetails(&mc)
// 			// sc.handler.OnChange(mc)
// 		}
// 	}
// }
