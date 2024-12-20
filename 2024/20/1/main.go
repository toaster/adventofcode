package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2024/20/racecondition"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	racetrack := racecondition.ParseRacetrack(io.ReadLines())
	fmt.Println(racetrack.CountCheatsSavingAtLeast(2, 100))
}
