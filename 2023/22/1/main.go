package main

import (
	"fmt"
	"slices"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	var bricks []*brick
	for _, line := range io.ReadLines() {
		c := math.ParseCuboid(line, "~")
		bricks = append(bricks, &brick{Cuboid: c})
	}
	slices.SortFunc(bricks, func(a, b *brick) int {
		if a.FrontBottomLeft.Z < b.FrontBottomLeft.Z {
			return -1
		}
		if a.FrontBottomLeft.Z == b.FrontBottomLeft.Z {
			return 0
		}
		return 1
	})
	freePlane := 1
	tops := map[int][]*brick{}
	for i, b := range bricks {
		height := b.FrontBottomLeft.Z - freePlane
		if height > 0 {
			b.FrontBottomLeft.Z -= height
			b.BackTopRight.Z -= height
		}
		if freePlane > 1 {
			settled := false
			for !settled && b.FrontBottomLeft.Z > 1 {
				b.FrontBottomLeft.Z--
				b.BackTopRight.Z--
				for j := i - 1; j >= 0; j-- {
					if bricks[j].Intersect(b.Cuboid) != nil {
						settled = true
						b.FrontBottomLeft.Z++
						b.BackTopRight.Z++
						break
					}
				}
			}
		}
		{
			b.FrontBottomLeft.Z--
			for _, other := range tops[b.FrontBottomLeft.Z] {
				if other.Intersect(b.Cuboid) != nil {
					other.supporting = append(other.supporting, b)
					b.supportedBy = append(b.supportedBy, other)
				}
			}
			b.FrontBottomLeft.Z++
		}
		tops[b.BackTopRight.Z] = append(tops[b.BackTopRight.Z], b)
		freePlane = max(freePlane, b.BackTopRight.Z+1)
	}
	disintegratable := 0
	for _, b := range bricks {
		disintegratable++
		for _, supported := range b.supporting {
			if len(supported.supportedBy) < 2 {
				disintegratable--
				break
			}
		}
	}
	fmt.Println(disintegratable)
}

type brick struct {
	*math.Cuboid
	supporting  []*brick
	supportedBy []*brick
}

func (b *brick) String() string {
	return fmt.Sprintf("%s supporting %v and supported by %v", b.Cuboid, len(b.supporting), len(b.supportedBy))
}
