package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	action rune
	amount int
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
		os.Exit(1)
	}

	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	var instructions []instruction
	for _, line := range lines {
		amount, _ := strconv.Atoi(line[1:])
		instructions = append(instructions, instruction{[]rune(line)[0], amount})
	}
	x := 0
	y := 0
	h := 'E'
	for _, i := range instructions {
		switch i.action {
		case 'F':
			x, y = move(h, i.amount, x, y)
		case 'L':
			switch i.amount {
			case 0, 360:
			case 90:
				h = turnLeft(h)
			case 180:
				h = turnLeft(h)
				h = turnLeft(h)
			case 270:
				h = turnLeft(h)
				h = turnLeft(h)
				h = turnLeft(h)
			default:
				panic("unexpected angle")
			}
		case 'R':
			switch i.amount {
			case 0, 360:
			case 90:
				h = turnRight(h)
			case 180:
				h = turnRight(h)
				h = turnRight(h)
			case 270:
				h = turnRight(h)
				h = turnRight(h)
				h = turnRight(h)
			default:
				panic("unexpected angle")
			}
		default:
			x, y = move(i.action, i.amount, x, y)
		}
	}
	fmt.Println("distance:", int(math.Abs(float64(x))+math.Abs(float64(y))))
}

func move(d rune, a, x, y int) (int, int) {
	switch d {
	case 'N':
		y += a
	case 'E':
		x += a
	case 'S':
		y -= a
	case 'W':
		x -= a
	}
	return x, y
}

func turnLeft(h rune) rune {
	switch h {
	case 'N':
		return 'W'
	case 'W':
		return 'S'
	case 'S':
		return 'E'
	case 'E':
		return 'N'
	}
	panic("unexpected direction")
}

func turnRight(h rune) rune {
	switch h {
	case 'N':
		return 'E'
	case 'E':
		return 'S'
	case 'S':
		return 'W'
	case 'W':
		return 'N'
	}
	panic("unexpected direction")
}
