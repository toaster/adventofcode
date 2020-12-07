package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/toaster/advent_of_code/2019/icc"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	program := icc.Parse(string(input))
	in := make(chan int)
	out := make(chan int)

	var beamPoints int
	size := 50
	min := 0
	minX := min
	maxX := minX + size
	minY := min
	maxY := minY + size
	startX := minX
	for y := minY; y < maxY; y++ {
		endX := maxX
		for x := minX; x < startX; x++ {
			fmt.Print(" ")
		}
		var started bool
		for x := startX; x < maxX; x++ {
			c := icc.New(in, out)
			c.Load(program)
			go c.Run()
			in <- x
			in <- y
			if <-out == 1 {
				fmt.Print("#")
				if !started {
					startX = x
					started = true
				}
				beamPoints++
			} else {
				fmt.Print(".")
				if started {
					endX = x
					break
				}
			}
		}
		for x := endX + 1; x < maxX; x++ {
			fmt.Print(" ")
		}
		fmt.Println(" width", endX-startX, "y", y)
	}
	fmt.Println(beamPoints)
}
