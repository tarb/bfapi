package bfapi

// ListArg parameter for ListMarketTypes/ListEventTypes/ListCompetitions/ListEvents/ListCountries/ListVenues
type ListArg struct {
	Filter MarketListFilter `json:"filter"`
	Locale string           `json:"locale,omitempty"`
}

//
type ListMarketCatalogueArg struct {
	Sort             string           `json:"sort,omitempty"`
	MarketProjection []string         `json:"marketProjection,omitempty"` //COMPETITION, EVENT, EVENT_TYPE, MARKET_START_TIME, MARKET_DESCRIPTION, RUNNER_DESCRIPTION, RUNNER_METADATA
	MaxResults       int              `json:"maxResults,omitempty"`
	Filter           MarketListFilter `json:"filter,omitempty"`
	Locale           string           `json:"locale,omitempty"`
}

//
type MarketListFilter struct {
	TextQuery          string        `json:"textQuery,omitempty"`
	MarketIds          []MarketID    `json:"marketIds,omitempty"`
	EventTypeIds       []string      `json:"eventTypeIds,omitempty"`
	EventIds           []string      `json:"eventIds,omitempty"`
	Countries          []string      `json:"marketCountries,omitempty"`
	MarketTypeCodes    []string      `json:"marketTypeCodes,omitempty"`
	MarketBettingTypes []string      `json:"marketBettingTypes,omitempty"`
	RaceTypes          []string      `json:"raceTypes,omitempty"`
	CompetitionIds     []string      `json:"competitionIds,omitempty"`
	Venues             []string      `json:"venues,omitempty"`
	WithOrders         []OrderStatus `json:"withOrders,omitempty"`
	StartTime          *TimeRange    `json:"marketStartTime,omitempty"`
	BspOnly            *bool         `json:"bspOnly,omitempty"`
	TurnInPlayEnabled  *bool         `json:"turnInPlayEnabled,omitempty"`
	InPlayOnly         *bool         `json:"inPlayOnly,omitempty"`
}

//
type MarketCatalogue struct {
	MarketID        MarketID           `json:"marketId,omitempty"`
	MarketName      string             `json:"marketName,omitempty"`
	MarketStartTime Time               `json:"marketStartTime,omitempty"`
	Description     *MarketDescription `json:"description,omitempty"`
	TotalMatched    float64            `json:"totalMatched,omitempty"`
	Runners         []RunnerCatalogue  `json:"runners,omitempty"`
	EventType       *EventType         `json:"eventType,omitempty"`
	Event           *Event             `json:"event,omitempty"`
	Competition     *Competition       `json:"competition,omitempty"`
}

//
type RunnerCatalogue struct {
	SelectionID  SelectionID     `json:"selectionId,omitempty"`
	RunnerName   string          `json:"runnerName,omitempty"`
	Handicap     float64         `json:"handicap,omitempty"`
	SortPriority int             `json:"sortPriority,omitempty"`
	Metadata     *RunnerMetadata `json:"metadata,omitempty"`
}

//
type RunnerMetadata struct {
	RunnerID                 string `json:"runnerId,omitempty"`
	SireName                 string `json:"SIRE_NAME,omitempty"`
	ClothNumberAlpha         string `json:"CLOTH_NUMBER_ALPHA,omitempty"`
	OfficialRating           string `json:"OFFICIAL_RATING,omitempty"`
	ColoursDescription       string `json:"COLOURS_DESCRIPTION,omitempty"`
	ColoursFileName          string `json:"COLOURS_FILENAME,omitempty"`
	ForecastPriceDenominator string `json:"FORECASTPRICE_DENOMINATOR,omitempty"`
	DamSireName              string `json:"DAMSIRE_NAME,omitempty"`
	WeightValue              string `json:"WEIGHT_VALUE,omitempty"`
	SexType                  string `json:"SEX_TYPE,omitempty"`
	DaysSinceLastRun         string `json:"DAYS_SINCE_LAST_RUN,omitempty"`
	Wearing                  string `json:"WEARING,omitempty"`
	OwnerName                string `json:"OWNER_NAME,omitempty"`
	DamYearBorn              string `json:"DAM_YEAR_BORN,omitempty"`
	SireBred                 string `json:"SIRE_BRED,omitempty"`
	JockeyName               string `json:"JOCKEY_NAME,omitempty"`
	DamBred                  string `json:"DAM_BRED,omitempty"`
	AdjustedRating           string `json:"ADJUSTED_RATING,omitempty"`
	ClothNumber              string `json:"CLOTH_NUMBER,omitempty"`
	SireYearBorn             string `json:"SIRE_YEAR_BORN,omitempty"`
	TrainerName              string `json:"TRAINER_NAME,omitempty"`
	ColourType               string `json:"COLOUR_TYPE,omitempty"`
	Age                      string `json:"AGE,omitempty"`
	DamsireBred              string `json:"DAMSIRE_BRED,omitempty"`
	JockeyClaim              string `json:"JOCKEY_CLAIM,omitempty"`
	Form                     string `json:"FORM,omitempty"`
	ForecastPriceNumerator   string `json:"FORECASTPRICE_NUMERATOR,omitempty"`
	Bred                     string `json:"BRED,omitempty"`
	DamName                  string `json:"DAM_NAME,omitempty"`
	DamSireYearBorn          string `json:"DAMSIRE_YEAR_BORN,omitempty"`
	StallDraw                string `json:"STALL_DRAW,omitempty"`
	WeightUnits              string `json:"WEIGHT_UNITS,omitempty"`
}

//
type CountryCodeResult struct {
	CountryCode string `json:"countryCode"` //The ISO-2 code for the event.  A list of ISO-2 codes is available via http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
	Count       int    `json:"marketCount"` //Count of markets associated with this Country Code
}

//
type VenueResult struct {
	Venue string `json:"venue"`
	Count int    `json:"marketCount"` //Count of markets associated with this Venue
}

//
type EventResult struct {
	Event Event `json:"event"`
	Count int   `json:"marketCount"` //Count of markets associated with this event
}

//
type Event struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	Timezone    string `json:"timezone,omitempty"`
	Venue       string `json:"venue,omitempty"`
	OpenDate    Time   `json:"openDate,omitempty"`
}

//
type EventTypeResult struct {
	EventType   EventType `json:"eventType"`   //The ID identifying the Event Type
	MarketCount int       `json:"marketCount"` //Count of markets associated with this eventType
}

//
type EventType struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

//
type MarketDescription struct {
	MarketTime             Time                   `json:"marketTime,omitempty"`
	SuspendTime            Time                   `json:"suspendTime,omitempty"`
	SettleTime             Time                   `json:"settleTime,omitempty"`
	BettingType            string                 `json:"bettingType,omitempty"` //ODDS/LINE/RANGE/ASIAN_HANDICAP_DOUBLE_LINE/ASIAN_HANDICAP_SINGLE_LINE/FIXED_ODDS
	MarketType             string                 `json:"marketType,omitempty"`
	Regulator              string                 `json:"regulator,omitempty"`
	Wallet                 string                 `json:"wallet,omitempty"`
	Rules                  string                 `json:"rules,omitempty"`
	Clarifications         string                 `json:"clarifications,omitempty"`
	MarketBaseRate         float64                `json:"marketBaseRate,omitempty"`
	EachWayDivisor         float64                `json:"eachWayDivisor,omitempty"`
	LineRangeInfo          MarketLineRangeInfo    `json:"lineRangeInfo,omitempty"`
	PriceLadderDescription PriceLadderDescription `json:"priceLadderDescription,omitempty"`
	PersistenceEnabled     bool                   `json:"persistenceEnabled,omitempty"`
	BspMarket              bool                   `json:"bspMarket,omitempty"`
	TurnInPlayEnabled      bool                   `json:"turnInPlayEnabled,omitempty"`
	DiscountAllowed        bool                   `json:"discountAllowed,omitempty"`
	RulesHasDate           bool                   `json:"rulesHasDate,omitempty"`
}

//
type MarketLineRangeInfo struct {
	MaxUnitValue float64 `json:"maxUnitValue,omitempty"`
	MinUnitValue float64 `json:"minUnitValue,omitempty"`
	Interval     float64 `json:"interval,omitempty"`
	MarketUnit   string  `json:"marketUnit,omitempty"`
}

//
type PriceLadderDescription struct {
	Type string `json:"type,omitempty"` //CLASSIC/FINEST/LINE_RANGE
}

//
type StartingPrice struct {
	BackStakeTaken    []PriceSize `json:"backStakeTaken,omitempty"`
	LayLiabilityTaken []PriceSize `json:"layLiabilityTaken,omitempty"`
	ActualSP          float64     `json:"actualSP,omitempty"`
}

//
type ExchangePrice struct {
	AvailableToBack []PriceSize `json:"availableToBack,omitempty"`
	AvailableToLay  []PriceSize `json:"availableToLay,omitempty"`
	TradedVolume    []PriceSize `json:"tradedVolume,omitempty"`
}

//
type PriceProjection struct {
	PriceData             []string              `json:"priceData,omitempty"` //SP_AVAILABLE,SP_TRADED,EX_BEST_OFFERS,EX_ALL_OFFERS,EX_ALL_OFFERS,EX_TRADED
	ExBestOffersOverrides ExBestOffersOverrides `json:"exBestOffersOverrides,omitempty"`
	Virtualise            bool                  `json:"virtualise,omitempty"`
	RolloverStakes        bool                  `json:"rolloverStakes,omitempty"`
}

//
type ExBestOffersOverrides struct {
	BestPricesDepth          int     `json:"bestPricesDepth,omitempty"`
	RollupModel              string  `json:"rollupModel,omitempty"` //STAKE,PAYOUT,MANAGED_LIABILITY,NONE
	RollupLimit              int     `json:"rollupLimit,omitempty"`
	RollupLiabilityThreshold float64 `json:"rollupLiabilityThreshold,omitempty"`
	RollupLiabilityFactor    int     `json:"rollupLiabilityFactor,omitempty"`
}

//
type ListMarketBookArg struct {
	MarketIds                     []MarketID      `json:"marketIds,omitempty"` //REQUIRED
	PriceProjection               PriceProjection `json:"priceProjection,omitempty"`
	OrderProjection               string          `json:"orderProjection,omitempty"` //ALL,EXECUTABLE,EXECUTION_COMPLETE
	MatchProjection               string          `json:"matchProjection,omitempty"` //NO_ROLLUP,ROLLED_UP_BY_PRICE,ROLLED_UP_BY_AVG_PRICE
	IncludeOverallPosition        bool            `json:"includeOverallPosition,omitempty"`
	PartitionMatchedByStrategyRef bool            `json:"partitionMatchedByStrategyRef,omitempty"`
	CustomerStrategyRefs          []string        `json:"customerStrategyRefs,omitempty"`
	CurrencyCode                  string          `json:"currencyCode,omitempty"`
	Locale                        string          `json:"locale,omitempty"`
	MatchedSince                  Time            `json:"matchedSince,omitempty"`
	BetIds                        []string        `json:"betIds,omitempty"`
}

//
type ListRunnerBookArg struct {
	MarketID                      MarketID        `json:"marketId,omitempty"`    //REQUIRED
	SelectionID                   SelectionID     `json:"selectionId,omitempty"` //The unique id for the selection in the market.
	PriceProjection               PriceProjection `json:"priceProjection,omitempty"`
	OrderProjection               string          `json:"orderProjection,omitempty"` //ALL,EXECUTABLE,EXECUTION_COMPLETE
	MatchProjection               string          `json:"matchProjection,omitempty"` //NO_ROLLUP,ROLLED_UP_BY_PRICE,ROLLED_UP_BY_AVG_PRICE
	IncludeOverallPosition        bool            `json:"includeOverallPosition,omitempty"`
	PartitionMatchedByStrategyRef bool            `json:"partitionMatchedByStrategyRef,omitempty"`
	CustomerStrategyRefs          []string        `json:"customerStrategyRefs,omitempty"`
	CurrencyCode                  string          `json:"currencyCode,omitempty"`
	Locale                        string          `json:"locale,omitempty"`
	MatchedSince                  Time            `json:"matchedSince,omitempty"`
	BetIds                        []string        `json:"betIds,omitempty"`
}

//
type MarketBook struct {
	MarketID              MarketID     `json:"marketId,omitempty"`
	BetDelay              int          `json:"betDelay,omitempty"`
	NumberOfWinners       int          `json:"numberOfWinners,omitempty"`
	NumberOfRunners       int          `json:"numberOfRunners,omitempty"`
	NumberOfActiveRunners int          `json:"numberOfActiveRunners,omitempty"`
	LastMatchTime         Time         `json:"lastMatchTime,omitempty"`
	TotalMatched          float64      `json:"totalMatched,omitempty"`
	TotalAvailable        float64      `json:"totalAvailable,omitempty"`
	Version               int64        `json:"version,omitempty"`
	Runners               []Runner     `json:"runners,omitempty"`
	KeyLineDescription    string       `json:"keyLineDescription,omitempty"` //?
	Status                MarketStatus `json:"status,omitempty"`             //INACTIVE,OPEN,SUSPENDED,CLOSED
	CrossMatching         bool         `json:"crossMatching,omitempty"`
	RunnersVoidable       bool         `json:"runnersVoidable,omitempty"`
	IsMarketDataDelayed   bool         `json:"isMarketDataDelayed,omitempty"`
	BspReconciled         bool         `json:"bspReconciled,omitempty"`
	Complete              bool         `json:"complete,omitempty"`
	Inplay                bool         `json:"inplay,omitempty"`
}

//
type Runner struct {
	SelectionID      SelectionID    `json:"selectionId,omitempty"`
	Handicap         float64        `json:"handicap,omitempty"`
	AdjustmentFactor float64        `json:"adjustmentFactor,omitempty"`
	LastPriceTraded  float64        `json:"lastPriceTraded,omitempty"`
	TotalMatched     float64        `json:"totalMatched,omitempty"`
	RemovalDate      Time           `json:"removalDate,omitempty"`
	Sp               StartingPrices `json:"sp,omitempty"`
	Ex               ExchangePrices `json:"ex,omitempty"`
	Status           RunnerStatus   `json:"status,omitempty"` //ACTIVE, REMOVED, WINNER, PLACED, LOSER, HIDDEN
	// orders            Order            `json:"orders,omitempty"`
	// matches           Match      `json:"matches,omitempty"`
	// matchesByStrategy map[string]Matches        `json:"matchesByStrategy,omitempty"`
}

//
type StartingPrices struct {
	NearPrice         float64     `json:"nearPrice,omitempty"`
	FarPrice          float64     `json:"farPrice,omitempty"`
	BackStakeTaken    []PriceSize `json:"backStakeTaken,omitempty"`
	LayLiabilityTaken []PriceSize `json:"layLiabilityTaken,omitempty"`
	ActualSP          float64     `json:"actualSP,omitempty"`
}

//
type ExchangePrices struct {
	AvailableToBack []PriceSize `json:"availableToBack,omitempty"`
	AvailableToLay  []PriceSize `json:"availableToLay,omitempty"`
	TradedVolume    []PriceSize `json:"tradedVolume,omitempty"`
}

//
type PriceSize struct {
	Price float64 `json:"price,omitempty"`
	Size  float64 `json:"size,omitempty"`
}

//
type MarketVersion struct {
	Version int64 `json:"version,omitempty"`
}

//
type PlaceOrderArg struct {
	MarketID            MarketID           `json:"marketId,omitempty"`
	Instructions        []PlaceInstruction `json:"instructions,omitempty"`
	CustomerRef         string             `json:"customerRef,omitempty"`
	MarketVersion       MarketVersion      `json:"marketVersion,omitempty"`
	CustomerStrategyRef string             `json:"customerStrategyRef,omitempty"`
	Async               bool               `json:"async,omitempty"`
}

//
type PlaceInstruction struct {
	SelectionID        SelectionID         `json:"selectionId,omitempty"`
	Handicap           string              `json:"handicap,omitempty"`
	CustomerOrderRef   string              `json:"customerOrderRef,omitempty"`
	LimitOrder         *LimitOrder         `json:"limitOrder,omitempty"`
	LimitOnCloseOrder  *LimitOnCloseOrder  `json:"limitOnCloseOrder,omitempty"`
	MarketOnCloseOrder *MarketOnCloseOrder `json:"marketOnCloseOrder,omitempty"`
	Side               Side                `json:"side,omitempty"`
	OrderType          OrderType           `json:"orderType,omitempty"`
}

//
type PlaceOrderRequest struct {
	MarketID            MarketID           `json:"marketId,omitempty"`
	Instructions        []PlaceInstruction `json:"instructions,omitempty"`
	CustomerRef         string             `json:"customerRef,omitempty"`
	MarketVersion       MarketVersion      `json:"marketVersion,omitempty"`
	CustomerStrategyRef string             `json:"customerStrategyRef,omitempty"`
	Async               bool               `json:"async,omitempty"`
}

//
type PlaceExecutionReport struct {
	CustomerRef        string                   `json:"customerRef"`
	MarketID           MarketID                 `json:"marketId"`
	Status             string                   `json:"status"`
	ErrorCode          string                   `json:"errorCode"`
	InstructionReports []PlaceInstructionReport `json:"instructionReports"`
}

//
type PlaceInstructionReport struct {
	Status              string           `json:"status"`
	ErrorCode           string           `json:"errorCode"`
	OrderStatus         string           `json:"orderStatus"`
	Instruction         PlaceInstruction `json:"instruction"`
	BetID               string           `json:"betId"`
	PlacedDate          string           `json:"placedDate"`
	AveragePriceMatched float64          `json:"averagePriceMatched"`
	SizeMatched         float64          `json:"sizeMatched"`
}

//
type LimitOnCloseOrder struct {
	Liability float64 `json:"liability,omitempty"`
	Price     float64 `json:"price,omitempty"`
}

//
type MarketOnCloseOrder struct {
	Liability float64 `json:"liability,omitempty"`
}

// LimitOrder - values are pointers so they can be ommitted if not set
type LimitOrder struct {
	MinFillSize     float64         `json:"minFillSize,omitempty"`
	BetTargetSize   float64         `json:"betTargetSize,omitempty"`
	Size            float64         `json:"size,omitempty"`
	Price           float64         `json:"price,omitempty"`
	TimeInForce     TimeInForce     `json:"timeInForce,omitempty"`
	BetTargetType   BetTargetType   `json:"betTargetType,omitempty"`
	PersistenceType PersistenceType `json:"persistenceType,omitempty"`
}

//
type CancelInstruction struct {
	BetID         string  `json:"betId,omitempty"`
	SizeReduction float64 `json:"sizeReduction,omitempty"`
}

//
type CancelOrderRequest struct {
	MarketID     MarketID            `json:"marketId,omitempty"`
	Instructions []CancelInstruction `json:"instructions,omitempty"`
	CustomerRef  string              `json:"customerRef,omitempty"`
}

//
type CancelInstructionReport struct {
	Status        string            `json:"status,omitempty"`
	ErrorCode     string            `json:"errorCode,omitempty"`
	Instruction   CancelInstruction `json:"instruction,omitempty"`
	SizeCancelled float64           `json:"sizeCancelled,omitempty"`
	CancelDate    string            `json:"cancelledDate,omitempty"`
}

//
type CancelExecutionReport struct {
	CustomerRef        string                    `json:"customerRef,omitempty"`
	Status             string                    `json:"status,omitempty"`
	ErrorCode          string                    `json:"errorCode,omitempty"`
	MarketID           MarketID                  `json:"marketId,omitempty"`
	InstructionReports []CancelInstructionReport `json:"instructionReports,omitempty"`
}

//
type ReplaceInstruction struct {
	BetID    string  `json:"betId,omitempty"`
	NewPrice float64 `json:"newPrice,omitempty"`
}

//
type ReplaceOrderRequest struct {
	MarketID      MarketID             `json:"marketId,omitempty"`
	Instructions  []ReplaceInstruction `json:"instructions,omitempty"`
	CustomerRef   string               `json:"customerRef,omitempty"`
	MarketVersion MarketVersion        `json:"marketVersion,omitempty"`
	Async         bool                 `json:"async,omitempty"`
}

//
type ReplaceInstructionReport struct {
	Status                  string                  `json:"status,omitempty"`
	ErrorCode               string                  `json:"errorCode,omitempty"`
	CancelInstructionReport CancelInstructionReport `json:"cancelInstructionReport,omitempty"`
	PlaceInstructionReport  PlaceInstructionReport  `json:"placeInstructionReport,omitempty"`
}

//
type ReplaceExecutionReport struct {
	CustomerRef        string                     `json:"customerRef,omitempty"`
	Status             string                     `json:"status,omitempty"`
	ErrorCode          string                     `json:"errorCode,omitempty"`
	MarketID           MarketID                   `json:"marketId,omitempty"`
	InstructionReports []ReplaceInstructionReport `json:"instructionReports,omitempty"`
}

//
type UpdateInstruction struct {
	BetID              string `json:"betId,omitempty"`
	NewPersistenceType string `json:"newPersistenceType,omitempty"`
}

//
type UpdateOrderRequest struct {
	MarketID     MarketID            `json:"marketId,omitempty"`
	Instructions []UpdateInstruction `json:"instructions,omitempty"`
	CustomerRef  string              `json:"customerRef,omitempty"`
}

//
type UpdateInstructionReport struct {
	Status      string            `json:"status,omitempty"`
	ErrorCode   string            `json:"errorCode,omitempty"`
	Instruction UpdateInstruction `json:"instruction,omitempty"`
}

//
type UpdateExecutionReport struct {
	CustomerRef        string                    `json:"customerRef,omitempty"`
	Status             string                    `json:"status,omitempty"`
	ErrorCode          string                    `json:"errorCode,omitempty"`
	MarketID           MarketID                  `json:"marketId,omitempty"`
	InstructionReports []UpdateInstructionReport `json:"instructionReports,omitempty"`
}

//
type ListClearedOrdersArgs struct {
	BetStatus              string     `json:"betStatus,omitempty"` //SETTLED | VOIDED | LAPSED | CANCELLED
	EventTypeIDs           []string   `json:"eventTypeIds,omitempty"`
	EventIDs               []string   `json:"eventIds,omitempty"`
	MarketIDs              []MarketID `json:"marketIds,omitempty"`
	RunnerIDs              []int64    `json:"runnerIds,omitempty"`
	BetIDs                 []string   `json:"betIds,omitempty"`
	CustomerOrderRefs      []string   `json:"customerOrderRefs,omitempty"`
	CustomerStrategyRefs   []string   `json:"customerStrategyRefs,omitempty"`
	Side                   Side       `json:"side,omitempty"` // BACK | LAY
	SettledDateRange       *TimeRange `json:"settledDateRange,omitempty"`
	GroupBy                string     `json:"groupBy,omitempty"` // EVENT_TYPE | EVENT | MARKET | SIDE | BET
	IncludeItemDescription bool       `json:"includeItemDescription,omitempty"`
	Locale                 string     `json:"locale,omitempty"`
	FromRecord             int        `json:"fromRecord,omitempty"`
	RecordCount            int        `json:"recordCount,omitempty"`
}

//
type ClearedOrderSummaryReport struct {
	ClearedOrders []ClearedOrderSummary `json:"clearedOrders"`
	MoreAvailable bool                  `json:"moreAvailable"`
}

//
type ClearedOrderSummary struct {
	MarketID            MarketID         `json:"marketId"`
	SelectionID         SelectionID      `json:"selectionId"`
	Handicap            float64          `json:"handicap"`
	ItemDescription     *ItemDescription `json:"itemDescription"`
	PlacedDate          Time             `json:"placedDate"`
	SettledDate         Time             `json:"settledDate"`
	LastMatchedDate     Time             `json:"lastMatchedDate"`
	BetCount            int              `json:"betCount"`
	PriceRequested      float64          `json:"priceRequested"`
	Commission          float64          `json:"commission"`
	PriceMatched        float64          `json:"priceMatched"`
	SizeSettled         float64          `json:"sizeSettled"`
	Profit              float64          `json:"profit"`
	SizeCancelled       float64          `json:"sizeCancelled"`
	EventTypeID         string           `json:"eventTypeId"`
	EventID             string           `json:"eventId"`
	BetID               string           `json:"betId"`
	BetOutcome          string           `json:"betOutcome"`
	CustomerOrderRef    string           `json:"customerOrderRef"`
	CustomerStrategyRef string           `json:"customerStrategyRef"`
	PersistenceType     PersistenceType  `json:"persistenceType"` // LAPSE | PERSIST | MARKET_ON_CLOSE
	OrderType           OrderType        `json:"orderType"`       // LIMIT | LIMIT_ON_CLOSE | MARKET_ON_CLOSE
	Side                Side             `json:"side"`            // BACK | LAY
	PriceReduced        bool             `json:"priceReduced"`
}

//
type ListCurrentOrdersArgs struct {
	BetIDs               []string   `json:"betIds,omitempty"`
	MarketIDs            []MarketID `json:"marketIds,omitempty"`
	OrderProjection      string     `json:"orderProjection,omitempty"` // ALL, EXECUTABLE, EXECUTION_COMPLETE
	CustomerOrderRefs    []string   `json:"customerOrderRefs,omitempty"`
	CustomerStrategyRefs []string   `json:"customerStrategyRefs,omitempty"`
	DateRange            *TimeRange `json:"dateRange,omitempty"`
	OrderBy              string     `json:"orderBy,omitempty"` // BY_MARKET,BY_MATCH_TIME,BY_PLACE_TIME,BY_SETTLED_TIME,BY_VOID_TIME
	SortDir              string     `json:"sortDir,omitempty"` //EARLIEST_TO_LATEST, LATEST_TO_EARLIEST
	FromRecord           int        `json:"fromRecord,omitempty"`
	RecordCount          int        `json:"recordCount,omitempty"`
}

//
type CurrentOrderSummaryReport struct {
	CurrentOrders []CurrentOrderSummary `json:"currentOrders"`
	MoreAvailable bool                  `json:"moreAvailable"`
}

//
type CurrentOrderSummary struct {
	MarketID            MarketID        `json:"marketId"`
	SelectionID         SelectionID     `json:"selectionId"`
	PlacedDate          Time            `json:"placedDate"`
	MatchedDate         Time            `json:"matchedDate"`
	PriceSize           PriceSize       `json:"priceSize"`
	Handicap            float64         `json:"handicap"`
	BspLiability        float64         `json:"bspLiability"`
	AveragePriceMatched float64         `json:"averagePriceMatched"`
	SizeMatched         float64         `json:"sizeMatched"`
	SizeRemaining       float64         `json:"sizeRemaining"`
	SizeLapsed          float64         `json:"sizeLapsed"`
	SizeCancelled       float64         `json:"sizeCancelled"`
	SizeVoided          float64         `json:"sizeVoided"`
	BetID               string          `json:"betId"`
	RegulatorAuthCode   string          `json:"regulatorAuthCode"`
	RegulatorCode       string          `json:"regulatorCode"`
	CustomerOrderRef    string          `json:"customerOrderRef"`
	CustomerStrategyRef string          `json:"customerStrategyRef"`
	Side                Side            `json:"side"`            // BACK | LAY
	Status              OrderStatus     `json:"status"`          // EXECUTABLE,EXECUTION_COMPLETE
	PersistenceType     PersistenceType `json:"persistenceType"` // LAPSE | PERSIST | MARKET_ON_CLOSE
	OrderType           OrderType       `json:"orderType"`       // LIMIT | LIMIT_ON_CLOSE | MARKET_ON_CLOSE
}

//
type ItemDescription struct {
	EventTypeDesc   string  `json:"eventTypeDesc"`
	EventDesc       string  `json:"eventDesc"`
	MarketDesc      string  `json:"marketDesc"`
	MarketType      string  `json:"marketType"`
	RunnerDesc      string  `json:"runnerDesc"`
	MarketStartTime Time    `json:"marketStartTime"`
	NumberOfWinners int     `json:"numberOfWinners"`
	EachWayDivisor  float64 `json:"eachWayDivisor"`
}

//
type TimeRange struct {
	To   Time `json:"to"`
	From Time `json:"from"`
}

//
type MarketTypeResult struct {
	MarketType  string `json:"marketType"`  // Market Type
	MarketCount int    `json:"marketCount"` // Count of markets associated with this marketType
}

//
type Competition struct {
	ID   string `json:"id"`   // id
	Name string `json:"name"` // name
}

//
type CompetitionResult struct {
	Competition       Competition `json:"competition"`       // Competition
	MarketCount       int         `json:"marketCount"`       // Count of markets associated with this competition
	CompetitionRegion string      `json:"competitionRegion"` // Region in which this competition is happening
}

//
type ListMarketProfitAndLossArg struct {
	MarketIDs          []MarketID `json:"marketIds,omitempty"`          // List of markets to calculate profit and loss
	IncludeSettledBets bool       `json:"includeSettledBets,omitempty"` // Option to include settled bets (partially settled markets only). Defaults to false if not specified.
	IncludeBspBets     bool       `json:"includeBspBets,omitempty"`     // Option to include BSP bets. Defaults to false if not specified.
	NetOfCommission    bool       `json:"netOfCommission,omitempty"`    // Option to return profit and loss net of users current commission rate for this market including any special tariffs. Defaults to false if not specified.
}

//
type MarketProfitAndLoss struct {
	MarketID          MarketID              `json:"marketId"`          // The unique identifier for the market
	CommissionApplied float64               `json:"commissionApplied"` // The commission rate applied to P&L values. Only returned if netOfCommision option is requested
	ProfitAndLosses   []RunnerProfitAndLoss `json:"profitAndLosses"`   // Calculated profit and loss data.
}

//
type RunnerProfitAndLoss struct {
	SelectionID SelectionID `json:"selectionId"` // The unique identifier for the selection
	IfWin       float64     `json:"ifWin"`       // Profit or loss for the market if this selection is the winner.
	IfLose      float64     `json:"ifLose"`      // Profit or loss for the market if this selection is the loser. Only returned for multi-winner odds markets.
	IfPlace     float64     `json:"ifPlace"`     // Profit or loss for the market if this selection is placed. Applies to marketType EACH_WAY only.
}
