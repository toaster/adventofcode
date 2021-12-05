package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	var numbers []int
	var boards []*board
	var current *board
	var y int
	for i, line := range io.ReadLines() {
		if i == 0 {
			numbers = parseNumbers(line, ",")
			continue
		}

		if line == "" {
			current = &board{fieldsByValues: map[int]*field{}}
			boards = append(boards, current)
			y = 0
			continue
		}

		nums := parseNumbers(line, " ")
		var row []*field
		for x, num := range nums {
			f := &field{
				m: false,
				v: num,
				x: x,
				y: y,
			}
			current.fieldsByValues[num] = f
			row = append(row, f)
		}
		current.fields = append(current.fields, row)
		y++
	}

	for ni, n := range numbers {
		for bi, b := range boards {
			if mark(b, n) {
				fmt.Println("board", bi+1, "wins at number", ni+1, "which is", n)
				sum := 0
				for _, f := range b.fieldsByValues {
					if !f.m {
						sum += f.v
					}
				}
				fmt.Println("final score", sum*n)
				os.Exit(0)
			}
		}
	}
}

func printBoards(boards []*board) {
	for _, b := range boards {
		for _, r := range b.fields {
			for _, f := range r {
				m := ' '
				if f.m {
					m = 'x'
				}
				fmt.Printf("%2d %c ", f.v, m)
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func mark(b *board, n int) bool {
	if f := b.fieldsByValues[n]; f != nil {
		f.m = true
		donex := true
		doney := true
		for i := 0; i < 5; i++ {
			if !b.fields[f.y][i].m {
				doney = false
			}

			if !b.fields[i][f.x].m {
				donex = false
			}
		}

		return donex || doney
	}
	return false
}

func parseNumbers(line string, sep string) (numbers []int) {
	for _, s := range strings.Split(line, sep) {
		if s == "" {
			continue
		}

		v, err := strconv.Atoi(s)
		if err != nil {
			io.ReportError("failed to parse input", err)
		}

		numbers = append(numbers, v)
	}
	return
}

type board struct {
	fieldsByValues map[int]*field
	fields         [][]*field
}

type field struct {
	m bool
	v int
	x int
	y int
}
