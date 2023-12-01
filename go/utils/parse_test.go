package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstAscii(t *testing.T) {
	testCases := []struct {
		input       string
		value       int
		shouldPanic bool
	}{
		{"A", 65, false},
		{"Ab", 65, false},
		{"", 0, true},
		{"ö", 0, true},
	}

	for _, tc := range testCases {
		if !tc.shouldPanic {
			got := FirstAsciiNr(tc.input)
			if got != tc.value {
				t.Errorf("FirstAscii(%q) => %d, want %d", tc.input, got, tc.value)
			}
		} else {
			assert.Panics(t, func() { _ = FirstAsciiNr(tc.input) },
				fmt.Sprintf("FirstAscii(%q) did not cause panic", tc.input))
		}
	}
}

func TestAtoi(t *testing.T) {
	testCases := []struct {
		input       string
		value       int
		shouldPanic bool
	}{
		{"1123", 1123, false},
		{"-12", -12, false},
		{"1.2", 1, true},
		{"", 0, true},
		{"ö", 0, true},
	}

	for _, tc := range testCases {
		if !tc.shouldPanic {
			got := Atoi(tc.input)
			if got != tc.value {
				t.Errorf("FirstAscii(%q) => %d, want %d", tc.input, got, tc.value)
			}
		} else {
			assert.Panics(t, func() { _ = Atoi(tc.input) },
				fmt.Sprintf("Input %q did not cause panic", tc.input))
		}
	}
}

func TestCut(t *testing.T) {
	testCases := []struct {
		input       string
		separator   string
		outLeft     string
		outRight    string
		shouldPanic bool
	}{
		{"a=b", "=", "a", "b", false},
		{"", "=", "", "", true},
		{"a=b=c", "=", "a", "b=c", false},
	}

	for _, tc := range testCases {
		if !tc.shouldPanic {
			gotLeft, gotRight := Cut(tc.input, tc.separator)
			if gotLeft != tc.outLeft || gotRight != tc.outRight {
				t.Errorf("Cut(%q, %q) => %q, %q, want %q, %q", tc.input, tc.separator,
					gotLeft, gotRight, tc.outLeft, tc.outRight)
			}
		} else {
			assert.Panics(t, func() { _, _ = Cut(tc.input, tc.separator) },
				fmt.Sprintf("Cut(%q, %q) did not cause panic", tc.input, tc.separator))
		}
	}
}
