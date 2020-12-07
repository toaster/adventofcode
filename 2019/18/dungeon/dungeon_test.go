package dungeon_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/toaster/advent_of_code/2019/18/dungeon"
)

func TestMap_MinimalStepsToCollectAllKeys(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"example 1": {
			input: "#########\n" +
				"#b.A.@.a#\n" +
				"#########\n",
			want: 8,
		},
		"example 2": {
			input: "########################\n" +
				"#f.D.E.e.C.b.A.@.a.B.c.#\n" +
				"######################.#\n" +
				"#d.....................#\n" +
				"########################\n",
			want: 86,
		},
		"example 3": {
			input: "########################\n" +
				"#...............b.C.D.f#\n" +
				"#.######################\n" +
				"#.....@.a.B.c.d.A.e.F.g#\n" +
				"########################\n",
			want: 132,
		},
		"example 4": {
			input: "#################\n" +
				"#i.G..c...e..H.p#\n" +
				"########.########\n" +
				"#j.A..b...f..D.o#\n" +
				"########@########\n" +
				"#k.E..a...g..B.n#\n" +
				"########.########\n" +
				"#l.F..d...h..C.m#\n" +
				"#################\n",
			want: 136,
		},
		"example 5": {
			input: "########################\n" +
				"#@..............ac.GI.b#\n" +
				"###d#e#f################\n" +
				"###A#B#C################\n" +
				"###g#h#i################\n" +
				"########################\n",
			want: 81,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			m := dungeon.Parse(tt.input)
			assert.Equal(t, tt.want, m.MinimalStepsToCollectAllKeys())
		})
	}
}
