package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	var a []int
	b := map[int]int{}
	for _, line := range io.ReadLines() {
		aAndB := strings.Split(line, "   ")
		a = append(a, io.ParseInt(aAndB[0]))
		b[io.ParseInt(aAndB[1])]++
	}

	similarity := 0
	for _, valueA := range a {
		similarity += valueA * b[valueA]
	}
	fmt.Println(similarity)
}
