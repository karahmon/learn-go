package main

import (
	"testing"
)

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-20", 100.0, 5.0, 20.0, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := divide(tt.dividend, tt.divisor)
			if tt.isErr {
				if err == nil {
					t.Error("Expected an error but didn't get one")
				}
			} else {
				if err != nil {
					t.Error("Did not expect an error but got one:", err.Error())
				}
				if got != tt.expected {
					t.Errorf("Expected %f but got %f", tt.expected, got)
				}
			}
		})
	}
}
