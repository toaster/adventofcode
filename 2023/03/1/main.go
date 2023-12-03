package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2023/03/gondolaengine"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	s := gondolaengine.ParseSchematic(io.ReadLines())
	fmt.Println(s.SumOfPartNumbers())
}
