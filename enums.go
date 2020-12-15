package bfapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
)

//
const (
	SideBack Side = 1
	SideLay  Side = 2
)

var (
	sideBackJSON = []byte(`"BACK"`)
	sideLayJSON  = []byte(`"LAY"`)
	sideBackStr  = "BACK"
	sideLayStr   = "LAY"
)

//
type Side uint8

//
func SideFromString(s string) Side {
	if len(s) == 0 {
		return 0
	}

	switch strings.ToLower(s)[:1] {
	case "b":
		return SideBack
	case "l":
		return SideLay
	default:
		return 0
	}
}

func (s Side) String() string {
	if s == SideBack {
		return sideBackStr
	} else if s == SideLay {
		return sideLayStr
	} else {
		return ""
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *Side) UnmarshalJSON(data []byte) error {

	// side will either be a "B" / "BACK"
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	// make sure theres actually a value
	if len(str) == 0 {
		return errors.New("value empty")
	}

	switch strings.ToLower(str)[:1] {
	case "b":
		*s = SideBack
	case "l":
		*s = SideLay
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (s Side) MarshalJSON() ([]byte, error) {
	if s == SideBack {
		return sideBackJSON, nil
	} else if s == SideLay {
		return sideLayJSON, nil
	} else {
		return nil, errors.New("invalid side value")
	}
}

//
const (
	BetStatusInit       BetStatus = 1
	BetStatusPending    BetStatus = 2
	BetStatusCanceling  BetStatus = 3
	BetStatusExecutable BetStatus = 4
	BetStatusExComplete BetStatus = 5
	BetStatusFailed     BetStatus = 6
)

var (
	betStatusExecutableJSON      = []byte(`"EXECUTABLE"`)
	betStatusExecutableShortJSON = []byte(`"E"`)
	betStatusExCompleteJSON      = []byte(`"EXECUTION_COMPLETE"`)
	betStatusExCompleteShortJSON = []byte(`"EC"`)
)

//
type BetStatus uint8

// UnmarshalJSON implements the json.Unmarshaler interface.
func (bs *BetStatus) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, betStatusExecutableShortJSON) {
		*bs = BetStatusExecutable
	} else if bytes.Equal(data, betStatusExCompleteShortJSON) {
		*bs = BetStatusExComplete
	} else if bytes.Equal(data, betStatusExecutableJSON) {
		*bs = BetStatusExecutable
	} else if bytes.Equal(data, betStatusExCompleteJSON) {
		*bs = BetStatusExComplete
	} else {
		return errors.New("invalid BetStatus value in unmarshal - " + string(data))
	}

	return nil
}

func (bs BetStatus) String() string {
	switch bs {
	case BetStatusInit:
		return "INITIAL"
	case BetStatusPending:
		return "PENDING"
	case BetStatusCanceling:
		return "CANCELING"
	case BetStatusExecutable:
		return "EXECUTABLE"
	case BetStatusExComplete:
		return "EXECUTION_COMPLETE"
	case BetStatusFailed:
		return "FAILED"
	default:
		return ""
	}
}

//
const (
	PersistenceTypeLapse         PersistenceType = 1
	PersistenceTypePersist       PersistenceType = 2
	PersistenceTypeMarketOnClose PersistenceType = 3
)

var (
	ptLapseJSON        = []byte(`"LAPSE"`)
	ptPersistJSON      = []byte(`"PERSIST"`)
	ptMocJSON          = []byte(`"MARKET_ON_CLOSE"`)
	ptLapseShortJSON   = []byte(`"L"`)
	ptPersistShortJSON = []byte(`"P"`)
	ptMocShortJSON     = []byte(`"MOC"`)
	ptLapseStr         = "LAPSE"
	ptPersistStr       = "PERSIST"
	ptMocStr           = "MARKET_ON_CLOSE"
)

//
type PersistenceType uint8

//
func PersistenceTypeFromString(s string) PersistenceType {
	switch strings.ToUpper(s) {
	case ptLapseStr:
		return PersistenceTypeLapse
	case ptPersistStr:
		return PersistenceTypePersist
	case ptMocStr:
		return PersistenceTypeMarketOnClose
	default:
		return 0
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (pt *PersistenceType) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, ptLapseShortJSON) {
		*pt = PersistenceTypeLapse
	} else if bytes.Equal(data, ptMocShortJSON) {
		*pt = PersistenceTypeMarketOnClose
	} else if bytes.Equal(data, ptPersistShortJSON) {
		*pt = PersistenceTypePersist
	} else if bytes.Equal(data, ptLapseJSON) {
		*pt = PersistenceTypeLapse
	} else if bytes.Equal(data, ptPersistJSON) {
		*pt = PersistenceTypePersist
	} else if bytes.Equal(data, ptMocJSON) {
		*pt = PersistenceTypeMarketOnClose
	} else {
		return errors.New("invalid PersistenceType value in unmarshal - " + string(data))
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (pt PersistenceType) MarshalJSON() ([]byte, error) {
	switch pt {
	case PersistenceTypeLapse:
		return ptLapseJSON, nil
	case PersistenceTypePersist:
		return ptPersistJSON, nil
	case PersistenceTypeMarketOnClose:
		return ptMocJSON, nil
	default:
		return nil, errors.New("invalid PersistenceType value")
	}
}

func (pt PersistenceType) String() string {
	if pt == PersistenceTypeLapse {
		return ptLapseStr
	} else if pt == PersistenceTypePersist {
		return ptPersistStr
	} else if pt == PersistenceTypeMarketOnClose {
		return ptMocStr
	} else {
		return ""
	}
}

//
const (
	OrderTypeLimit         OrderType = 1
	OrderTypeLimitOnClose  OrderType = 2
	OrderTypeMarketOnClose OrderType = 3
)

var (
	otLimitJSON      = []byte(`"LIMIT"`)
	otMocJSON        = []byte(`"MARKET_ON_CLOSE"`)
	otLocJSON        = []byte(`"LIMIT_ON_CLOSE"`)
	otLimitShortJSON = []byte(`"L"`)
	otMocShortJSON   = []byte(`"MOC"`)
	otLocShortJSON   = []byte(`"LOC"`)
	otLimitStr       = "LIMIT"
	otMocStr         = "MARKET_ON_CLOSE"
	otLocStr         = "LIMIT_ON_CLOSE"
)

//
type OrderType uint8

//
func OrderTypeFromString(s string) OrderType {
	switch strings.ToUpper(s) {
	case otLimitStr:
		return OrderTypeLimit
	case otMocStr:
		return OrderTypeMarketOnClose
	case otLocStr:
		return OrderTypeLimitOnClose
	default:
		return 0
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ot *OrderType) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, otLimitShortJSON) {
		*ot = OrderTypeLimit
	} else if bytes.Equal(data, otMocShortJSON) {
		*ot = OrderTypeMarketOnClose
	} else if bytes.Equal(data, otLocShortJSON) {
		*ot = OrderTypeLimitOnClose
	} else if bytes.Equal(data, otLimitJSON) {
		*ot = OrderTypeLimit
	} else if bytes.Equal(data, otMocJSON) {
		*ot = OrderTypeMarketOnClose
	} else if bytes.Equal(data, otLocJSON) {
		*ot = OrderTypeLimitOnClose
	} else {
		return errors.New("invalid BetStatus value in unmarshal - " + string(data))
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (ot OrderType) MarshalJSON() ([]byte, error) {
	switch ot {
	case OrderTypeLimit:
		return otLimitJSON, nil
	case OrderTypeMarketOnClose:
		return otMocJSON, nil
	case OrderTypeLimitOnClose:
		return otLocJSON, nil
	default:
		return nil, errors.New("invalid PersistenceType value")
	}
}

func (ot OrderType) String() string {
	if ot == OrderTypeLimit {
		return otLimitStr
	} else if ot == OrderTypeLimitOnClose {
		return otLocStr
	} else if ot == OrderTypeMarketOnClose {
		return otMocStr
	} else {
		return ""
	}
}

//
type RunnerStatus uint8

//
const (
	RunnerStatusActive  RunnerStatus = 1
	RunnerStatusRemoved RunnerStatus = 2
	RunnerStatusWinner  RunnerStatus = 3
	RunnerStatusPlaced  RunnerStatus = 4
	RunnerStatusLoser   RunnerStatus = 5
	RunnerStatusHidden  RunnerStatus = 6
)

var (
	runnerStatusActiveJSON  = []byte(`"ACTIVE"`)
	runnerStatusRemovedJSON = []byte(`"REMOVED"`)
	runnerStatusWinnerJSON  = []byte(`"WINNER"`)
	runnerStatusPlacedJSON  = []byte(`"PLACED"`)
	runnerStatusLoserJSON   = []byte(`"LOSER"`)
	runnerStatusHiddenJSON  = []byte(`"HIDDEN"`)
	runnerStatusActiveStr   = "ACTIVE"
	runnerStatusRemovedStr  = "REMOVED"
	runnerStatusWinnerStr   = "WINNER"
	runnerStatusPlacedStr   = "PLACED"
	runnerStatusLoserStr    = "LOSER"
	runnerStatusHiddenStr   = "HIDDEN"
)

// UnmarshalJSON implements the json.Unmarshaler interface.
func (rs *RunnerStatus) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, runnerStatusActiveJSON) {
		*rs = RunnerStatusActive
	} else if bytes.Equal(data, runnerStatusRemovedJSON) {
		*rs = RunnerStatusRemoved
	} else if bytes.Equal(data, runnerStatusWinnerJSON) {
		*rs = RunnerStatusWinner
	} else if bytes.Equal(data, runnerStatusPlacedJSON) {
		*rs = RunnerStatusPlaced
	} else if bytes.Equal(data, runnerStatusLoserJSON) {
		*rs = RunnerStatusLoser
	} else if bytes.Equal(data, runnerStatusHiddenJSON) {
		*rs = RunnerStatusHidden
	} else {
		return errors.New("invalid RunnerStatus value in unmarshal - " + string(data))
	}

	return nil
}

//
func (rs RunnerStatus) String() string {
	switch rs {
	case RunnerStatusActive:
		return runnerStatusActiveStr
	case RunnerStatusRemoved:
		return runnerStatusRemovedStr
	case RunnerStatusWinner:
		return runnerStatusWinnerStr
	case RunnerStatusPlaced:
		return runnerStatusPlacedStr
	case RunnerStatusLoser:
		return runnerStatusLoserStr
	case RunnerStatusHidden:
		return runnerStatusHiddenStr
	default:
		return ""
	}
}

//
type MarketStatus uint8

//
const (
	MarketStatusPending   MarketStatus = 0
	MarketStatusInActive  MarketStatus = 1
	MarketStatusOpen      MarketStatus = 2
	MarketStatusSuspended MarketStatus = 3
	MarketStatusClosed    MarketStatus = 4
)

var (
	marketStatusInActiveJSON  = []byte(`"INACTIVE"`)
	marketStatusOpenJSON      = []byte(`"OPEN"`)
	marketStatusSuspendedJSON = []byte(`"SUSPENDED"`)
	marketStatusClosedJSON    = []byte(`"CLOSED"`)
	marketStatusInActiveStr   = "INACTIVE"
	marketStatusOpenStr       = "OPEN"
	marketStatusSuspendedStr  = "SUSPENDED"
	marketStatusClosedStr     = "CLOSED"
)

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ms *MarketStatus) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, marketStatusOpenJSON) {
		*ms = MarketStatusOpen
	} else if bytes.Equal(data, marketStatusInActiveJSON) {
		*ms = MarketStatusInActive
	} else if bytes.Equal(data, marketStatusSuspendedJSON) {
		*ms = MarketStatusSuspended
	} else if bytes.Equal(data, marketStatusClosedJSON) {
		*ms = MarketStatusClosed
	} else {
		return errors.New("invalid MarketStatus value in unmarshal - " + string(data))
	}

	return nil
}

//
func (ms MarketStatus) String() string {
	switch ms {
	case MarketStatusOpen:
		return marketStatusOpenStr
	case MarketStatusInActive:
		return marketStatusInActiveStr
	case MarketStatusSuspended:
		return marketStatusSuspendedStr
	case MarketStatusClosed:
		return marketStatusClosedStr
	default:
		return ""
	}
}

//
type SubStatus uint8

//
const (
	SubStatusAll         SubStatus = 1
	SubStatusActivated   SubStatus = 2
	SubStatusExpired     SubStatus = 3
	SubStatusCancelled   SubStatus = 4
	SubStatusUnactivated SubStatus = 5
)

var (
	subStatusAllJSON         = []byte(`"ALL"`)
	subStatusActivatedJSON   = []byte(`"ACTIVATED"`)
	subStatusExpiredJSON     = []byte(`"EXPIRED"`)
	subStatusCancelledJSON   = []byte(`"CANCELLED"`)
	subStatusUnactivatedJSON = []byte(`"UNACTIVATED"`)
)

//
func (ss *SubStatus) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, subStatusActivatedJSON) {
		*ss = SubStatusActivated
	} else if bytes.Equal(data, subStatusUnactivatedJSON) {
		*ss = SubStatusUnactivated
	} else if bytes.Equal(data, subStatusCancelledJSON) {
		*ss = SubStatusCancelled
	} else if bytes.Equal(data, subStatusExpiredJSON) {
		*ss = SubStatusExpired
	} else if bytes.Equal(data, subStatusAllJSON) {
		*ss = SubStatusAll
	} else {
		return errors.New("invalid SubStatus value in unmarshal - " + string(data))
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (ss SubStatus) MarshalJSON() ([]byte, error) {
	switch ss {
	case SubStatusActivated:
		return subStatusActivatedJSON, nil
	case SubStatusExpired:
		return subStatusExpiredJSON, nil
	case SubStatusCancelled:
		return subStatusCancelledJSON, nil
	case SubStatusUnactivated:
		return subStatusUnactivatedJSON, nil
	default:
		return nil, errors.New("invalid SubStatus value")
	}
}

//
func (ss SubStatus) String() string {
	switch ss {
	case SubStatusActivated:
		return "Activated"
	case SubStatusExpired:
		return "Expired"
	case SubStatusCancelled:
		return "Cancelled"
	case SubStatusUnactivated:
		return "Unactivated"
	case SubStatusAll:
		return "All"
	default:
		return ""
	}
}

//
type ChangeType uint8

//
const (
	ChangeTypeSubImage   ChangeType = 1
	ChangeTypeResubDelta ChangeType = 2
	ChangeTypeHeartbeat  ChangeType = 3
)

var (
	changeTypeSubImageJSON   = []byte(`"SUB_IMAGE"`)
	changeTypeResubDeltaJSON = []byte(`"RESUB_DELTA"`)
	changeTypeHeartbeat      = []byte(`"HEARTBEAT"`)
)

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ct *ChangeType) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, changeTypeSubImageJSON) {
		*ct = ChangeTypeSubImage
	} else if bytes.Equal(data, changeTypeHeartbeat) {
		*ct = ChangeTypeHeartbeat
	} else if bytes.Equal(data, changeTypeResubDeltaJSON) {
		*ct = ChangeTypeResubDelta
	} else {
		return errors.New("invalid RunnerStatus value in unmarshal - " + string(data))
	}

	return nil
}

//
func (ct ChangeType) String() string {
	switch ct {
	case ChangeTypeSubImage:
		return "SUB_IMAGE"
	case ChangeTypeResubDelta:
		return "RESUB_DELTA"
	case ChangeTypeHeartbeat:
		return "HEARTBEAT"
	default:
		return ""
	}
}

//
type GrantType uint8

//
const (
	GrantTypeAuthorizationCode GrantType = 1
	GrantTypeRefreshToken      GrantType = 2
)

var (
	grantTypeAuthorizationCodeJSON = []byte(`"AUTHORIZATION_CODE"`)
	grantTypeRefreshTokenJSON      = []byte(`"REFRESH_TOKEN"`)
)

// MarshalJSON implements the json.Marshaler interface.
func (gt GrantType) String() string {
	switch gt {
	case GrantTypeAuthorizationCode:
		return string(grantTypeAuthorizationCodeJSON[1:19])
	case GrantTypeRefreshToken:
		return string(grantTypeRefreshTokenJSON[1:14])
	default:
		return ""
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (gt GrantType) MarshalJSON() ([]byte, error) {
	switch gt {
	case GrantTypeAuthorizationCode:
		return grantTypeAuthorizationCodeJSON, nil
	case GrantTypeRefreshToken:
		return grantTypeRefreshTokenJSON, nil
	default:
		return nil, errors.New("invalid GrantType value")
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (gt *GrantType) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, grantTypeAuthorizationCodeJSON) {
		*gt = GrantTypeAuthorizationCode
	} else if bytes.Equal(data, grantTypeRefreshTokenJSON) {
		*gt = GrantTypeRefreshToken
	} else {
		return errors.New("invalid GrantType value in unmarshal - " + string(data))
	}

	return nil
}

//
type TokenType uint8

//
const (
	TokenTypeBearer TokenType = 1
)

var (
	tokenTypeBearerJSON = []byte(`"BEARER"`)
)

// MarshalJSON implements the json.Marshaler interface.
func (tt TokenType) String() string {
	switch tt {
	case TokenTypeBearer:
		return string(tokenTypeBearerJSON[1:7])
	default:
		return ""
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (tt TokenType) MarshalJSON() ([]byte, error) {
	switch tt {
	case TokenTypeBearer:
		return tokenTypeBearerJSON, nil
	default:
		return nil, errors.New("invalid TokenType value")
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (tt *TokenType) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, tokenTypeBearerJSON) {
		*tt = TokenTypeBearer
	} else {
		return errors.New("invalid TokenType value in unmarshal - " + string(data))
	}

	return nil
}

//
type SegmentType uint8

//
const (
	SegmentTypeSegStart SegmentType = 1
	SegmentTypeSeg      SegmentType = 2
	SegmentTypeSegEnd   SegmentType = 3
)

var (
	segmentTypeSegStartJSON = []byte(`"SEG_START"`)
	segmentTypeSegJSON      = []byte(`"SEG"`)
	segmentTypeSegEndJSON   = []byte(`"SEG_END"`)
)

// MarshalJSON implements the json.Marshaler interface.
func (st SegmentType) String() string {
	switch st {
	case SegmentTypeSegStart:
		return string(segmentTypeSegStartJSON[1:10])
	case SegmentTypeSeg:
		return string(segmentTypeSegJSON[1:4])
	case SegmentTypeSegEnd:
		return string(segmentTypeSegEndJSON[1:8])
	default:
		return ""
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (st SegmentType) MarshalJSON() ([]byte, error) {
	switch st {
	case SegmentTypeSegStart:
		return segmentTypeSegStartJSON, nil
	case SegmentTypeSeg:
		return segmentTypeSegJSON, nil
	case SegmentTypeSegEnd:
		return segmentTypeSegEndJSON, nil
	default:
		return nil, nil
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (st *SegmentType) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, segmentTypeSegJSON) {
		*st = SegmentTypeSeg
	} else if bytes.Equal(data, segmentTypeSegStartJSON) {
		*st = SegmentTypeSegStart
	} else if bytes.Equal(data, segmentTypeSegEndJSON) {
		*st = SegmentTypeSegEnd
	} else {
		return errors.New("invalid SegmentType value in unmarshal - " + string(data))
	}

	return nil
}

//
const (
	BetTargetTypeBackerProfit BetTargetType = 1
	BetTargetTypePayout       BetTargetType = 2
)

var (
	bttBackerProfitJSON = []byte(`"BACKERS_PROFIT"`)
	bttPayoutJSON       = []byte(`"PAYOUT"`)
	bttBackerProfitStr  = "BACKERS_PROFIT"
	bttPayoutStr        = "PAYOUT"
)

//
type BetTargetType uint8

//
func BetTargetTypeFromString(s string) BetTargetType {
	switch strings.ToUpper(s) {
	case bttBackerProfitStr:
		return BetTargetTypeBackerProfit
	case bttPayoutStr:
		return BetTargetTypePayout
	default:
		panic("invalid BetTargetType string")
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (btt *BetTargetType) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, bttBackerProfitJSON) {
		*btt = BetTargetTypeBackerProfit
	} else if bytes.Equal(data, bttPayoutJSON) {
		*btt = BetTargetTypePayout
	} else {
		return errors.New("invalid BetTargetType value in unmarshal - " + string(data))
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (btt BetTargetType) MarshalJSON() ([]byte, error) {
	switch btt {
	case BetTargetTypeBackerProfit:
		return bttBackerProfitJSON, nil
	case BetTargetTypePayout:
		return bttPayoutJSON, nil
	default:
		return nil, errors.New("invalid BetTargetType value in marshal")
	}
}

func (btt BetTargetType) String() string {
	switch btt {
	case BetTargetTypeBackerProfit:
		return bttBackerProfitStr
	case BetTargetTypePayout:
		return bttPayoutStr
	default:
		return ""
	}
}

//
const (
	TimeInForceFillOrKill TimeInForce = 1
)

var (
	tifFillOrKillJSON = []byte(`"FILL_OR_KILL"`)
	tifFillOrKillStr  = "FILL_OR_KILL"
)

//
type TimeInForce uint8

//
func TimeInForceFromString(s string) TimeInForce {
	switch strings.ToUpper(s) {
	case tifFillOrKillStr:
		return TimeInForceFillOrKill
	default:
		panic("invalid TimeInForce string")
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (tif *TimeInForce) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, tifFillOrKillJSON) {
		*tif = TimeInForceFillOrKill
	} else {
		return errors.New("invalid TimeInForce value in unmarshal - " + string(data))
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (tif TimeInForce) MarshalJSON() ([]byte, error) {
	switch tif {
	case TimeInForceFillOrKill:
		return tifFillOrKillJSON, nil
	default:
		return nil, errors.New("invalid TimeInForce value in marshal")
	}
}

func (tif TimeInForce) String() string {
	switch tif {
	case TimeInForceFillOrKill:
		return tifFillOrKillStr
	default:
		return ""
	}
}
