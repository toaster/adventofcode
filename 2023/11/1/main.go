package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2023/11/universe"
	"github.com/toaster/advent_of_code/internal/io"
)

const expansionFactor = 2

func main() {
	u := universe.Parse(io.ReadLines())
	u.Expand(expansionFactor)
	fmt.Println(u.ComputeShortestPathBetweenAllGalaxies())
}
