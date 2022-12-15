package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "<max x or y>")
		os.Exit(1)
	}

	maxXY := io.ParseInt(os.Args[1])

	var sensors []*sensor
	for _, line := range io.ReadLines() {
		sensors = append(sensors, parseSensor(line))
	}

	for y := 0; y <= maxXY; y++ {
		var ranges []*math.Range
		for _, s := range sensors {
			rowDistance := s.pos.ManhattanDistance(math.Point2D{X: s.pos.X, Y: y})
			horizontalRange := s.scannedDistance - rowDistance
			if horizontalRange < 0 {
				continue
			}

			ranges = append(ranges, &math.Range{Start: s.pos.X - horizontalRange, End: s.pos.X + horizontalRange})
		}
		sort.Slice(ranges, func(i, j int) bool { return ranges[i].Start < ranges[j].Start })
		r := ranges[0]
		for _, next := range ranges[1:] {
			if merged, ok := r.Merge(next); ok {
				r = merged
			} else {
				fmt.Println((r.End+1)*4000000 + y)
				os.Exit(0)
			}
		}
	}
	os.Exit(1)
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
