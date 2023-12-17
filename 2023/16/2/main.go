package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2023/16/lavacave"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	cave := lavacave.Parse(io.ReadLines())
	maxEnergized := 0
	for _, beam := range cave.PossibleInputBeams() {
		energized := cave.CountEnergizedFloorTiles(beam)
		if energized > maxEnergized {
			maxEnergized = energized
		}
	}
	fmt.Println(maxEnergized)
}
