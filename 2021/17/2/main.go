package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	input := strings.Split(strings.Split(io.ReadLines()[0], ": ")[1], ", ")
	minX, maxX := parseRange(input[0])
	minY, maxY := parseRange(input[1])

	vxMax := maxX
	vxMin := int(math.Sqrt(0.25+2*float64(minX)) - 0.5)
	fmt.Println("vx:", vxMin, vxMax)

	vyMin := minY
	vyMax := -minY - 1
	fmt.Println("vy:", vyMin, vyMax)

	count := 0
	for xVelocity := vxMin; xVelocity <= vxMax; xVelocity++ {
		for yVelocity := vyMin; yVelocity <= vyMax; yVelocity++ {
			x := 0
			y := 0
			hit := false
			vx := xVelocity
			vy := yVelocity
			for i := 0; x <= maxX && y >= minY; i++ {
				x += vx
				y += vy
				if vx > 0 {
					vx--
				}
				vy--
				// fmt.Printf("%d: %d,%d\n", i, x, y)
				if x >= minX && x <= maxX && y >= minY && y <= maxY {
					hit = true
					fmt.Printf("HIT @%d: %d,%d\n", i, x, y)
					break
				}
			}
			if hit {
				count++
			}
		}
	}
	fmt.Println(count)
}

func parseRange(raw string) (min, max int) {
	points := strings.Split(strings.Split(raw, "=")[1], "..")
	var err error
	min, err = strconv.Atoi(points[0])
	io.ReportError("failed to parse range", err)
	max, err = strconv.Atoi(points[1])
	io.ReportError("failed to parse range", err)
	return
}
