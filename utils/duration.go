package utils

import (
	"time"
)

// Duration custom duration struct
type Duration time.Duration

// MarshalKV marshal duration
func (d Duration) MarshalKV() (string, error) {
	return time.Duration(d).String(), nil
}

// UnmarshalKV unmarshal duration
func (d *Duration) UnmarshalKV(value string) error {
	pd, err := time.ParseDuration(value)
	if err != nil {
		return err
	}
	*d = Duration(pd)
	return nil
}

// String format to string
func (d Duration) String() string {
	return time.Duration(d).String()
}

// Duration get duration
func (d Duration) Duration() time.Duration {
	return time.Duration(d)
}

// Number get number
func (d Duration) Number() uint64 {
	dt := time.Duration(d)
	if dt.Hours() >= 1 {
		return uint64(dt.Hours())
	}
	if dt.Minutes() >= 1 {
		return uint64(dt.Minutes())
	}
	if dt.Seconds() >= 1 {
		return uint64(dt.Seconds())
	}
	if dt.Milliseconds() >= 1 {
		return uint64(dt.Milliseconds())
	}
	if dt.Microseconds() >= 1 {
		return uint64(dt.Microseconds())
	}
	if dt.Nanoseconds() >= 1 {
		return uint64(dt.Nanoseconds())
	}
	return 0
}

// Unit get unit
func (d Duration) Unit() string {
	dt := time.Duration(d)
	if dt.Hours() >= 1 {
		return "h"
	}
	if dt.Minutes() >= 1 {
		return "m"
	}
	if dt.Seconds() >= 1 {
		return "s"
	}
	if dt.Milliseconds() >= 1 {
		return "ms"
	}
	if dt.Microseconds() >= 1 {
		return "us"
	}
	if dt.Nanoseconds() >= 1 {
		return "ns"
	}
	return ""
}
