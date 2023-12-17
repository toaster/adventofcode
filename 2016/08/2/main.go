package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	const width = 50
	const height = 6
	pixels := map[math.Point2D]bool{}
	for _, line := range io.ReadLines() {
		components := strings.Split(line, " ")
		switch components[0] {
		case "rect":
			widthAndHeight := io.ParseInts(components[1], "x")
			for y := 0; y < widthAndHeight[1]; y++ {
				for x := 0; x < widthAndHeight[0]; x++ {
					pixels[math.Point2D{X: x, Y: y}] = true
				}
			}
		case "rotate":
			index := io.ParseInt(components[2][2:])
			amount := io.ParseInt(components[4])
			switch components[1] {
			case "row":
				for x := width - 1; x >= 0; x-- {
					pixels[math.Point2D{X: x + amount, Y: index}] = pixels[math.Point2D{X: x, Y: index}]
				}
				for x := 0; x < amount; x++ {
					pixels[math.Point2D{X: x, Y: index}] = pixels[math.Point2D{X: x + width, Y: index}]
				}
			case "column":
				for y := height - 1; y >= 0; y-- {
					pixels[math.Point2D{X: index, Y: y + amount}] = pixels[math.Point2D{X: index, Y: y}]
				}
				for y := 0; y < amount; y++ {
					pixels[math.Point2D{X: index, Y: y}] = pixels[math.Point2D{X: index, Y: y + height}]
				}
			}
		}
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if pixels[math.Point2D{X: x, Y: y}] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
