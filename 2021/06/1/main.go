package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	fishes := map[int]int{}
	for _, n := range io.ParseInts(io.ReadLines()[0], ",") {
		fishes[n]++
	}

	for i := 0; i < 80; i++ {
		spawning := fishes[0]
		for j := 0; j < 8; j++ {
			fishes[j] = fishes[j+1]
		}
		fishes[6] += spawning
		fishes[8] = spawning
	}

	count := 0
	for i := 0; i < 9; i++ {
		count += fishes[i]
	}
	fmt.Println(count)
}
