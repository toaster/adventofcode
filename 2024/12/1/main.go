package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2024/12/garden"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	m := garden.ParseMap(io.ReadLines())
	fmt.Println(m.ComputeFencingCosts())
}
