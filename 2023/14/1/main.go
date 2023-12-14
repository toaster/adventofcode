package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2023/14/reflectordish"
	"github.com/toaster/advent_of_code/internal/io"
)

const (
	ball rockType = 'O'
	cube rockType = '#'
	none rockType = '.'
)

type rockType rune

func main() {
	p := reflectordish.ParseControlPlatform(io.ReadLines())
	p.TiltNorth()
	fmt.Println(p.NorthLoad())
}
