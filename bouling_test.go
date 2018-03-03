package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// import "testing"

func TestGetScore(t *testing.T) {
	tests := []struct {
		throws string
		result int
	}{
		{"| 1 4 | 4 5 | 6 / | 5 / | X | 0 1 | 7 / | 6 / | X | 2 / 6 |", 133},
		{"| X | X | X | X | X | X | X | X | X | X X X |", 300},
		{"| X | 1 0 | 6 / | X | X | X | X | X | X | X X X |", 242},
		{"| X | 1 0 | 6 / | X | X | X | X | X | X | 1 2 |", 189},
		{"| 1 1 | 1 1 | 1 1 | 1 1 | 1 1| 1 1| 1 1| 1 1 | 1 1 | 1 1 |", 20},
	}

	for _, tc := range tests {
		score := getScore(tc.throws)
		assert.Equal(t, tc.result, score)
	}
}

func TestGetFrameScore(t *testing.T) {
	tests := []struct {
		frame  string
		result int
	}{
		{"XX", 20},
		{"6/", 10},
		{"12", 3},
	}

	for _, tc := range tests {
		score := getFrameScore(tc.frame)
		assert.Equal(t, tc.result, score)
	}
}
