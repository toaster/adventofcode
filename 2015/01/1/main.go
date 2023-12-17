package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	input := strings.TrimSpace(io.ReadAll())
	level := 0
	for _, c := range input {
		switch c {
		case '(':
			level++
		case ')':
			level--
		}
	}
	fmt.Println(level)
}
