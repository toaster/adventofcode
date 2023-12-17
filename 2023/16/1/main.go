package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2023/16/lavacave"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	cave := lavacave.Parse(io.ReadLines())
	fmt.Println(cave.CountEnergizedFloorTiles(lavacave.TopLeftBeamHeadingEast()))
}
