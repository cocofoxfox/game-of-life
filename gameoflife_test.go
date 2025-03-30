package main

import (
	"testing"
)

func TestNextGeneration(t *testing.T) {
	tests := []struct {
		name     string
		input    map[Cell]bool
		expected map[Cell]bool
	}{
		{
			name: "Stable 2x2 block remains the same",
			input: map[Cell]bool{
				{0, 0}: true, {0, 1}: true,
				{1, 0}: true, {1, 1}: true,
			},
			expected: map[Cell]bool{
				{0, 0}: true, {0, 1}: true,
				{1, 0}: true, {1, 1}: true,
			},
		},
		{
			name: "Glider",
			input: map[Cell]bool{
				{0, 1}: true,
				{1, 2}: true,
				{2, 0}: true,
				{2, 1}: true,
				{2, 2}: true,
			},
			expected: map[Cell]bool{
				{1, 0}: true,
				{2, 1}: true,
				{2, 2}: true,
				{1, 2}: true,
				{3, 1}: true,
			},
		},
		{
			name: "Large coordinates",
			input: map[Cell]bool{
				{-2000000000000, -2000000000000}: true,
				{-2000000000001, -2000000000001}: true,
				{-2000000000000, -2000000000001}: true,
			},
			// TODO: this still fails, after one iteration seems generate more elements, no time to finish
			expected: map[Cell]bool{
				{-2000000000000, -2000000000000}: true,
				{-2000000000001, -2000000000001}: true,
				{-2000000000000, -2000000000001}: true,
			},
		},
		{
			name: "Single cell dies",
			input: map[Cell]bool{
				{0, 0}: true,
			},
			expected: map[Cell]bool{},
		},
		{
			name:     "Empty input remains empty",
			input:    map[Cell]bool{},
			expected: map[Cell]bool{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nextGeneration(tt.input)
			if !equalSets(got, tt.expected) {
				t.Errorf("nextGeneration() = %v; want %v", got, tt.expected)
			}
		})
	}
}

func equalSets(a, b map[Cell]bool) bool {
	if len(a) != len(b) {
		return false
	}
	for cell := range a {
		if !b[cell] {
			return false
		}
	}
	return true
}
