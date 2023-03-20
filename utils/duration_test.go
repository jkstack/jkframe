package utils

import (
	"testing"
	"time"
)

func TestDurationUnit(t *testing.T) {
	d := Duration(time.Hour)
	if d.Unit() != "h" {
		t.Fatalf("unexpected duration unit: %s", d.Unit())
	}
	d = Duration(time.Minute)
	if d.Unit() != "m" {
		t.Fatalf("unexpected duration unit: %s", d.Unit())
	}
	d = Duration(time.Second)
	if d.Unit() != "s" {
		t.Fatalf("unexpected duration unit: %s", d.Unit())
	}
	d = Duration(time.Millisecond)
	if d.Unit() != "ms" {
		t.Fatalf("unexpected duration unit: %s", d.Unit())
	}
	d = Duration(time.Microsecond)
	if d.Unit() != "us" {
		t.Fatalf("unexpected duration unit: %s", d.Unit())
	}
	d = Duration(time.Nanosecond)
	if d.Unit() != "ns" {
		t.Fatalf("unexpected duration unit: %s", d.Unit())
	}
}
