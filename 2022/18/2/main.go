package main

import (
	"fmt"
	"math"

	"github.com/toaster/advent_of_code/internal/io"
	aocmath "github.com/toaster/advent_of_code/internal/math"
)

var maxIntRange = aocmath.Range{Start: math.MinInt64, End: math.MaxInt64}

func main() {
	points := map[aocmath.Point3D]bool{}
	max := aocmath.Point3D{X: math.MinInt64, Y: math.MinInt64, Z: math.MinInt64}
	min := aocmath.Point3D{X: math.MaxInt64, Y: math.MaxInt64, Z: math.MaxInt64}
	for _, line := range io.ReadLines() {
		coords := io.ParseInts(line, ",")
		p := aocmath.Point3D{X: coords[0], Y: coords[1], Z: coords[2]}
		points[p] = true
		if p.X < min.X {
			min.X = p.X
		}
		if p.Y < min.Y {
			min.Y = p.Y
		}
		if p.Z < min.Z {
			min.Z = p.Z
		}
		if p.X > max.X {
			max.X = p.X
		}
		if p.Y > max.Y {
			max.Y = p.Y
		}
		if p.Z > max.Z {
			max.Z = p.Z
		}
	}

	surface := 0
	outside := map[aocmath.Point3D]bool{}
	inside := map[aocmath.Point3D]bool{}
	for p := range points {
		for _, n := range p.Neighbours(maxIntRange, maxIntRange, maxIntRange) {
			if !points[n] && !inside[n] && !outside[n] {
				visited := map[aocmath.Point3D]bool{n: true}
				m := inside
				if grow(n, min, max, points, visited) {
					m = outside
				}
				for v := range visited {
					m[v] = true
				}
			}
			if outside[n] {
				surface++
			}
		}
	}
	fmt.Println(surface)
}

func grow(p aocmath.Point3D, min, max aocmath.Point3D, points, visited map[aocmath.Point3D]bool) bool {
	if p.X < min.X || p.Y < min.Y || p.Z < min.Z {
		return true
	}

	if p.X > max.X || p.Y > max.Y || p.Z > max.Z {
		return true
	}

	outside := false
	for _, n := range p.Neighbours(maxIntRange, maxIntRange, maxIntRange) {
		if !points[n] && !visited[n] {
			visited[n] = true
			outside = outside || grow(n, min, max, points, visited)
		}
	}
	return outside
}
