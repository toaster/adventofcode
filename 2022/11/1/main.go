package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2022/11/monkey"
	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	monkeys := monkey.ParseMonkeys(io.ReadLines())
	for i := 0; i < 20; i++ {
		monkey.PlayRound(monkeys, func(level int) int { return level / 3 })
	}
	var inspectionCounts []int
	for _, m := range monkeys {
		inspectionCounts = append(inspectionCounts, m.InspectionCount)
	}
	inspectionCounts = math.Sort(inspectionCounts)
	fmt.Println(inspectionCounts[len(inspectionCounts)-1] * inspectionCounts[len(inspectionCounts)-2])
}
