package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2024/22/monkeymarket"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	market := monkeymarket.New(io.ParseInts(io.ReadAll(), "\n"))
	fmt.Println(market.ComputeMaxAmountOfBananas(2000))
}
