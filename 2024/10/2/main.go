package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2024/10/topography"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	m := topography.ParseMap(io.ReadLines())
	fmt.Println(m.SumUpRatings())
}
