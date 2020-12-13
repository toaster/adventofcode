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
	wx := 10
	wy := 1
	for _, i := range instructions {
		switch i.action {
		case 'F':
			x += wx * i.amount
			y += wy * i.amount
		case 'L':
			switch i.amount {
			case 0, 360:
			case 90:
				wx, wy = turnLeft(wx, wy)
			case 180:
				wx, wy = turnLeft(wx, wy)
				wx, wy = turnLeft(wx, wy)
			case 270:
				wx, wy = turnLeft(wx, wy)
				wx, wy = turnLeft(wx, wy)
				wx, wy = turnLeft(wx, wy)
			default:
				panic("unexpected angle")
			}
		case 'R':
			switch i.amount {
			case 0, 360:
			case 90:
				wx, wy = turnRight(wx, wy)
			case 180:
				wx, wy = turnRight(wx, wy)
				wx, wy = turnRight(wx, wy)
			case 270:
				wx, wy = turnRight(wx, wy)
				wx, wy = turnRight(wx, wy)
				wx, wy = turnRight(wx, wy)
			default:
				panic("unexpected angle")
			}
		default:
			wx, wy = move(i.action, i.amount, wx, wy)
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

func turnLeft(x, y int) (int, int) {
	return -y, x
}

func turnRight(x, y int) (int, int) {
	return y, -x
}
