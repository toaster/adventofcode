package orbit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/toaster/advent_of_code/2019/06/1/orbit"
)

func TestCount(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"example": {
			input: "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L",
			want:  42,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, orbit.Count(tt.input))
		})
	}
}
