package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	var entry math.Point2D
	var exit math.Point2D
	lines := io.ReadLines()
	xWall := math.Range{Start: 0, End: len(lines[0]) - 1}
	yWall := math.Range{Start: 0, End: len(lines) - 1}
	valley := map[math.Point2D]bool{}
	var blizzards []*blizzard
	for y, line := range lines {
		switch y {
		case yWall.Start:
			entry = math.Point2D{X: strings.IndexRune(line, '.'), Y: y}
		case yWall.End:
			exit = math.Point2D{X: strings.IndexRune(line, '.'), Y: y}
		default:
			for x, s := range line {
				if s == '#' {
					continue
				}
				p := math.Point2D{X: x, Y: y}
				var b *blizzard
				switch s {
				case '>':
					b = &blizzard{head: east}
				case 'v':
					b = &blizzard{head: south}
				case '<':
					b = &blizzard{head: west}
				case '^':
					b = &blizzard{head: north}
				}
				if b != nil {
					b.pos = p
					valley[p] = true
					blizzards = append(blizzards, b)
				}
			}
		}
	}
	sum := 0
	minutes := wander(entry, exit, xWall, yWall, blizzards)
	fmt.Println(minutes)
	sum += minutes
	minutes = wander(exit, entry, xWall, yWall, blizzards)
	fmt.Println(minutes)
	sum += minutes
	minutes = wander(entry, exit, xWall, yWall, blizzards)
	fmt.Println(minutes)
	sum += minutes
	fmt.Println(sum)
}

func moveBlizzards(blizzards []*blizzard, xRange, yRange math.Range) map[math.Point2D]bool {
	valley := map[math.Point2D]bool{}
	for _, b := range blizzards {
		p := b.pos.Add(math.Point2D(b.head))
		if p.X == xRange.Start {
			p.X = xRange.End - 1
		} else if p.X == xRange.End {
			p.X = xRange.Start + 1
		} else if p.Y == yRange.Start {
			p.Y = yRange.End - 1
		} else if p.Y == yRange.End {
			p.Y = yRange.Start + 1
		}
		b.pos = p
		valley[p] = true
	}
	return valley
}

func printValley(valley map[math.Point2D]bool, xRange, yRange math.Range) {
	for y := yRange.Start; y <= yRange.End; y++ {
		for x := xRange.Start; x <= xRange.End; x++ {
			if x == xRange.Start || x == xRange.End || y == yRange.Start || y == yRange.End {
				fmt.Print("#")
			} else if valley[math.Point2D{X: x, Y: y}] {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func wander(entry, exit math.Point2D, xWall, yWall math.Range, blizzards []*blizzard) (minutes int) {
	cur := map[math.Point2D]bool{entry: true}
	xRange := math.Range{Start: xWall.Start + 1, End: xWall.End - 1}
	yRange := math.Range{Start: yWall.Start + 1, End: yWall.End - 1}
	for {
		minutes++
		valley := moveBlizzards(blizzards, xWall, yWall)
		next := map[math.Point2D]bool{}
		for p := range cur {
			if !valley[p] {
				next[p] = true
			}
			for _, n := range p.Neighbours(xWall, yWall) {
				if n == exit {
					return
				}
				if !xRange.Covers(n.X) || !yRange.Covers(n.Y) {
					continue
				}
				if !valley[n] {
					next[n] = true
				}
			}
		}
		cur = next
	}
}

type blizzard struct {
	pos  math.Point2D
	head heading
}

type heading math.Point2D

var (
	east  = heading{X: 1}
	south = heading{Y: 1}
	west  = heading{X: -1}
	north = heading{Y: -1}
)
