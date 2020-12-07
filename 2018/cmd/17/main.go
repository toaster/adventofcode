package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type pos struct {
	x, y int
}

type ground int

const (
	sand ground = iota
	moist
	wet
	clay
	spring
)

type scan struct {
	slice                          map[pos]ground
	width, height, offset, voffset int
}

func (g ground) c() byte {
	switch g {
	case moist:
		return '|'
	case wet:
		return '~'
	case clay:
		return '#'
	case spring:
		return '+'
	default:
		return '.'
	}
}

func (s *scan) set(c pos, g ground) {
	if c.x <= s.offset {
		a := s.offset - c.x + 1
		s.width += a
		s.offset -= a
	} else if c.x >= s.offset+s.width-1 {
		s.width += c.x - s.offset - s.width + 2
	}
	if c.y > s.height-1 {
		s.height = c.y + 1
	}
	if s.voffset == 0 || c.y < s.voffset {
		s.voffset = c.y
	}
	s.slice[c] = clay
}

func p(x, y int) pos {
	return pos{x, y}
}

func main() {
	inputFile := os.Args[1]
	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	verticalRegex := regexp.MustCompile("x=(\\d+), y=(\\d+)..(\\d+)")
	horizontalRegex := regexp.MustCompile("y=(\\d+), x=(\\d+)..(\\d+)")

	s := scan{
		slice:  map[pos]ground{p(500, 0): spring},
		offset: 499,
		width:  3,
		height: 1,
	}
	lines := strings.Split(strings.TrimSuffix(string(b), "\n"), "\n")
	for _, l := range lines {
		m := verticalRegex.FindStringSubmatch(l)
		if m != nil {
			x := atoi(m[1])
			for y := atoi(m[2]); y <= atoi(m[3]); y++ {
				s.set(p(x, y), clay)
			}
		} else {
			m := horizontalRegex.FindStringSubmatch(l)
			y := atoi(m[1])
			for x := atoi(m[2]); x <= atoi(m[3]); x++ {
				s.set(p(x, y), clay)
			}
		}
	}

	printMap(&s)

	process(&s, p(500, 1))

	printMap(&s)
}

func printMap(s *scan) {
	reachableCount := 1 - s.voffset
	storedCount := 0
	for y := 0; y < s.height; y++ {
		for x := s.offset; x < s.offset+s.width; x++ {
			g := s.slice[p(x, y)]
			fmt.Print(string(g.c()))
			if g == wet || g == moist {
				reachableCount++
			}
			if g == wet {
				storedCount++
			}
		}
		fmt.Println("")
	}
	fmt.Println("Reachable:", reachableCount, "Stored:", storedCount)
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func process(s *scan, c pos) bool {
	fmt.Println("p-down:", c)
	// time.Sleep(200 * time.Millisecond)
	// printMap(s)
	if processDownOf(s, c) {
		blocked := scanLeftOf(s, c)
		blocked = scanRightOf(s, c) && blocked
		if blocked {
			floodLeft(s, c)
			floodRight(s, c)
			s.slice[c] = wet
			return true
		}
	}
	return false
}

func processDownOf(s *scan, c pos) bool {
	s.slice[c] = moist
	down := p(c.x, c.y+1)
	if down.y == s.height {
		return false
	}
	if s.slice[down] == clay || s.slice[down] == wet {
		return true
	}
	return process(s, down)
}

func scanLeftOf(s *scan, c pos) bool {
	left := p(c.x-1, c.y)
	g := s.slice[left]
	if g == sand {
		s.slice[left] = moist
		if processDownOf(s, left) {
			return scanLeftOf(s, left)
		}
		fmt.Println("free down of left of:", c)
	}
	fmt.Println("no sand left of:", c, g)
	return g == clay
}

func scanRightOf(s *scan, c pos) bool {
	right := p(c.x+1, c.y)
	g := s.slice[right]
	if g == sand {
		s.slice[right] = moist
		if processDownOf(s, right) {
			return scanRightOf(s, right)
		}
	}
	return g == clay
}

func floodLeft(s *scan, c pos) {
	if left := p(c.x-1, c.y); s.slice[left] == moist {
		floodLeft(s, left)
	}
	s.slice[c] = wet
}

func floodRight(s *scan, c pos) {
	if right := p(c.x+1, c.y); s.slice[right] == moist {
		floodRight(s, right)
	}
	s.slice[c] = wet
}
