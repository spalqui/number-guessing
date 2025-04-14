package main

import "testing"

func TestIsGuessCorrect(t *testing.T) {
	tests := []struct {
		name          string
		guess         int
		correctNumber int
		attempt       int
		expected      bool
	}{
		{"Guess too low", 1, 50, 1, false},
		{"Guess too high", 100, 50, 1, false},
		{"Guess correct", 50, 50, 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isGuessCorrect(tt.guess, tt.correctNumber, tt.attempt)
			if result != tt.expected {
				t.Errorf("isGuessCorrect(%d, %d, %d) = %v; want %v", tt.guess, tt.correctNumber, tt.attempt, result, tt.expected)
			}
		})
	}
}

func TestMapDifficultyToMaxAttempts(t *testing.T) {
	tests := []struct {
		name     string
		level    int
		expected int
	}{
		{"Easy", 1, 10},
		{"Medium", 2, 5},
		{"Hard", 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mapDifficultyToMaxAttempts(tt.level)
			if result != tt.expected {
				t.Errorf("mapapDifficultyToMaxAttempts(%d) = %d; want %d", tt.level, result, tt.expected)
			}
		})
	}
}
