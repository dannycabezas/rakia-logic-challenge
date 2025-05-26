package main

import (
	"testing"
)

func TestNumDecodings(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{"empty string", "", "0"},
		{"single zero", "0", "0"},
		{"zero six", "06", "0"},
		{"single digit one", "1", "1"},
		{"ten", "10", "1"},          // J
		{"twelve", "12", "2"},       // AB, L
		{"twenty seven", "27", "1"}, // BG (no), B G
		{"two two six", "226", "3"},
		{"one one one zero six", "11106", "2"},
		{"three zero one", "301", "0"}, // 30 is invalid
		{"one two three four five", "12345", "3"},
		{"long sequence of ones (100)", "1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111", "573147844013817084101"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualOutput := numDecodings(tc.input)
			actualOutputStr := actualOutput.String()

			if actualOutputStr != tc.expectedOutput {
				t.Errorf("Input: %s, Expected: %s, Got: %s", tc.input, tc.expectedOutput, actualOutputStr)
			}
		})
	}
}
