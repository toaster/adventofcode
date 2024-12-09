package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2024/08/easterbunny"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	m := easterbunny.ParseAntennaMap(io.ReadLines())
	fmt.Println(m.CountAntinodes(false))
}
