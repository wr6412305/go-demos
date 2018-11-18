package math

import "testing"

func TestMaxPositive(t *testing.T) {
	if Max(1, 2, 3) != 3 {
		t.Error("Error: Max(1, 2, 3) != 3")
	}
}

func TestMaxNegative(t *testing.T) {
	if Max(-1, -2-3) != -1 {
		t.Error("Error: Max(-1, -2, -3) != -1")
	}
}
