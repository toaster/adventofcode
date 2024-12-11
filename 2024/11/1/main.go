package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2024/11/pluto"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	stones := pluto.ParseStones(io.ReadAll())
	fmt.Println(stones.CountAfter(25))
}
