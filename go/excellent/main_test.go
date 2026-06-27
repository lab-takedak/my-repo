package main

import "testing"

#test
func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "Even" {
		t.Errorf("expected: even, actual: '%s'", result)
	}
}
