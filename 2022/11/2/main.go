package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2022/11/monkey"
	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	monkeys := monkey.ParseMonkeys(io.ReadLines())
	var divisors []int
	for _, m := range monkeys {
		divisors = append(divisors, m.TestDivisor)
	}
	divisor := math.LCM(divisors[0], divisors[1], divisors[2:]...)
	for i := 0; i < 10000; i++ {
		monkey.PlayRound(monkeys, func(level int) int { return level % divisor })
	}
	var inspectionCounts []int
	for _, m := range monkeys {
		inspectionCounts = append(inspectionCounts, m.InspectionCount)
	}
	inspectionCounts = math.Sort(inspectionCounts)
	fmt.Println(inspectionCounts[len(inspectionCounts)-1] * inspectionCounts[len(inspectionCounts)-2])
}
