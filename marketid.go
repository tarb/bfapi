package bfapi

import "github.com/pkg/errors"

//
type MarketID uint64

// ErrMarketIDZero returned when serialising default value marketID
var ErrMarketIDZero = errors.New("MarketID is zero")

//
func NewMarketID(d []byte) MarketID {
	multiplier := 1
	num := 0

	for i := len(d) - 1; i > -1; i-- {
		if d[i] == '.' {
			continue
		} else if d[i] > 57 || d[i] < 48 {
			panic("bad input [" + string(d) + "] in marketID, found char " + string(rune(d[i])))
		}

		num += int(d[i]-48) * multiplier
		multiplier *= 10
	}

	return MarketID(num)
}

//
func (m MarketID) Encode() []byte {
	// count digits
	n := 0
	for i := m; i != 0; i /= 10 {
		n++
	}

	bs := make([]byte, n+1)
	for i := n; i > -1; i-- {
		bs[i] = byte(48 + m%10)
		m /= 10
	}
	bs[0], bs[1] = bs[1], '.'

	return bs
}

//
func (m MarketID) MarshalJSON() ([]byte, error) {
	if m == 0 {
		return nil, ErrMarketIDZero
	}

	// the same as encode, but wrapped with ""
	// count digits
	n := 0
	for i := m; i != 0; i /= 10 {
		n++
	}

	bs := make([]byte, n+3)
	for i := n + 1; i > 0; i-- {
		bs[i] = byte(48 + m%10)
		m /= 10
	}
	bs[0], bs[1], bs[2], bs[n+2] = '"', bs[2], '.', '"'

	return bs, nil
}

//
func (m *MarketID) UnmarshalJSON(data []byte) error {
	// d should be in the form "1.160296443" - we need to strip the "" from each end
	if len(data) < 3 || data[0] != '"' || data[len(data)-1] != '"' {
		return errors.Errorf("could not unmarshal marketID, json in bad format %s", string(data))
	}

	*m = NewMarketID(data[1 : len(data)-1])
	return nil
}

// MarketIDsDiff returns the difference between the 2 slices
func MarketIDsDiff(sa, sb []MarketID) []MarketID {
	var diff []MarketID

	for i := 0; i < 2; i++ {
		for _, s1 := range sa {
			found := false
			for _, s2 := range sb {
				if s1 == s2 {
					found = true
					break
				}
			}
			if !found {
				diff = append(diff, s1)
			}
		}

		if i == 0 {
			sa, sb = sb, sa
		}
	}

	return diff
}

//
type SelectionID uint64
