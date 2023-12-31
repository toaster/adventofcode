package space_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/toaster/advent_of_code/2019/space"
	"github.com/toaster/advent_of_code/internal/math"
)

func TestMap_MaxVisibleAsteroids(t *testing.T) {
	tests := map[string]struct {
		input     []string
		wantMax   int
		wantPoint math.Point2D
	}{
		"example 1": {
			input: []string{
				".#..#",
				".....",
				"#####",
				"....#",
				"...##",
			},
			wantMax:   8,
			wantPoint: math.Point2D{X: 3, Y: 4},
		},
		"example 2": {
			input: []string{
				"......#.#.",
				"#..#.#....",
				"..#######.",
				".#.#.###..",
				".#..#.....",
				"..#....#.#",
				"#..#....#.",
				".##.#..###",
				"##...#..#.",
				".#....####",
			},
			wantMax:   33,
			wantPoint: math.Point2D{X: 5, Y: 8},
		},
		"example 3": {
			input: []string{
				"#.#...#.#.",
				".###....#.",
				".#....#...",
				"##.#.#.#.#",
				"....#.#.#.",
				".##..###.#",
				"..#...##..",
				"..##....##",
				"......#...",
				".####.###.",
			},
			wantMax:   35,
			wantPoint: math.Point2D{X: 1, Y: 2},
		},
		"example 4": {
			input: []string{
				".#..#..###",
				"####.###.#",
				"....###.#.",
				"..###.##.#",
				"##.##.#.#.",
				"....###..#",
				"..#.#..#.#",
				"#..#.#.###",
				".##...##.#",
				".....#.#..",
			},
			wantMax:   41,
			wantPoint: math.Point2D{X: 6, Y: 3},
		},
		"example 5": {
			input: []string{
				".#..##.###...#######",
				"##.############..##.",
				".#.######.########.#",
				".###.#######.####.#.",
				"#####.##.#.##.###.##",
				"..#####..#.#########",
				"####################",
				"#.####....###.#.#.##",
				"##.#################",
				"#####.##.###..####..",
				"..######..##.#######",
				"####.##.####...##..#",
				".#####..#.######.###",
				"##...#.##########...",
				"#.##########.#######",
				".####.#.###.###.#.##",
				"....##.##.###..#####",
				".#.#.###########.###",
				"#.#.#.#####.####.###",
				"###.##.####.##.#..##",
			},
			wantMax:   210,
			wantPoint: math.Point2D{X: 11, Y: 13},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			m := space.ParseMap(strings.Join(tt.input, "\n"))
			count, point := m.MaxVisibleAsteroids()
			assert.Equal(t, tt.wantMax, count)
			assert.Equal(t, tt.wantPoint, point)
		})
	}
}

func TestMap_VaporizeAsteroids(t *testing.T) {
	tests := map[string]struct {
		input []string
		pos   math.Point2D
		want  []math.Point2D
	}{
		"example 1": {
			input: []string{
				".#....#####...#..",
				"##...##.#####..##",
				"##...#...#.#####.",
				"..#.....X...###..",
				"..#.#.....#....##",
			},
			pos: math.Point2D{X: 8, Y: 3},
			want: []math.Point2D{
				{8, 1},
				{9, 0},
				{9, 1},
				{10, 0},
				{9, 2},
				{11, 1},
				{12, 1},
				{11, 2},
				{15, 1},
				{12, 2},
				{13, 2},
				{14, 2},
				{15, 2},
				{12, 3},
				{16, 4},
				{15, 4},
				{10, 4},
				{4, 4},
				{2, 4},
				{2, 3},
				{0, 2},
				{1, 2},
				{0, 1},
				{1, 1},
				{5, 2},
				{1, 0},
				{5, 1},
				{6, 1},
				{6, 0},
				{7, 0},
				{8, 0},
				{10, 1},
				{14, 0},
				{16, 1},
				{13, 3},
				{14, 3},
			},
		},
		// "example 5": {
		// 	input: []string{
		// 		".#..##.###...#######",
		// 		"##.############..##.",
		// 		".#.######.########.#",
		// 		".###.#######.####.#.",
		// 		"#####.##.#.##.###.##",
		// 		"..#####..#.#########",
		// 		"####################",
		// 		"#.####....###.#.#.##",
		// 		"##.#################",
		// 		"#####.##.###..####..",
		// 		"..######..##.#######",
		// 		"####.##.####...##..#",
		// 		".#####..#.######.###",
		// 		"##...#.##########...",
		// 		"#.##########.#######",
		// 		".####.#.###.###.#.##",
		// 		"....##.##.###..#####",
		// 		".#.#.###########.###",
		// 		"#.#.#.#####.####.###",
		// 		"###.##.####.##.#..##",
		// 	},
		// 	want: 210,
		// },
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			m := space.ParseMap(strings.Join(tt.input, "\n"))
			assert.Equal(t, tt.want, m.VaporizeAsteroids(tt.pos))
		})
	}
}
