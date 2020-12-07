package fft_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/toaster/advent_of_code/2019/16/fft"
)

func TestPerform(t *testing.T) {
	tests := map[string]struct {
		input      []int
		phaseCount int
		want       []int
	}{
		"example 1": {
			input:      []int{1, 2, 3, 4, 5, 6, 7, 8},
			phaseCount: 1,
			want:       []int{4, 8, 2, 2, 6, 1, 5, 8},
		},
		"example 2": {
			input:      []int{1, 2, 3, 4, 5, 6, 7, 8},
			phaseCount: 2,
			want:       []int{3, 4, 0, 4, 0, 4, 3, 8},
		},
		"example 3": {
			input:      []int{1, 2, 3, 4, 5, 6, 7, 8},
			phaseCount: 3,
			want:       []int{0, 3, 4, 1, 5, 5, 1, 8},
		},
		"example 4": {
			input:      []int{1, 2, 3, 4, 5, 6, 7, 8},
			phaseCount: 4,
			want:       []int{0, 1, 0, 2, 9, 4, 9, 8},
		},
		"example 5": {
			input:      []int{8, 0, 8, 7, 1, 2, 2, 4, 5, 8, 5, 9, 1, 4, 5, 4, 6, 6, 1, 9, 0, 8, 3, 2, 1, 8, 6, 4, 5, 5, 9, 5},
			phaseCount: 100,
			want:       []int{2, 4, 1, 7, 6, 1, 7, 6, 4, 8, 0, 9, 1, 9, 0, 4, 6, 1, 1, 4, 0, 3, 8, 7, 6, 3, 1, 9, 5, 5, 9, 5},
		},
		"example 6": {
			input:      []int{1, 9, 6, 1, 7, 8, 0, 4, 2, 0, 7, 2, 0, 2, 2, 0, 9, 1, 4, 4, 9, 1, 6, 0, 4, 4, 1, 8, 9, 9, 1, 7},
			phaseCount: 100,
			want:       []int{7, 3, 7, 4, 5, 4, 1, 8, 5, 5, 7, 2, 5, 7, 2, 5, 9, 1, 4, 9, 4, 6, 6, 5, 9, 9, 6, 3, 9, 9, 1, 7},
		},
		"example 7": {
			input:      []int{6, 9, 3, 1, 7, 1, 6, 3, 4, 9, 2, 9, 4, 8, 6, 0, 6, 3, 3, 5, 9, 9, 5, 9, 2, 4, 3, 1, 9, 8, 7, 3},
			phaseCount: 100,
			want:       []int{5, 2, 4, 3, 2, 1, 3, 3, 2, 9, 2, 9, 9, 8, 6, 0, 6, 8, 8, 0, 4, 9, 5, 9, 7, 4, 8, 6, 9, 8, 7, 3},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, fft.Perform(tt.input, tt.phaseCount))
		})
	}
}

func TestPerformHuge(t *testing.T) {
	tests := map[string]struct {
		input      []int
		phaseCount int
		want       []int
	}{
		"example 5": {
			input:      []int{0, 3, 0, 3, 6, 7, 3, 2, 5, 7, 7, 2, 1, 2, 9, 4, 4, 0, 6, 3, 4, 9, 1, 5, 6, 5, 4, 7, 4, 6, 6, 4},
			phaseCount: 100,
			want:       []int{8, 4, 4, 6, 2, 0, 2, 6},
		},
		"example 6": {
			input:      []int{0, 2, 9, 3, 5, 1, 0, 9, 6, 9, 9, 9, 4, 0, 8, 0, 7, 4, 0, 7, 5, 8, 5, 4, 4, 7, 0, 3, 4, 3, 2, 3},
			phaseCount: 100,
			want:       []int{7, 8, 7, 2, 5, 2, 7, 0},
		},
		"example 7": {
			input:      []int{0, 3, 0, 8, 1, 7, 7, 0, 8, 8, 4, 9, 2, 1, 9, 5, 9, 7, 3, 1, 1, 6, 5, 4, 4, 6, 8, 5, 0, 5, 1, 7},
			phaseCount: 100,
			want:       []int{5, 3, 5, 5, 3, 7, 3, 1},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, fft.PerformHuge(tt.input, tt.phaseCount))
		})
	}
}
