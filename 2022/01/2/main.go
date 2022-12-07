package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	lines := append(io.ReadLines(), "")
	max := []int{0, 0, 0}
	cur := 0
	for _, line := range lines {
		if line == "" {
			if cur > max[0] {
				max = max[1:]
				if cur > max[0] {
					if cur > max[1] {
						max = append(max, cur)
					} else {
						max = append(max[0:1], cur, max[1])
					}
				} else {
					max = append([]int{cur}, max...)
				}
			}
			cur = 0
		} else {
			cur += io.ParseInt(line)
		}
	}
	fmt.Println(max[0] + max[1] + max[2])
}
