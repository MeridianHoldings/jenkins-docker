package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	actual := Sum()
	expected := 2
	if actual != expected {
		t.Errorf("Test failed")
	}
}