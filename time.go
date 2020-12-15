package bfapi

import (
	"encoding/json"
	"time"
)

//
type Time uint64

//
func (t Time) ToStdTime() time.Time { return time.Unix(0, int64(t*1e6)) }

//
func FromStdTime(stdT time.Time) Time {
	var t Time

	if stdT != (time.Time{}) {
		t = Time(stdT.UnixNano() / 1e6)
	}

	return t
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a number in millisecond time
// from linux epoch or a quoted string in RFC 3339 format.
func (t *Time) UnmarshalJSON(data []byte) error {

	// if data is a millisecond value
	var i uint64
	if err := json.Unmarshal(data, &i); err == nil {
		*t = Time(i)
		return nil
	}

	// Ignore string null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}

	// Fractional seconds are handled implicitly by Parse.
	stdT, err := time.Parse(`"`+time.RFC3339+`"`, string(data))
	if err != nil {
		return err
	}

	if stdT != (time.Time{}) {
		*t = Time(stdT.UnixNano() / 1e6)
	}

	return err
}

// A Duration represents the elapsed time between two instants
// as an int64 millisecond count. This differs from Go's std
// library which is built off nanosecond time.
type Duration int64

//
func (d Duration) ToStdDuration() time.Duration { return time.Duration(d * 1e6) }
