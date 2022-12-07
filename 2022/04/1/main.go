package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

type pair struct {
	a *math.Range
	b *math.Range
}

func main() {
	var pairs []*pair
	for _, line := range io.ReadLines() {
		rawPairs := strings.Split(line, ",")
		numsA := io.ParseInts(rawPairs[0], "-")
		numsB := io.ParseInts(rawPairs[1], "-")
		pairs = append(pairs, &pair{
			a: &math.Range{Start: numsA[0], End: numsA[1]},
			b: &math.Range{Start: numsB[0], End: numsB[1]},
		})
	}

	count := 0
	for _, p := range pairs {
		if p.a.Includes(p.b) || p.b.Includes(p.a) {
			count++
		}
	}
	fmt.Println(count)
}
