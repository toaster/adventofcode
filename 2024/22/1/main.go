package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2024/22/monkeymarket"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	result := 0
	for _, line := range io.ReadLines() {
		buyer := monkeymarket.NewBuyer(io.ParseInt(line))
		result += buyer.ComputeSecret(2000)
	}
	fmt.Println(result)
}
