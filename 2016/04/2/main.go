package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	for _, line := range io.ReadLines() {
		letterCounts := map[rune]int{}
		var letters []rune
		components := strings.Split(line, "-")
		l := len(components)
		sc := strings.Split(components[l-1], "[")
		sectorID := io.ParseInt(sc[0])
		checksum := sc[1][:len(sc[1])-1]
		for i := 0; i < l-1; i++ {
			for _, c := range components[i] {
				if letterCounts[c] == 0 {
					letters = append(letters, c)
				}
				letterCounts[c]++
			}
		}
		slices.SortFunc(letters, func(a, b rune) int {
			if letterCounts[a] < letterCounts[b] {
				return 1
			} else if letterCounts[b] < letterCounts[a] {
				return -1
			}

			if a < b {
				return -1
			}

			return 1
		})
		if checksum == string(letters[:5]) {
			name := decrypt(components[:l-1], sectorID)
			if name == "northpole object storage" {
				fmt.Println(sectorID)
				break
			}
		}
	}
}

func decrypt(words []string, id int) string {
	var buf []rune
	shift := int32(id % 26)
	for _, word := range words {
		if len(buf) > 0 {
			buf = append(buf, ' ')
		}
		for _, c := range word {
			x := c - 'a'
			x += shift
			x %= 26
			buf = append(buf, x+'a')
		}
	}
	return string(buf)
}
