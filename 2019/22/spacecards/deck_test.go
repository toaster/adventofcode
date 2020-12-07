package spacecards_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/toaster/advent_of_code/2019/22/spacecards"
)

func TestDeck_Shuffle(t *testing.T) {
	tests := map[string]struct {
		size  int
		input string
		want  []int
	}{
		"example 1": {
			size: 10,
			input: "deal with increment 7\n" +
				"deal into new stack\n" +
				"deal into new stack\n",
			want: []int{0, 3, 6, 9, 2, 5, 8, 1, 4, 7},
		},
		"example 2": {
			size: 10,
			input: "cut 6\n" +
				"deal with increment 7\n" +
				"deal into new stack\n",
			want: []int{3, 0, 7, 4, 1, 8, 5, 2, 9, 6},
		},
		"example 3": {
			size: 10,
			input: "deal with increment 7\n" +
				"deal with increment 9\n" +
				"cut -2\n",
			want: []int{6, 3, 0, 7, 4, 1, 8, 5, 2, 9},
		},
		"example 4": {
			size: 10,
			input: "deal into new stack\n" +
				"cut -2\n" +
				"deal with increment 7\n" +
				"cut 8\n" +
				"cut -4\n" +
				"deal with increment 7\n" +
				"cut 3\n" +
				"deal with increment 9\n" +
				"deal with increment 3\n" +
				"cut -1\n",
			want: []int{9, 2, 5, 8, 1, 4, 7, 0, 3, 6},
		},
		// TODO: Add test cases.
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := spacecards.NewDeck(tt.size)
			d.Shuffle(tt.input)
			assert.Equal(t, tt.want, d.Cards)
		})
	}
}
