package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2022/elfdevice"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	lines := io.ReadLines()
	indexSum := 0
	for i := 0; i < len(lines); i += 3 {
		p1 := elfdevice.ParsePacket(lines[i])
		p2 := elfdevice.ParsePacket(lines[i+1])
		fmt.Println("== Pair", i/3+1, "==")
		if p1.Compare(p2) == -1 {
			indexSum += i/3 + 1
		}
	}
	fmt.Println(indexSum)
}
