package utils

import "testing"

func TestVersionLess(t *testing.T) {
	a := Version([3]int{1, 0, 0})
	b := Version([3]int{1, 1, 1})
	if !a.Less(b) {
		t.Fatalf("unexpected less method")
	}
}

func TestVersionLessEqual(t *testing.T) {
	a := Version([3]int{1, 0, 0})
	b := Version([3]int{1, 1, 1})
	if !a.LessEqual(b) {
		t.Fatalf("unexpected less_equal method")
	}
	if !a.LessEqual(a) {
		t.Fatalf("unexpected less_equal method")
	}
}

func TestVersionGreater(t *testing.T) {
	a := Version([3]int{1, 0, 0})
	b := Version([3]int{1, 1, 1})
	if !b.Greater(a) {
		t.Fatalf("unexpected greater method")
	}
}

func TestVersionGreaterEqual(t *testing.T) {
	a := Version([3]int{1, 0, 0})
	b := Version([3]int{1, 1, 1})
	if !b.Greater(a) {
		t.Fatalf("unexpected greater_equal method")
	}
	if !b.GreaterEqual(b) {
		t.Fatalf("unexpected greater_equal method")
	}
}

func TestVersionEqual(t *testing.T) {
	a := Version([3]int{1, 0, 0})
	if !a.Equal(a) {
		t.Fatalf("unexpected equal method")
	}
}
