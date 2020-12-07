package pluto_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/toaster/advent_of_code/2019/20/pluto"
)

func TestMaze_ShortestPath(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"example 1": {
			input: "         A\n" +
				"         A\n" +
				"  #######.#########\n" +
				"  #######.........#\n" +
				"  #######.#######.#\n" +
				"  #######.#######.#\n" +
				"  #######.#######.#\n" +
				"  #####  B    ###.#\n" +
				"BC...##  C    ###.#\n" +
				"  ##.##       ###.#\n" +
				"  ##...DE  F  ###.#\n" +
				"  #####    G  ###.#\n" +
				"  #########.#####.#\n" +
				"DE..#######...###.#\n" +
				"  #.#########.###.#\n" +
				"FG..#########.....#\n" +
				"  ###########.#####\n" +
				"             Z\n" +
				"             Z\n",
			want: 23,
		},
		"example 2": {
			input: "                   A\n" +
				"                   A\n" +
				"  #################.#############\n" +
				"  #.#...#...................#.#.#\n" +
				"  #.#.#.###.###.###.#########.#.#\n" +
				"  #.#.#.......#...#.....#.#.#...#\n" +
				"  #.#########.###.#####.#.#.###.#\n" +
				"  #.............#.#.....#.......#\n" +
				"  ###.###########.###.#####.#.#.#\n" +
				"  #.....#        A   C    #.#.#.#\n" +
				"  #######        S   P    #####.#\n" +
				"  #.#...#                 #......VT\n" +
				"  #.#.#.#                 #.#####\n" +
				"  #...#.#               YN....#.#\n" +
				"  #.###.#                 #####.#\n" +
				"DI....#.#                 #.....#\n" +
				"  #####.#                 #.###.#\n" +
				"ZZ......#               QG....#..AS\n" +
				"  ###.###                 #######\n" +
				"JO..#.#.#                 #.....#\n" +
				"  #.#.#.#                 ###.#.#\n" +
				"  #...#..DI             BU....#..LF\n" +
				"  #####.#                 #.#####\n" +
				"YN......#               VT..#....QG\n" +
				"  #.###.#                 #.###.#\n" +
				"  #.#...#                 #.....#\n" +
				"  ###.###    J L     J    #.#.###\n" +
				"  #.....#    O F     P    #.#...#\n" +
				"  #.###.#####.#.#####.#####.###.#\n" +
				"  #...#.#.#...#.....#.....#.#...#\n" +
				"  #.#####.###.###.#.#.#########.#\n" +
				"  #...#.#.....#...#.#.#.#.....#.#\n" +
				"  #.###.#####.###.###.#.#.#######\n" +
				"  #.#.........#...#.............#\n" +
				"  #########.###.###.#############\n" +
				"           B   J   C\n" +
				"           U   P   P\n",
			want: 58,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			m := pluto.Parse(tt.input)
			assert.Equal(t, tt.want, m.ShortestPath())
		})
	}
}
