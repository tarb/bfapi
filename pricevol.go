package bfapi

import (
	json "encoding/json"
	"fmt"
	"strconv"
)

//
type PriceVol float64

//
func (p PriceVol) String() string {
	s := strconv.FormatFloat(float64(p), 'f', -1, 64)

	switch s {
	case "+Inf":
		return s[1:]
	default:
		return s
	}
}

//
func (p *PriceVol) UnmarshalJSON(b []byte) error {
	// handles all correct float64 values
	var f float64

	err := json.Unmarshal(b, &f)
	if err == nil {
		*p = PriceVol(f)
		return nil
	}

	// handles string values
	var s string
	err = json.Unmarshal(b, &s)

	if err != nil {
		return fmt.Errorf("could not unmarshall PriceVol into float64 or string: %v", err)
	}

	switch s {
	case "Infinity":
		*p = PriceVol(-1)
		// *p = PriceVol(math.Inf(+1))
	case "null", "Null", "NULL":
		fallthrough
	case "NaN":
		*p = PriceVol(-1)
		// *p = PriceVol(math.NaN())
	default:
		return fmt.Errorf("could not unmarshall PriceVol, string value wanted (inf, null, nan) received (%v)", s)
	}

	return nil
}
