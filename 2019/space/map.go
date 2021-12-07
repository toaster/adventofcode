package space

import (
	"sort"
	"strings"

	math2019 "github.com/toaster/advent_of_code/2019/math"
	"github.com/toaster/advent_of_code/internal/math"
)

// Map shows asteroids in a space segment.
type Map struct {
	asteroids [][]bool
	width     int
	height    int
	lcm       int
}

// Point represents a point, probably on a map ;).
type Point struct {
	X int
	Y int
}

// ParseMap creates a Map out of an input string.
func ParseMap(input string) *Map {
	m := Map{}
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, square := range line {
			if x == 0 {
				m.asteroids = append(m.asteroids, []bool{})
			}
			m.asteroids[y] = append(m.asteroids[y], square == '#')
		}
	}
	m.width = len(m.asteroids[0])
	m.height = len(m.asteroids)

	var numbers []int
	for i := 0; i < m.width; i++ {
		numbers = append(numbers, i)
	}
	m.lcm = math2019.LCM(numbers)

	return &m
}

// MaxVisibleAsteroids computes the maximum amount of other asteroids visible from an asteroid.
func (m *Map) MaxVisibleAsteroids() (int, Point) {
	max := 0
	maxP := Point{}
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			if !m.asteroids[y][x] {
				continue
			}
			angles := m.scan(x, y)
			if len(angles) > max {
				max = len(angles)
				maxP = Point{x, y}
			}
		}
	}
	return max, maxP
}

// VaporizeAsteroids turns the giant laser at pos and vaporizes all asteroids. It returns the coordinates
// of the vaporized asteroids in order.
func (m *Map) VaporizeAsteroids(pos Point) []Point {
	asteroids := m.scan(pos.X, pos.Y)
	rays := sortableRays{}
	for point, asteroids := range asteroids {
		rays = append(rays, ray{point, asteroids})
	}
	sort.Sort(rays)
	var vaporized []Point
	for len(rays) > 0 {
		newRays := rays[:0]
		for _, ray := range rays {
			vaporized = append(vaporized, ray.asteroids[0])
			ray.asteroids = ray.asteroids[1:]
			if len(ray.asteroids) > 0 {
				newRays = append(newRays, ray)
			}
		}
		rays = newRays
	}
	return vaporized
}

type ray struct {
	normPoint Point
	asteroids []Point
}

type sortableRays []ray

var _ sort.Interface = sortableRays(nil)

func (s sortableRays) Len() int {
	return len(s)
}

func (s sortableRays) Less(i, j int) bool {
	r1 := s[i]
	r2 := s[j]
	if r1.normPoint.X == 0 {
		if r2.normPoint.X == 0 {
			return r2.normPoint.Y > r1.normPoint.Y
		}
		if r1.normPoint.Y < 0 {
			return true
		}
		return r2.normPoint.X < 0
	}
	if r1.normPoint.X > 0 {
		if r2.normPoint.X < 0 {
			return true
		}
		if r2.normPoint.X == 0 {
			return r2.normPoint.Y > 0
		}
		return r2.normPoint.Y > r1.normPoint.Y
	}
	return r2.normPoint.X < 0 && r2.normPoint.Y < r1.normPoint.Y
}

func (s sortableRays) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (m *Map) scan(x int, y int) map[Point][]Point {
	asteroids := map[Point][]Point{}
	for sy := 0; sy < m.height; sy++ {
		for sx := 0; sx < m.width; sx++ {
			if sx == x && sy == y {
				continue
			}
			if m.asteroids[sy][sx] {
				dx := sx - x
				dy := sy - y
				var np Point
				if dx == 0 {
					np = Point{X: dx, Y: dy / math.AbsInt(dy)}
				} else {
					f := m.lcm / math.AbsInt(dx)
					np = Point{X: f * dx, Y: f * dy}
				}
				if len(asteroids[np]) == 0 {
					asteroids[np] = []Point{{sx, sy}}
				} else if math.AbsInt(asteroids[np][0].X-x) < math.AbsInt(dx) {
					asteroids[np] = append(asteroids[np], Point{sx, sy})
				} else {
					asteroids[np] = append([]Point{{sx, sy}}, asteroids[np]...)
				}
			}
		}
	}
	return asteroids
}
