package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	lines := io.ReadLines()
	input := lines[:len(lines)-2]
	tiles, start := parseMap(input)
	directions := parseDirections(lines[len(lines)-1])

	cur := tiles[start]
	heading := 0
	for _, d := range directions {
		for i := 0; i < d.move; i++ {
			var next *tile
			var headingChange int
			switch heading {
			case 0:
				next = cur.right
				headingChange = cur.rightTurn
			case 1:
				next = cur.bottom
				headingChange = cur.bottomTurn
			case 2:
				next = cur.left
				headingChange = cur.leftTurn
			case 3:
				next = cur.top
				headingChange = cur.topTurn
			}
			if !next.open {
				break
			}
			cur = next
			fmt.Println("move", heading, "to", cur.pos, "and change heading to", heading+headingChange)
			heading += headingChange
		}
		heading = (heading + d.turn + 4) % 4
		fmt.Println("turn", d.turn, "to", heading)
	}
	fmt.Println(1000*cur.pos.Y + 4*cur.pos.X + heading)
}

var nets = map[int]map[math.Point2D]func(int, int, int) (math.Point2D, int){
	// example:
	4 + 32 + 64 + 128 + 4096 + 8192: {
		{2, 0}: func(offset int, heading int, size int) (math.Point2D, int) {
			switch heading {
			case 0:
				return math.Point2D{X: 4 * size, Y: 3*size - offset}, 2
			case 2:
				return math.Point2D{X: size + 1 + offset, Y: size + 1}, -1
			case 3:
				return math.Point2D{X: size - offset, Y: size + 1}, -2
			default:
				panic(fmt.Sprintf("example 2,0 unexpected heading %d (off %d, size %d)", heading, offset, size))
			}
		},
		{0, 1}: func(offset int, heading int, size int) (math.Point2D, int) {
			switch heading {
			case 1:
				return math.Point2D{X: 3*size - offset, Y: 3 * size}, 2
			case 2:
				return math.Point2D{X: 4*size - offset, Y: 3 * size}, 1
			case 3:
				return math.Point2D{X: 3*size - offset, Y: 1}, -2
			default:
				panic(fmt.Sprintf("example 0,1 unexpected heading %d (off %d, size %d)", heading, offset, size))
			}
		},
		{1, 1}: func(offset int, heading int, size int) (math.Point2D, int) {
			switch heading {
			case 1:
				return math.Point2D{X: 2*size + 1, Y: 3*size - offset}, -1
			case 3:
				return math.Point2D{X: 2*size + 1, Y: 1 + offset}, -3
			default:
				panic(fmt.Sprintf("example 1,1 unexpected heading %d (off %d, size %d)", heading, offset, size))
			}
		},
		{2, 1}: func(offset int, heading int, size int) (math.Point2D, int) {
			switch heading {
			case 0:
				return math.Point2D{X: 4*size - offset, Y: 2*size + 1}, 1
			default:
				panic(fmt.Sprintf("example 2,1 unexpected heading %d (off %d, size %d)", heading, offset, size))
			}
		},
		{2, 2}: func(offset int, heading int, size int) (math.Point2D, int) {
			switch heading {
			case 1:
				return math.Point2D{X: size - offset, Y: 2 * size}, 2
			case 2:
				return math.Point2D{X: 2*size - offset, Y: 2 * size}, 1
			default:
				panic(fmt.Sprintf("example 2,2 unexpected heading %d (off %d, size %d)", heading, offset, size))
			}
		},
		{3, 2}: func(offset int, heading int, size int) (math.Point2D, int) {
			switch heading {
			case 0:
				return math.Point2D{X: 3 * size, Y: size - offset}, 2
			case 1:
				return math.Point2D{X: 1, Y: 2*size - offset}, -1
			case 3:
				return math.Point2D{X: 3 * size, Y: 2*size - offset}, -1
			default:
				panic(fmt.Sprintf("example 3,2 unexpected heading %d (off %d, size %d)", heading, offset, size))
			}
		},
	},
	// input:
	2 + 4 + 64 + 1024 + 2048 + 32768: {
		{1, 0}: func(offset int, heading int, size int) (math.Point2D, int) {
			switch heading {
			case 2:
				return math.Point2D{X: 1, Y: 3*size - offset}, -2
			case 3:
				return math.Point2D{X: 1, Y: 3*size + 1 + offset}, -3
			default:
				panic(fmt.Sprintf("input 1,0 unexpected heading %d (off %d, size %d)", heading, offset, size))
			}
		},
		{2, 0}: func(offset int, heading int, size int) (math.Point2D, int) {
			switch heading {
			case 0:
				return math.Point2D{X: 2 * size, Y: 3*size - offset}, 2
			case 1:
				return math.Point2D{X: 2 * size, Y: size + 1 + offset}, 1
			case 3:
				return math.Point2D{X: 1 + offset, Y: 4 * size}, 0
			default:
				panic(fmt.Sprintf("input 2,0 unexpected heading %d (off %d, size %d)", heading, offset, size))
			}
		},
		{1, 1}: func(offset int, heading int, size int) (math.Point2D, int) {
			switch heading {
			case 0:
				return math.Point2D{X: 2*size + 1 + offset, Y: size}, 3
			case 2:
				return math.Point2D{X: 1 + offset, Y: 2*size + 1}, -1
			default:
				panic(fmt.Sprintf("input 1,1 unexpected heading %d (off %d, size %d)", heading, offset, size))
			}
		},
		{0, 2}: func(offset int, heading int, size int) (math.Point2D, int) {
			switch heading {
			case 2:
				return math.Point2D{X: size + 1, Y: size - offset}, -2
			case 3:
				return math.Point2D{X: size + 1, Y: size + 1 + offset}, -3
			default:
				panic(fmt.Sprintf("input 0,2 unexpected heading %d (off %d, size %d)", heading, offset, size))
			}
		},
		{1, 2}: func(offset int, heading int, size int) (math.Point2D, int) {
			switch heading {
			case 0:
				return math.Point2D{X: 3 * size, Y: size - offset}, 2
			case 1:
				return math.Point2D{X: size, Y: 3*size + 1 + offset}, 1
			default:
				panic(fmt.Sprintf("input 1,2 unexpected heading %d (off %d, size %d)", heading, offset, size))
			}
		},
		{0, 3}: func(offset int, heading int, size int) (math.Point2D, int) {
			switch heading {
			case 0:
				return math.Point2D{X: size + 1 + offset, Y: 3 * size}, 3
			case 1:
				return math.Point2D{X: 2*size + 1 + offset, Y: 1}, 0
			case 2:
				return math.Point2D{X: size + 1 + offset, Y: 1}, -1
			default:
				panic(fmt.Sprintf("input 0,3 unexpected heading %d (off %d, size %d)", heading, offset, size))
			}
		},
	},
}

func connectTiles(tiles map[math.Point2D]*tile, maxX int, maxY int) {
	var size int
	if maxX > maxY {
		if maxX/2 < maxY {
			// 4x3
			size = maxX / 4
		} else {
			// 5x2
			size = maxX / 5
		}
	} else {
		if maxY/2 < maxX {
			// 3x4
			size = maxY / 4
		} else {
			// 2x5
			size = maxY / 5
		}
	}
	key := 0
	for y := 0; y < maxY/size; y++ {
		for x := 0; x < maxX/size; x++ {
			if tiles[math.Point2D{X: x*size + 1, Y: y*size + 1}] != nil {
				fmt.Println("x, y:", x*size+1, y*size+1, "=>", 1<<(y*5+x))
				key |= 1 << (y*5 + x)
			}
		}
	}
	net := nets[key]
	fmt.Println("key:", key, "=> net:", net)
	fmt.Println("size:", size)
	for p, t := range tiles {
		netPos := math.Point2D{X: (t.pos.X - 1) / size, Y: (t.pos.Y - 1) / size}
		t.bottom = tiles[p.AddXY(0, 1)]
		var pos math.Point2D
		if t.bottom == nil {
			pos, t.bottomTurn = net[netPos]((p.X-1)%size, 1, size)
			t.bottom = tiles[pos]
			fmt.Printf("pos: %#v, netPos %#v: %p b=> %#v\n", t.pos, netPos, net[netPos], pos)
			if t.bottom == nil {
				panic(fmt.Sprintf("bottom not found: %#v", t.pos))
			}
		}
		t.left = tiles[p.AddXY(-1, 0)]
		if t.left == nil {
			pos, t.leftTurn = net[netPos]((p.Y-1)%size, 2, size)
			t.left = tiles[pos]
			fmt.Printf("pos: %#v, netPos %#v: %p l=> %#v\n", t.pos, netPos, net[netPos], pos)
			if t.left == nil {
				panic(fmt.Sprintf("left not found: %#v", t.pos))
			}
		}
		t.right = tiles[p.AddXY(1, 0)]
		if t.right == nil {
			pos, t.rightTurn = net[netPos]((p.Y-1)%size, 0, size)
			t.right = tiles[pos]
			fmt.Printf("pos: %#v, netPos %#v: %p r=> %#v\n", t.pos, netPos, net[netPos], pos)
			if t.right == nil {
				panic(fmt.Sprintf("right not found: %#v", t.pos))
			}
		}
		t.top = tiles[p.AddXY(0, -1)]
		if t.top == nil {
			pos, t.topTurn = net[netPos]((p.X-1)%size, 3, size)
			t.top = tiles[pos]
			fmt.Printf("pos: %#v, netPos %#v: %p t=> %#v\n", t.pos, netPos, net[netPos], pos)
			if t.top == nil {
				panic(fmt.Sprintf("top not found: %#v", t.pos))
			}
		}
	}
}

func parseDirections(s string) (directions []direction) {
	for len(s) > 0 {
		if s[0] == 'R' {
			directions = append(directions, direction{0, 1})
			s = s[1:]
		} else if s[0] == 'L' {
			directions = append(directions, direction{0, -1})
			s = s[1:]
		}
		i := 1
		for ; i < len(s); i++ {
			if s[i] == 'R' || s[i] == 'L' {
				break
			}
		}
		directions = append(directions, direction{io.ParseInt(s[:i]), 0})
		s = s[i:]
	}
	return
}

func parseMap(input []string) (map[math.Point2D]*tile, math.Point2D) {
	var start math.Point2D
	tiles := map[math.Point2D]*tile{}
	maxX := 0
	maxY := 0
	for y, line := range input {
		if y+1 > maxY {
			maxY = y + 1
		}
		for x, c := range line {
			if x+1 > maxX {
				maxX = x + 1
			}
			if c == ' ' {
				continue
			}
			p := math.Point2D{X: x + 1, Y: y + 1}
			if len(tiles) == 0 {
				start = p
			}
			tiles[p] = &tile{pos: p, open: c == '.'}
		}
	}
	connectTiles(tiles, maxX, maxY)
	return tiles, start
}

type direction struct {
	move int
	turn int
}

type tile struct {
	bottom     *tile
	bottomTurn int
	left       *tile
	leftTurn   int
	open       bool
	pos        math.Point2D
	right      *tile
	rightTurn  int
	top        *tile
	topTurn    int
}
