package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2023/02/cubegame"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	games := cubegame.ParseRecords(io.ReadLines())
	fmt.Println(games.SumOfPowersOfMinimumSets())
}
