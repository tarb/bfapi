package bfapi

import (
	"bufio"
	"crypto/rand"
	"crypto/tls"
	"encoding/json"
	"errors"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/buger/jsonparser"
	"github.com/mailru/easyjson"
)

//
type StreamHandler interface {
	OnConnect(ConnectionMessage)
	OnStatus(StatusMessage)
	OnChange(ChangeMessage)
	OnClose(error)
}

//
type Stream interface {
	Listen()
	Close()
	Details() ConnDetails
}

//
type tcpStream struct {
	conn     net.Conn
	closer   sync.Once
	certs    []tls.Certificate
	appkey   string
	submsg   SubsMessage
	handler  StreamHandler
	details  atomic.Value // holding *ConnDetails
	getToken func() Token
}

//
type ConnDetails struct {
	Heartbeat Duration
	Conflate  Duration
	InitClk   string
	Clk       string
	Pt        Time
	System    Time
	Latency   Duration
	Inited    bool
}

/*
	experimental npt implementation for more accurate latency info
		server 0.europe.pool.ntp.org
		server 1.europe.pool.ntp.org
		server 2.europe.pool.ntp.org
		server 3.europe.pool.ntp.org
*/
// var latencyOffset time.Duration

// func init() {
// 	response, err := ntp.Query("0.europe.pool.ntp.org")
// 	if err != nil {
// 		panic("could not fetch ntp")
// 	}
// 	latencyOffset = response.ClockOffset
// }

/* experimental npt implementation for more accurate latency info */

// ConnTimeout  Dial timeout
const ConnTimeout = 30 * time.Second

//
var ErrHeartbeatTimeout = errors.New("heartbeat timeout")

// regex for finding incomming op codes - no longer used
// var re = regexp.MustCompile(`"op"\s*:\s*"([a-zA-Z]+)\s*"`)
// op := string(re.FindSubmatch(jsonBytes)[1])

// NewTCPStream returns a new tcp stream connecting to the Betfair stream API
func (c *Client) NewTCPStream(hdlr StreamHandler, submsg SubsMessage) Stream {
	stream := &tcpStream{
		handler:  hdlr,
		getToken: c.Token,
		appkey:   c.appKey,
		submsg:   submsg,
	}

	if c.certificate != nil {
		stream.certs = append(stream.certs, *c.certificate)
	}

	stream.details.Store(ConnDetails{})

	return stream
}

//
func (sc *tcpStream) listenFailure() {
	dur := ConnTimeout
	ticker := time.NewTicker(dur)

	for t := range ticker.C {
		d := sc.Details()
		hb := d.Heartbeat.ToStdDuration()

		if hb != 0 && (hb*2) != dur {
			dur = hb * 2
			ticker = time.NewTicker(dur)
		}

		if t.Sub(d.Pt.ToStdTime()) > hb {
			sc.close(ErrHeartbeatTimeout)
			break
		}
	}
}

//
func (sc *tcpStream) Listen() {
	var err error
	sc.conn, err = tls.DialWithDialer(&net.Dialer{Timeout: ConnTimeout, KeepAlive: ConnTimeout}, "tcp", streamHost, &tls.Config{
		Certificates:       sc.certs,
		InsecureSkipVerify: true,
		Renegotiation:      tls.RenegotiateFreelyAsClient,
		Rand:               rand.Reader,
	})
	if err != nil {
		sc.handler.OnClose(err)
		return
	}

	// reset the closer
	sc.closer = sync.Once{}

	// write auth and sub msgs - this is the only writing we need to do to the socket
	go func() {
		enc := json.NewEncoder(sc.conn)

		if err := enc.Encode(AuthMessage{Op: Auth, ID: 0, Session: sc.getToken().Token, AppKey: sc.appkey}); err != nil {
			sc.close(err)
			return
		}

		details := sc.Details()
		submsg := sc.submsg
		submsg.InitialClk = details.InitClk
		submsg.Clk = details.Clk
		submsg.HeartbeatMs = details.Heartbeat
		// submsg.SegmentationEnabled = true  // cant actually seem to make betfair use this... even with the limit raised to 2000 markets

		if err := enc.Encode(submsg); err != nil {
			sc.close(err)
		}
	}()

	// lisen on the socket
	in := bufio.NewReaderSize(sc.conn, 2<<20)
	for {
		jsonBytes, err := in.ReadSlice('\n')
		if err != nil {
			sc.close(err)
			return
		}

		op, err := jsonparser.GetUnsafeString(jsonBytes, "op")
		if err != nil {
			sc.close(err)
			return
		}

		switch op {
		case Conn:
			var cm ConnectionMessage
			if err := easyjson.Unmarshal(jsonBytes, &cm); err != nil {
				sc.close(err)
				return
			}

			sc.handler.OnConnect(cm)
			go sc.listenFailure()

		case Status:
			var sm StatusMessage
			if err := easyjson.Unmarshal(jsonBytes, &sm); err != nil {
				sc.close(err)
				return
			}

			sc.handler.OnStatus(sm)

		case Mcm, Ocm:
			var mc ChangeMessage
			if err := easyjson.Unmarshal(jsonBytes, &mc); err != nil {
				sc.close(err)
				return
			}

			// if mc.SegmentType == SegmentTypeSegStart {

			// }

			sc.updateDetails(&mc)
			sc.handler.OnChange(mc)
		}
	}
}

//
func (sc *tcpStream) close(err error) {
	sc.closer.Do(func() {
		sc.conn.Close()
		sc.handler.OnClose(err)
	})
}

//
func (sc *tcpStream) Close() { sc.close(nil) }

//
func (sc *tcpStream) Details() ConnDetails {
	return sc.details.Load().(ConnDetails)
}

//
func (sc *tcpStream) updateDetails(mc *ChangeMessage) {
	var details ConnDetails
	if v := sc.details.Load(); v != nil {
		details = v.(ConnDetails)
	}

	if mc.Ct == ChangeTypeHeartbeat {
		details.Inited = true
	}
	if mc.InitialClk != "" {
		details.InitClk = mc.InitialClk
	}
	if mc.Clk != "" {
		details.Clk = mc.Clk
	}
	if mc.HeartbeatMs != 0 {
		details.Heartbeat = mc.HeartbeatMs
	}
	if mc.ConflateMs != 0 {
		details.Conflate = mc.ConflateMs
	}
	if mc.Pt != 0 {
		details.Pt = mc.Pt
		details.System = Time(time.Now().UnixNano() / 1e6)
		details.Latency = Duration(int64(details.System) - int64(mc.Pt))
	}

	sc.details.Store(details)
}
