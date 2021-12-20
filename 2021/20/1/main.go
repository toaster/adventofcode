package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	lines := io.ReadLines()
	enhancementMap := map[int]bool{}
	for i, c := range lines[0] {
		if c == '#' {
			enhancementMap[i] = true
		}
	}
	img := &image{
		pixels:      map[math.Point2D]bool{},
		topLeft:     math.Point2D{X: 0, Y: 0},
		bottomRight: math.Point2D{X: len(lines[2]) - 1, Y: len(lines) - 3},
	}
	for y, line := range lines[2:] {
		for x, c := range line {
			if c == '#' {
				img.pixels[math.Point2D{X: x, Y: y}] = true
			}
		}
	}

	cur := img
	printImage(cur)
	for i := 0; i < 2; i++ {
		cur = enhance(cur, enhancementMap)
		// printImage(cur)
	}
	fmt.Println(len(cur.pixels))
}

func enhance(img *image, enhancementMap map[int]bool) (enhanced *image) {
	newBackground := false
	if img.background {
		newBackground = enhancementMap[511]
	} else {
		newBackground = enhancementMap[0]
	}
	enhanced = &image{
		background:  newBackground,
		pixels:      map[math.Point2D]bool{},
		topLeft:     img.topLeft.SubtractXY(1, 1),
		bottomRight: img.bottomRight.AddXY(1, 1),
	}
	for y := enhanced.topLeft.Y; y <= enhanced.bottomRight.Y; y++ {
		for x := enhanced.topLeft.X; x <= enhanced.bottomRight.X; x++ {
			p := math.Point2D{X: x, Y: y}
			if enhancementMap[enhanceIndex(img, p)] {
				enhanced.pixels[p] = true
			}
		}
	}
	return
}

func enhanceIndex(img *image, p math.Point2D) (index int) {
	// printEnhancementWindow(img, p)
	for y := -1; y < 2; y++ {
		for x := -1; x < 2; x++ {
			index <<= 1
			if img.pixel(p.AddXY(x, y)) {
				index |= 1
			}
		}
	}
	// fmt.Println("=>", index)
	return
}

func printEnhancementWindow(img *image, p math.Point2D) {
	for y := p.Y - 1; y < p.Y+2; y++ {
		for x := p.X - 1; x < p.X+2; x++ {
			if img.pixels[(math.Point2D{X: x, Y: y})] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func printImage(img *image) {
	for y := img.topLeft.Y - 5; y <= img.bottomRight.Y+5; y++ {
		for x := img.topLeft.X - 5; x <= img.bottomRight.X+5; x++ {
			if img.pixel(math.Point2D{X: x, Y: y}) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type image struct {
	background  bool
	pixels      map[math.Point2D]bool
	topLeft     math.Point2D
	bottomRight math.Point2D
}

func (i image) pixel(p math.Point2D) bool {
	if p.IsLessThan(i.topLeft) || p.IsGreaterThan(i.bottomRight) {
		return i.background
	}
	return i.pixels[p]
}
