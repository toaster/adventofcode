package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/crypto"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	doorID := strings.TrimSpace(io.ReadAll())
	x := 0
	for i := 0; i < 8; i++ {
		var hash string
		x, hash = crypto.DetectMD5HashWithPrefix(doorID, "00000", x+1)
		fmt.Printf("%c", hash[5])
	}
	fmt.Println()
}
