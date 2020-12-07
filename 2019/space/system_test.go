package space_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/toaster/advent_of_code/2019/space"
)

func TestSystem_Simulate(t *testing.T) {
	ex1Input := "<x=-1, y=0, z=2>\n" +
		"<x=2, y=-10, z=-7>\n" +
		"<x=4, y=-8, z=8>\n" +
		"<x=3, y=5, z=-1>\n"
	ex2Input := "<x=-8, y=-10, z=0>\n" +
		"<x=5, y=5, z=10>\n" +
		"<x=2, y=-7, z=3>\n" +
		"<x=9, y=-8, z=-3>\n"
	tests := map[string]struct {
		input string
		steps int
		want  *space.System
	}{
		"example 1.0": {
			input: ex1Input,
			steps: 0,
			want: &space.System{
				Moons: []*space.Moon{
					{Pos: space.Vect{X: -1, Y: 0, Z: 2}, Vel: space.Vect{X: 0, Y: 0, Z: 0}},
					{Pos: space.Vect{X: 2, Y: -10, Z: -7}, Vel: space.Vect{X: 0, Y: 0, Z: 0}},
					{Pos: space.Vect{X: 4, Y: -8, Z: 8}, Vel: space.Vect{X: 0, Y: 0, Z: 0}},
					{Pos: space.Vect{X: 3, Y: 5, Z: -1}, Vel: space.Vect{X: 0, Y: 0, Z: 0}},
				},
			},
		},
		"example 1.1": {
			input: ex1Input,
			steps: 1,
			want: &space.System{
				Moons: []*space.Moon{
					{Pos: space.Vect{X: 2, Y: -1, Z: 1}, Vel: space.Vect{X: 3, Y: -1, Z: -1}},
					{Pos: space.Vect{X: 3, Y: -7, Z: -4}, Vel: space.Vect{X: 1, Y: 3, Z: 3}},
					{Pos: space.Vect{X: 1, Y: -7, Z: 5}, Vel: space.Vect{X: -3, Y: 1, Z: -3}},
					{Pos: space.Vect{X: 2, Y: 2, Z: 0}, Vel: space.Vect{X: -1, Y: -3, Z: 1}},
				},
			},
		},
		"example 1.2": {
			input: ex1Input,
			steps: 2,
			want: &space.System{
				Moons: []*space.Moon{
					{Pos: space.Vect{X: 5, Y: -3, Z: -1}, Vel: space.Vect{X: 3, Y: -2, Z: -2}},
					{Pos: space.Vect{X: 1, Y: -2, Z: 2}, Vel: space.Vect{X: -2, Y: 5, Z: 6}},
					{Pos: space.Vect{X: 1, Y: -4, Z: -1}, Vel: space.Vect{X: 0, Y: 3, Z: -6}},
					{Pos: space.Vect{X: 1, Y: -4, Z: 2}, Vel: space.Vect{X: -1, Y: -6, Z: 2}},
				},
			},
		},
		"example 1.3": {
			input: ex1Input,
			steps: 3,
			want: &space.System{
				Moons: []*space.Moon{
					{Pos: space.Vect{X: 5, Y: -6, Z: -1}, Vel: space.Vect{X: 0, Y: -3, Z: 0}},
					{Pos: space.Vect{X: 0, Y: 0, Z: 6}, Vel: space.Vect{X: -1, Y: 2, Z: 4}},
					{Pos: space.Vect{X: 2, Y: 1, Z: -5}, Vel: space.Vect{X: 1, Y: 5, Z: -4}},
					{Pos: space.Vect{X: 1, Y: -8, Z: 2}, Vel: space.Vect{X: 0, Y: -4, Z: 0}},
				},
			},
		},
		"example 1.4": {
			input: ex1Input,
			steps: 4,
			want: &space.System{
				Moons: []*space.Moon{
					{Pos: space.Vect{X: 2, Y: -8, Z: 0}, Vel: space.Vect{X: -3, Y: -2, Z: 1}},
					{Pos: space.Vect{X: 2, Y: 1, Z: 7}, Vel: space.Vect{X: 2, Y: 1, Z: 1}},
					{Pos: space.Vect{X: 2, Y: 3, Z: -6}, Vel: space.Vect{X: 0, Y: 2, Z: -1}},
					{Pos: space.Vect{X: 2, Y: -9, Z: 1}, Vel: space.Vect{X: 1, Y: -1, Z: -1}},
				},
			},
		},
		"example 1.5": {
			input: ex1Input,
			steps: 5,
			want: &space.System{
				Moons: []*space.Moon{
					{Pos: space.Vect{X: -1, Y: -9, Z: 2}, Vel: space.Vect{X: -3, Y: -1, Z: 2}},
					{Pos: space.Vect{X: 4, Y: 1, Z: 5}, Vel: space.Vect{X: 2, Y: 0, Z: -2}},
					{Pos: space.Vect{X: 2, Y: 2, Z: -4}, Vel: space.Vect{X: 0, Y: -1, Z: 2}},
					{Pos: space.Vect{X: 3, Y: -7, Z: -1}, Vel: space.Vect{X: 1, Y: 2, Z: -2}},
				},
			},
		},
		"example 1.6": {
			input: ex1Input,
			steps: 6,
			want: &space.System{
				Moons: []*space.Moon{
					{Pos: space.Vect{X: -1, Y: -7, Z: 3}, Vel: space.Vect{X: 0, Y: 2, Z: 1}},
					{Pos: space.Vect{X: 3, Y: 0, Z: 0}, Vel: space.Vect{X: -1, Y: -1, Z: -5}},
					{Pos: space.Vect{X: 3, Y: -2, Z: 1}, Vel: space.Vect{X: 1, Y: -4, Z: 5}},
					{Pos: space.Vect{X: 3, Y: -4, Z: -2}, Vel: space.Vect{X: 0, Y: 3, Z: -1}},
				},
			},
		},
		"example 1.7": {
			input: ex1Input,
			steps: 7,
			want: &space.System{
				Moons: []*space.Moon{
					{Pos: space.Vect{X: 2, Y: -2, Z: 1}, Vel: space.Vect{X: 3, Y: 5, Z: -2}},
					{Pos: space.Vect{X: 1, Y: -4, Z: -4}, Vel: space.Vect{X: -2, Y: -4, Z: -4}},
					{Pos: space.Vect{X: 3, Y: -7, Z: 5}, Vel: space.Vect{X: 0, Y: -5, Z: 4}},
					{Pos: space.Vect{X: 2, Y: 0, Z: 0}, Vel: space.Vect{X: -1, Y: 4, Z: 2}},
				},
			},
		},
		"example 1.8": {
			input: ex1Input,
			steps: 8,
			want: &space.System{
				Moons: []*space.Moon{
					{Pos: space.Vect{X: 5, Y: 2, Z: -2}, Vel: space.Vect{X: 3, Y: 4, Z: -3}},
					{Pos: space.Vect{X: 2, Y: -7, Z: -5}, Vel: space.Vect{X: 1, Y: -3, Z: -1}},
					{Pos: space.Vect{X: 0, Y: -9, Z: 6}, Vel: space.Vect{X: -3, Y: -2, Z: 1}},
					{Pos: space.Vect{X: 1, Y: 1, Z: 3}, Vel: space.Vect{X: -1, Y: 1, Z: 3}},
				},
			},
		},
		"example 1.9": {
			input: ex1Input,
			steps: 9,
			want: &space.System{
				Moons: []*space.Moon{
					{Pos: space.Vect{X: 5, Y: 3, Z: -4}, Vel: space.Vect{X: 0, Y: 1, Z: -2}},
					{Pos: space.Vect{X: 2, Y: -9, Z: -3}, Vel: space.Vect{X: 0, Y: -2, Z: 2}},
					{Pos: space.Vect{X: 0, Y: -8, Z: 4}, Vel: space.Vect{X: 0, Y: 1, Z: -2}},
					{Pos: space.Vect{X: 1, Y: 1, Z: 5}, Vel: space.Vect{X: 0, Y: 0, Z: 2}},
				},
			},
		},
		"example 1.10": {
			input: ex1Input,
			steps: 10,
			want: &space.System{
				Moons: []*space.Moon{
					{Pos: space.Vect{X: 2, Y: 1, Z: -3}, Vel: space.Vect{X: -3, Y: -2, Z: 1}},
					{Pos: space.Vect{X: 1, Y: -8, Z: 0}, Vel: space.Vect{X: -1, Y: 1, Z: 3}},
					{Pos: space.Vect{X: 3, Y: -6, Z: 1}, Vel: space.Vect{X: 3, Y: 2, Z: -3}},
					{Pos: space.Vect{X: 2, Y: 0, Z: 4}, Vel: space.Vect{X: 1, Y: -1, Z: -1}},
				},
			},
		},
		"example 2.0": {
			input: ex2Input,
			steps: 0,
			want: &space.System{
				Moons: []*space.Moon{
					{Pos: space.Vect{X: -8, Y: -10, Z: 0}, Vel: space.Vect{X: 0, Y: 0, Z: 0}},
					{Pos: space.Vect{X: 5, Y: 5, Z: 10}, Vel: space.Vect{X: 0, Y: 0, Z: 0}},
					{Pos: space.Vect{X: 2, Y: -7, Z: 3}, Vel: space.Vect{X: 0, Y: 0, Z: 0}},
					{Pos: space.Vect{X: 9, Y: -8, Z: -3}, Vel: space.Vect{X: 0, Y: 0, Z: 0}},
				},
			},
		},
		"example 2.10": {
			input: ex2Input,
			steps: 10,
			want: &space.System{
				Moons: []*space.Moon{
					{Pos: space.Vect{X: -9, Y: -10, Z: 1}, Vel: space.Vect{X: -2, Y: -2, Z: -1}},
					{Pos: space.Vect{X: 4, Y: 10, Z: 9}, Vel: space.Vect{X: -3, Y: 7, Z: -2}},
					{Pos: space.Vect{X: 8, Y: -10, Z: -3}, Vel: space.Vect{X: 5, Y: -1, Z: -2}},
					{Pos: space.Vect{X: 5, Y: -10, Z: 3}, Vel: space.Vect{X: 0, Y: -4, Z: 5}},
				},
			},
		},
		"example 2.20": {
			input: ex2Input,
			steps: 20,
			want: &space.System{
				Moons: []*space.Moon{
					{Pos: space.Vect{X: -10, Y: 3, Z: -4}, Vel: space.Vect{X: -5, Y: 2, Z: 0}},
					{Pos: space.Vect{X: 5, Y: -25, Z: 6}, Vel: space.Vect{X: 1, Y: 1, Z: -4}},
					{Pos: space.Vect{X: 13, Y: 1, Z: 1}, Vel: space.Vect{X: 5, Y: -2, Z: 2}},
					{Pos: space.Vect{X: 0, Y: 1, Z: 7}, Vel: space.Vect{X: -1, Y: -1, Z: 2}},
				},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s := space.ParseSystem(tt.input)
			s.Simulate(tt.steps)
			assert.Equal(t, tt.want, s)
		})
	}
}

func TestSystem_ComputePeriod(t *testing.T) {
	ex1Input := "<x=-1, y=0, z=2>\n" +
		"<x=2, y=-10, z=-7>\n" +
		"<x=4, y=-8, z=8>\n" +
		"<x=3, y=5, z=-1>\n"
	ex2Input := "<x=-8, y=-10, z=0>\n" +
		"<x=5, y=5, z=10>\n" +
		"<x=2, y=-7, z=3>\n" +
		"<x=9, y=-8, z=-3>\n"
	tests := map[string]struct {
		input string
		want  int
	}{
		"example 1": {
			input: ex1Input,
			want:  2772,
		},
		"example 2": {
			input: ex2Input,
			want:  4686774924,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s := space.ParseSystem(tt.input)
			assert.Equal(t, tt.want, s.ComputePeriod())
		})
	}
}
