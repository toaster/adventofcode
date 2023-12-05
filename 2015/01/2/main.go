package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	input := strings.TrimSpace(io.ReadAll())
	level := 0
	for i, c := range input {
		switch c {
		case '(':
			level++
		case ')':
			level--
		}
		if level == -1 {
			fmt.Println(i + 1)
			break
		}
	}
}
