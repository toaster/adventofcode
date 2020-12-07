package icc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	icc2 "github.com/toaster/advent_of_code/2019/07/2/icc"
)

func TestICC_LoadAndRun(t *testing.T) {
	tests := map[string]struct {
		program string
		in      int
		wantOut int
	}{
		"example 1 #1": {
			program: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9",
			in:      0,
			wantOut: 0,
		},
		"example 1 #2": {
			program: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9",
			in:      1,
			wantOut: 1,
		},
		"example 2 #1": {
			program: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1",
			in:      0,
			wantOut: 0,
		},
		"example 2 #2": {
			program: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1",
			in:      -1,
			wantOut: 1,
		},
		"example 3 #1": {
			program: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			in:      7,
			wantOut: 999,
		},
		"example 3 #2": {
			program: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			in:      8,
			wantOut: 1000,
		},
		"example 3 #3": {
			program: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			in:      9,
			wantOut: 1001,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			in := make(chan int, 10)
			out := make(chan int, 10)
			c := icc2.New(in, out)
			c.Load(tt.program)
			in <- tt.in
			c.Run()
			assert.Equal(t, tt.wantOut, <-out)
		})
	}
}

func TestICC_ComputeOptimalThrusterConfig(t *testing.T) {
	tests := map[string]struct {
		program    string
		wantSignal int
		wantPhases []int
	}{
		"example 1": {
			program:    "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0",
			wantSignal: 43210,
			wantPhases: []int{4, 3, 2, 1, 0},
		},
		"example 2": {
			program:    "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0",
			wantSignal: 54321,
			wantPhases: []int{0, 1, 2, 3, 4},
		},
		"example 3": {
			program:    "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0",
			wantSignal: 65210,
			wantPhases: []int{1, 0, 4, 3, 2},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s, p := icc2.ComputeOptimalThrusterConfig(tt.program)
			assert.Equal(t, tt.wantSignal, s)
			assert.Equal(t, tt.wantPhases, p)
		})
	}
}

func TestICC_ComputeOptimalLoopedThrusterConfig(t *testing.T) {
	tests := map[string]struct {
		program    string
		wantSignal int
		wantPhases []int
	}{
		"example 1": {
			program:    "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5",
			wantSignal: 139629729,
			wantPhases: []int{9, 8, 7, 6, 5},
		},
		"example 2": {
			program:    "3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10",
			wantSignal: 18216,
			wantPhases: []int{9, 7, 8, 5, 6},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s, p := icc2.ComputeOptimalLoopedThrusterConfig(tt.program)
			assert.Equal(t, tt.wantSignal, s)
			assert.Equal(t, tt.wantPhases, p)
		})
	}
}
