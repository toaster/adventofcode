package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	initializationSteps := strings.Split(strings.TrimSpace(io.ReadAll()), ",")
	var hashSum int
	for _, step := range initializationSteps {
		hashSum += computeHash(step)
	}
	fmt.Println(hashSum)
}

func computeHash(step string) int {
	hash := 0
	for _, c := range []byte(step) {
		hash += int(c)
		hash *= 17
		hash &= 0xFF
	}
	return hash
}
