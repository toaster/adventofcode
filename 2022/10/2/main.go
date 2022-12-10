package main

import (
	"github.com/toaster/advent_of_code/2022/elfdevice"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	input := io.ReadLines()
	gpu := elfdevice.LoadVideoProgram(input)
	gpu.Run()
	gpu.PrintCRT()
}
