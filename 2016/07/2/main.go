package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	sslCount := 0
	for _, address := range io.ReadLines() {
		if supportsSSL(address) {
			sslCount++
		}
	}
	fmt.Println(sslCount)
}

func supportsSSL(address string) bool {
	inHypernetSequence := false
	babMap := map[string]bool{}
	abaMap := map[string]bool{}
	for i := 0; i < len(address); i++ {
		c := address[i]
		if c == '[' {
			inHypernetSequence = true
			continue
		} else if c == ']' {
			inHypernetSequence = false
			continue
		}
		if i < 2 {
			continue
		}
		b := address[i-1]
		if b == '[' || b == ']' || address[i-2] != c || b == c {
			continue
		}

		babOrAba := string([]byte{b, c, b})
		abaOrBab := string([]byte{c, b, c})
		if inHypernetSequence {
			if babMap[abaOrBab] {
				return true
			}
			abaMap[babOrAba] = true
		} else {
			if abaMap[abaOrBab] {
				return true
			}
			babMap[babOrAba] = true
		}
	}
	return false
}
