package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	lines := io.ReadLines()
	pos1 := parseLine(lines[0])
	pos2 := parseLine(lines[1])
	fieldCount := 10
	d := &deterministicDice{sideCount: 100}
	winScore := 1000
	score1 := 0
	score2 := 0
	for score1 < winScore && score2 < winScore {
		pos1, score1 = turn(d, fieldCount, pos1, score1)
		if score1 < winScore {
			pos2, score2 = turn(d, fieldCount, pos2, score2)
		}
	}
	fmt.Println(score1, score2, math.MinInt(score1, score2)*d.Rolls())
}

func turn(d *deterministicDice, fieldCount int, pos, score int) (newPos, newScore int) {
	throw1 := d.Roll()
	throw2 := d.Roll()
	throw3 := d.Roll()
	newPos = ((pos - 1 + throw1 + throw2 + throw3) % fieldCount) + 1
	newScore = score + newPos
	fmt.Printf("Player rolls %d+%d+%d and moves to space %d for a total score of %d.\n", throw1, throw2, throw3, newPos, newScore)
	return
}

func parseLine(line string) int {
	value, err := strconv.Atoi(strings.Split(line, ": ")[1])
	io.ReportError("", err)
	return value
}

type dice interface {
	Roll() int
	Rolls() int
}

type deterministicDice struct {
	next      int
	sideCount int
	rolls     int
}

var _ dice = (*deterministicDice)(nil)

func (d *deterministicDice) Roll() int {
	v := d.next + 1
	d.next = v % d.sideCount
	d.rolls++
	return v
}

func (d *deterministicDice) Rolls() int {
	return d.rolls
}
