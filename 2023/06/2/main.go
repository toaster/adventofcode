package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	lines := io.ReadLines()
	time := io.ParseInt(strings.ReplaceAll(lines[0][6:], " ", ""))
	distanceToBeat := io.ParseInt(strings.ReplaceAll(lines[1][9:], " ", ""))
	waysToWin := 0
	for j := 1; j < time; j++ {
		if j*(time-j) > distanceToBeat {
			waysToWin++
		}
	}
	fmt.Println(waysToWin)
}
