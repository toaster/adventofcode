package icc_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/toaster/advent_of_code/2019/05/2/icc"
)

func TestICC_Load(t *testing.T) {
	tests := map[string]struct {
		program string
		in      string
		wantOut string
	}{
		"example 1 #1": {
			program: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9",
			in:      "0\n",
			wantOut: "> 0\n",
		},
		"example 1 #2": {
			program: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9",
			in:      "1\n",
			wantOut: "> 1\n",
		},
		"example 2 #1": {
			program: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1",
			in:      "0\n",
			wantOut: "> 0\n",
		},
		"example 2 #2": {
			program: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1",
			in:      "-1\n",
			wantOut: "> 1\n",
		},
		"example 3 #1": {
			program: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			in:      "7\n",
			wantOut: "> 999\n",
		},
		"example 3 #2": {
			program: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			in:      "8\n",
			wantOut: "> 1000\n",
		},
		"example 3 #3": {
			program: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
			in:      "9\n",
			wantOut: "> 1001\n",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			out := bytes.NewBuffer([]byte{})
			c := icc.New(strings.NewReader(tt.in), out)
			c.Load(tt.program)
			c.Run()
			assert.Equal(t, tt.wantOut, out.String())
		})
	}
}
