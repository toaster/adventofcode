package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	lines := io.ReadLines()
	directions := lines[0]
	connections := map[string][2]string{}
	var cur []string
	for _, line := range lines[2:] {
		srcAndDest := strings.Split(line, " = ")
		src := srcAndDest[0]
		if src[2] == 'A' {
			cur = append(cur, src)
		}
		dest := srcAndDest[1]
		connections[src] = [2]string{dest[1:4], dest[6:9]}
	}
	// okay, empirically shown:
	// - there are loops
	// - all loops start at the first direction, i.e. there is no offset into the loops
	//   -> that’s special, because that means that the Z field has to lead to the same field as the starting A field,
	//      and it implies that there is no other Z field (A:Z => 1:1)
	loops := make([]int, len(cur))
	loopCount := 0
	step := 0
	for loopCount < len(cur) {
		for _, d := range directions {
			for i, c := range cur {
				if c[2] == 'Z' {
					if loops[i] == 0 {
						loops[i] = step
						loopCount++
					}
				}
			}

			for i, c := range cur {
				if d == 'L' {
					cur[i] = connections[c][0]
				} else {
					cur[i] = connections[c][1]
				}
			}
			step++
		}
	}
	neededSteps := loops[0]
	if len(loops) > 1 {
		neededSteps = math.LCM(loops[0], loops[1], loops[2:]...)
	}
	fmt.Println(neededSteps)
}
