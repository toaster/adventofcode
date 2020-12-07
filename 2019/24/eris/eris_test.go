package eris_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/toaster/advent_of_code/2019/24/eris"
)

func TestMap_SimulateUntilRepeat(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"example": {
			input: "....#\n" +
				"#..#.\n" +
				"#..##\n" +
				"..#..\n" +
				"#....\n",
			want: 2129920,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			m := eris.New(tt.input)
			assert.Equal(t, tt.want, m.SimulateUntilRepeat())
		})
	}
}
