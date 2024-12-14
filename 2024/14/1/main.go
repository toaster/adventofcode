package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2024/14/ebhq"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	hallway := ebhq.ParseHallway(io.ReadLines())
	fmt.Println(hallway.ComputeSafetyFactor(100))
}
