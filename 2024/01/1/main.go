package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/util"
)

func main() {
	var a, b []int
	for _, line := range io.ReadLines() {
		aAndB := strings.Split(line, "   ")
		a = append(a, io.ParseInt(aAndB[0]))
		b = append(b, io.ParseInt(aAndB[1]))
	}
	a = util.Sort(a)
	b = util.Sort(b)

	distance := 0
	for i, valueA := range a {
		distance += int(math.Abs(float64(valueA - b[i])))
	}
	fmt.Println(distance)
}
