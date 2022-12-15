package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "<row to inspect>")
		os.Exit(1)
	}

	row := io.ParseInt(os.Args[1])

	var sensors []*sensor
	for _, line := range io.ReadLines() {
		sensors = append(sensors, parseSensor(line))
	}

	beaconsInPoints := map[math.Point2D]bool{}
	points := map[math.Point2D]bool{}
	for _, s := range sensors {
		rowDistance := s.pos.ManhattanDistance(math.Point2D{X: s.pos.X, Y: row})
		horizontalRange := s.scannedDistance - rowDistance
		if horizontalRange < 0 {
			continue
		}

		if s.beaconPos.Y == row {
			beaconsInPoints[s.beaconPos] = true
		}
		for i := -horizontalRange; i <= horizontalRange; i++ {
			points[math.Point2D{X: s.pos.X + i, Y: row}] = true
		}
	}
	fmt.Println(len(points) - len(beaconsInPoints))
}

func parseSensor(input string) *sensor {
	s := &sensor{}

	input = input[12:]
	i := strings.IndexRune(input, ',')
	s.pos.X = io.ParseInt(input[:i])

	input = input[i+4:]
	i = strings.IndexRune(input, ':')
	s.pos.Y = io.ParseInt(input[:i])

	input = input[i+25:]
	i = strings.IndexRune(input, ',')
	s.beaconPos.X = io.ParseInt(input[:i])

	input = input[i+4:]
	s.beaconPos.Y = io.ParseInt(input)

	s.scannedDistance = s.pos.ManhattanDistance(s.beaconPos)
	return s
}

type sensor struct {
	beaconPos       math.Point2D
	pos             math.Point2D
	scannedDistance int
}
