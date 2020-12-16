package bfapi

//RequestMessages OP codes
const (
	Auth      string = "authentication"
	MarketSub string = "marketSubscription"
	OrderSub  string = "orderSubscription"
	Heartbeat string = "heartbeat"
)

//ResponseMessages OP codes
const (
	Conn   string = "connection"
	Status string = "status"
	Mcm    string = "mcm"
	Ocm    string = "ocm"
)

//
type AuthMessage struct {
	Op      string `json:"op,omitempty"`
	ID      int    `json:"id,omitempty"`
	Session string `json:"session,omitempty"`
	AppKey  string `json:"appKey,omitempty"`
}

//
type HeartbeatMessage struct {
	Op string `json:"op,omitempty"`
	ID int    `json:"id,omitempty"`
}

//
type SubsMessage struct {
	Op                  string              `json:"op,omitempty"`
	InitialClk          string              `json:"initialClk,omitempty"`
	Clk                 string              `json:"clk,omitempty"`
	ID                  int                 `json:"id,omitempty"`
	HeartbeatMs         Duration            `json:"heartbeatMs,omitempty"`
	ConflateMs          Duration            `json:"conflateMs"`
	MarketFilter        *MarketStreamFilter `json:"marketFilter,omitempty"`
	MarketDataFilter    *MarketDataFilter   `json:"marketDataFilter,omitempty"`
	OrderFilter         *OrderFilter        `json:"orderFilter,omitempty"`
	SegmentationEnabled bool                `json:"segmentationEnabled"`
}

//
type MarketDataFilter struct {
	LadderLevels int      `json:"ladderLevels,omitempty"`
	Fields       []string `json:"fields,omitempty"`
}

//
type MarketStreamFilter struct {
	MarketIds         []string `json:"marketIds,omitempty"`
	CountryCodes      []string `json:"countryCodes,omitempty"`
	BettingTypes      []string `json:"bettingTypes,omitempty"`
	MarketTypes       []string `json:"marketTypes,omitempty"`
	Venues            []string `json:"venues,omitempty"`
	EventTypeIds      []string `json:"eventTypeIds,omitempty"`
	EventIds          []string `json:"eventIds,omitempty"`
	TurnInPlayEnabled bool     `json:"turnInPlayEnabled,omitempty"`
	BspMarket         bool     `json:"bspMarket,omitempty"`
}

//
type OrderFilter struct {
	AccountIds                    []int64  `json:"accountIds,omitempty"`
	CustomerStrategyRefs          []string `json:"customerStrategyRefs,omitempty"`
	IncludeOverallPosition        bool     `json:"includeOverallPosition"`
	PartitionMatchedByStrategyRef bool     `json:"partitionMatchedByStrategyRef"`
}

//
type ConnectionMessage struct {
	Op           string `json:"op"`
	ConnectionID string `json:"connectionId"`
	ID           int    `json:"id"`
}

//
type StatusMessage struct {
	ID               int    `json:"id"`
	Op               string `json:"op"`
	ErrorMessage     string `json:"errorMessage"`
	ErrorCode        string `json:"errorCode"`
	ConnectionID     string `json:"connectionId"`
	StatusCode       string `json:"statusCode"`
	ConnectionClosed bool   `json:"connectionClosed"`
}

//
type ChangeMessage struct {
	Op          string         `json:"op"`
	ID          int            `json:"id"`
	Clk         string         `json:"clk"`
	HeartbeatMs Duration       `json:"heartbeatMs"`
	Pt          Time           `json:"pt"`
	InitialClk  string         `json:"initialClk"`
	ConflateMs  Duration       `json:"conflateMs"`
	SegmentType SegmentType    `json:"segmentType"`
	Mc          []MarketChange `json:"mc"`
	Oc          []OrderChange  `json:"oc"`
	Ct          ChangeType     `json:"ct"`
}

//
type MarketChange struct {
	ID               MarketID          `json:"id"`
	Rc               []RunnerChange    `json:"rc"`
	MarketDefinition *MarketDefinition `json:"marketDefinition"`
	Tv               PriceVol          `json:"tv"` //total volume
	Img              bool              `json:"img"`
	Con              bool              `json:"con"`
}

//
type RunnerChange struct {
	ID    SelectionID   `json:"id"`    //RunnerId
	Spn   PriceVol      `json:"spn"`   //predicted sp
	Spf   PriceVol      `json:"sbf"`   //starting price projection prices
	Ltp   PriceVol      `json:"ltp"`   //last traded price
	Tv    PriceVol      `json:"tv"`    //traded volume
	Bdatb [][3]PriceVol `json:"bdatb"` //available to back
	Bdatl [][3]PriceVol `json:"bdatl"` //available to lay
	Batb  [][3]PriceVol `json:"batb"`  //available to back
	Batl  [][3]PriceVol `json:"batl"`  //available to lay
	Atb   [][2]PriceVol `json:"atb"`   //available to back
	Atl   [][2]PriceVol `json:"atl"`   //available to lay
	Trd   [][2]PriceVol `json:"trd"`   //traded bets
	Spb   [][2]PriceVol `json:"spb"`   //starting price bets
	Spl   [][2]PriceVol `json:"spl"`   //starting price lays
}

//
type MarketDefinition struct {
	//priceLadderDefinition -missing
	//keyLineDefinition - missing
	Venue                 string             `json:"venue,omitempty"`
	Timezone              string             `json:"timezone,omitempty"`
	MarketType            string             `json:"marketType,omitempty"`
	CountryCode           string             `json:"countryCode,omitempty"`
	BettingType           string             `json:"bettingType,omitempty"`
	EventID               string             `json:"eventId,omitempty"`
	EventTypeID           string             `json:"eventTypeId,omitempty"`
	Version               int64              `json:"version,omitempty"`
	BetDelay              int                `json:"betDelay,omitempty"`
	NumberOfWinners       int                `json:"numberOfWinners,omitempty"`
	NumberOfActiveRunners int                `json:"numberOfActiveRunners,omitempty"`
	EachWayDivisor        float64            `json:"eachWayDivisor,omitempty"`
	MarketBaseRate        float64            `json:"marketBaseRate,omitempty"`
	LineMaxUnit           float64            `json:"lineMaxUnit,omitempty"`
	LineMinUnit           float64            `json:"lineMinUnit,omitempty"`
	LineInterval          float64            `json:"lineInterval,omitempty"`
	Regulators            []string           `json:"regulators,omitempty"`
	Runners               []RunnerDefinition `json:"runners,omitempty"`
	SuspendTime           Time               `json:"suspendTime,omitempty"`
	SettledTime           Time               `json:"settledTime,omitempty"`
	OpenDate              Time               `json:"openDate,omitempty"`
	MarketTime            Time               `json:"marketTime,omitempty"`
	Status                MarketStatus       `json:"status,omitempty"`
	InPlay                bool               `json:"inPlay,omitempty"`
	BspMarket             bool               `json:"bspMarket,omitempty"`
	CrossMatching         bool               `json:"crossMatching,omitempty"`
	RunnersVoidable       bool               `json:"runnersVoidable,omitempty"`
	TurnInPlayEnabled     bool               `json:"turnInPlayEnabled,omitempty"`
	PersistenceEnabled    bool               `json:"persistenceEnabled,omitempty"`
	DiscountAllowed       bool               `json:"discountAllowed,omitempty"`
	Complete              bool               `json:"complete,omitempty"`
	BspReconciled         bool               `json:"bspReconciled,omitempty"`
}

//
type RunnerDefinition struct {
	SelectionName    string       `json:"name,omitempty"` // historic data only
	SortPriority     int          `json:"sortPriority,omitempty"`
	Sp               PriceVol     `json:"bsp,omitempty"`
	SelectionID      SelectionID  `json:"id,omitempty"`
	AdjustmentFactor PriceVol     `json:"adjustmentFactor,omitempty"`
	RemovalDate      Time         `json:"removalDate,omitempty"`
	Handicap         float64      `json:"hc,omitempty"`
	Status           RunnerStatus `json:"status,omitempty"`
}

//
type OrderChange struct {
	MarketID MarketID      `json:"id"`
	Runners  []RunnerOrder `json:"orc"`
	Closed   *bool         `json:"closed"`
}

//
type RunnerOrder struct {
	SelectionID  SelectionID                `json:"id"`
	Handicap     float64                    `json:"hc"`
	Unmatched    []UnmatchedOrder           `json:"uo"`
	StratMatched map[string]StrategyMatched `json:"smc"`
	MatchedBacks [][2]PriceVol              `json:"mb"`
	MatchedLays  [][2]PriceVol              `json:"ml"`
	Image        bool                       `json:"fullImage"`
}

//
type StrategyMatched struct {
	MatchedBacks [][2]PriceVol `json:"mb"`
	MatchedLays  [][2]PriceVol `json:"ml"`
}

//
type UnmatchedOrder struct {
	ID           string          `json:"id"`
	Ref          string          `json:"rfo"`
	Price        PriceVol        `json:"p"`
	Size         PriceVol        `json:"s"`
	BspLiability PriceVol        `json:"bsp"`
	PlacedAt     Time            `json:"pd"`
	MatchedAt    Time            `json:"md"`
	LapsedAt     Time            `json:"ld"`
	AvgPrice     PriceVol        `json:"avp"`
	SizeMatched  PriceVol        `json:"sm"`
	SizeRemaing  PriceVol        `json:"sr"`
	SizeLapse    PriceVol        `json:"sl"`
	SizeCanceled PriceVol        `json:"sc"`
	SizeVoided   PriceVol        `json:"sv"`
	Side         Side            `json:"side"`
	Status       OrderStatus     `json:"status"` //(E = EXECUTABLE, EC = EXECUTION_COMPLETE)
	PerType      PersistenceType `json:"pt"`     //(L = LAPSE, P = PERSIST, MOC = Market On Close)
	OrderType    OrderType       `json:"ot"`     //(L = LIMIT, MOC = MARKET_ON_CLOSE, LOC = LIMIT_ON_CLOSE)
}
