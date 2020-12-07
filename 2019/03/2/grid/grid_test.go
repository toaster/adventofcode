package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_markAndSearch(t *testing.T) {
	tests := map[string]struct {
		input1   string
		input2   string
		wantDist int
	}{
		"sample 1": {
			input1:   "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			input2:   "U62,R66,U55,R34,D71,R55,D58,R83",
			wantDist: 610,
		},
		"sample 2": {
			input1:   "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			input2:   "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			wantDist: 410,
		},
		// TODO: Add test cases.
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			g := map[Pos]int{}
			Mark(g, tt.input1)
			assert.Equal(t, tt.wantDist, Search(g, tt.input2))
		})
	}
}
