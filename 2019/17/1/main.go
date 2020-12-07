package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/toaster/advent_of_code/2019/icc"
	"github.com/toaster/advent_of_code/2019/space"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	program := icc.Parse(string(input))
	in := make(chan int)
	out := make(chan int)
	raw := bytes.NewBuffer([]byte{})
	go func() {
		for {
			fmt.Fprint(raw, string(rune(<-out)))
		}
	}()
	c := icc.New(in, out)
	c.Load(program)
	c.Run()

	var width, height int
	plan := map[space.Point]rune{}
	for y, line := range strings.Split(strings.TrimSpace(raw.String()), "\n") {
		width = len(line)
		height = y + 1
		for x, p := range line {
			plan[space.Point{X: x, Y: y}] = p
		}
	}

	alignments := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if plan[space.Point{X: x, Y: y}] != '#' {
				continue
			}
			if x > 0 && plan[space.Point{X: x - 1, Y: y}] != '#' {
				continue
			}
			if x < width-1 && plan[space.Point{X: x + 1, Y: y}] != '#' {
				continue
			}
			if y > 0 && plan[space.Point{X: x, Y: y - 1}] != '#' {
				continue
			}
			if y < height-1 && plan[space.Point{X: x, Y: y + 1}] != '#' {
				continue
			}
			alignments += x * y
		}
	}
	fmt.Println(alignments)
}
