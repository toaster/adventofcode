package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2024/04/wordsearch"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	s := wordsearch.ParseWordSearch(io.ReadLines())
	fmt.Println(s.CountX("MAS"))
}
