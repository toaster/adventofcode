package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	sheet := map[point]bool{}
	sheetComplete := false
	var instructions []instruction
	for _, line := range io.ReadLines() {
		if line == "" {
			sheetComplete = true
			continue
		}

		if sheetComplete {
			raw := strings.Split(strings.Split(line, " ")[2], "=")
			value, err := strconv.Atoi(raw[1])
			io.ReportError("", err)
			instructions = append(instructions, instruction{axis(raw[0][0]), value})
		} else {
			coords := io.ParseInts(line, ",")
			sheet[point{coords[0], coords[1]}] = true
		}
	}
	fmt.Println(len(sheet))

	printSheet(sheet)
	for _, i := range instructions {
		if i.a == axisY {
			for p := range sheet {
				if p.y > i.value {
					sheet[point{p.x, 2*i.value - p.y}] = true
					delete(sheet, p)
				}
			}
		} else {
			for p := range sheet {
				if p.x > i.value {
					sheet[point{2*i.value - p.x, p.y}] = true
					delete(sheet, p)
				}
			}
		}
		printSheet(sheet)
	}
	printSheet(sheet)
	fmt.Println(len(sheet))
}

func printSheet(sheet map[point]bool) {
	for y := 0; y < 20; y++ {
		for x := 0; x < 100; x++ {
			if sheet[point{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type axis rune

const (
	axisX axis = 'x'
	axisY axis = 'y'
)

type instruction struct {
	a     axis
	value int
}

type point struct {
	x int
	y int
}
