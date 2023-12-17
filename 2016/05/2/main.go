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
	pass := [8]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
	count := 0
	for count < 8 {
		var hash string
		x, hash = crypto.DetectMD5HashWithPrefix(doorID, "00000", x+1)
		if i := hash[5] - '0'; i < 8 && pass[i] == ' ' {
			pass[i] = hash[6]
			fmt.Printf("\r%s", string(pass[:]))
			count++
		}
	}
	fmt.Println()
}
