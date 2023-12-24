package main

import (
	"fmt"
	math2 "math"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	input := io.ReadLines()
	favouriteNumber := io.ParseInt(input[0])
	start := parsePoint(input[1])
	isOpenSpace := func(p math.Point2D) bool {
		x := p.X
		y := p.Y
		if x < 0 || y < 0 {
			return false
		}

		return math.CountSetBits(x*x+3*x+2*x*y+y+y*y+favouriteNumber)%2 == 0
	}
	wholeMap := math.Rectangle2D{TopLeft: math.Point2D{}, BottomRight: math.Point2D{X: math2.MaxInt, Y: math2.MaxInt}}
	adjacents := func(p math.Point2D) []math.Point2D {
		var adjacents []math.Point2D
		for _, n := range p.Neighbours(wholeMap) {
			if isOpenSpace(n) {
				adjacents = append(adjacents, n)
			}
		}
		return adjacents
	}
	fmt.Println(math.CountPossibleDestinations(adjacents, start, 50))
}

func parsePoint(input string) math.Point2D {
	startCoordinates := io.ParseInts(input, ",")
	return math.Point2D{X: startCoordinates[0], Y: startCoordinates[1]}
}
