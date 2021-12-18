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

	vxMax := math.Sqrt(0.25+2*float64(maxX)) - 0.5
	vxMin := math.Sqrt(0.25+2*float64(minX)) - 0.5
	xVelocity := int(vxMax)
	fmt.Println("vx:", vxMin, vxMax, "->", xVelocity)

	// downSteps := maxY - minY
	// fmt.Println("a:", downSteps, "Ymax:", maxY, "Ymin:", minY)
	// maxHeight := (downSteps+1)*(downSteps+1)/2 + downSteps/2 + minY
	// fmt.Println("maxHeight:", maxHeight)
	//
	// vyMax := (math.Sqrt(1-4*float64(2*maxY+downSteps*downSteps-downSteps)) - 1) / 2
	// vyMin := (math.Sqrt(1-4*float64(2*minY+downSteps*downSteps-downSteps)) - 1) / 2
	// yVelocity := int(vyMax)
	// fmt.Println("yx:", vyMin, vyMax, "->", yVelocity)
	// yVelocity = 9

	hp := 0
	yhp := 0
	for yVelocity := 0; yVelocity < 1000; yVelocity++ {
		x := 0
		y := 0
		max := 0
		hit := false
		vx := xVelocity
		vy := yVelocity
		for i := 0; x <= maxX && y >= minY; i++ {
			x += vx
			y += vy
			if y > max {
				max = y
			}
			if vx > 0 {
				vx--
			}
			vy--
			fmt.Printf("%d: %d,%d\n", i, x, y)
			if x >= minX && x <= maxX && y >= minY && y <= maxY {
				hit = true
				fmt.Println("HIT")
				break
			}
		}
		// if hp > 0 && !hit {
		// 	break
		// }
		if hit {
			hp = max
			yhp = yVelocity
		}
	}
	fmt.Println("hp:", hp, "@:", yhp)
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
