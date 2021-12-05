package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	io.ReportError("failed to read standard input", err)

	var numbers []int
	var boards []*board
	var current *board
	var y int
	for i, line := range strings.Split(strings.Trim(string(input), "\n"), "\n") {
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
		allWon := true
		for bi, b := range boards {
			if b.won {
				continue
			}
			if mark(b, n) {
				b.won = true
				fmt.Println("board", bi+1, "wins at number", ni+1, "which is", n)
				fmt.Println("its score is:", computeScore(b, n))
			} else {
				allWon = false
			}
		}
		if allWon {
			fmt.Println("STOP")
			break
		}
	}
}

func computeScore(b *board, n int) (score int) {
	for _, f := range b.fieldsByValues {
		if !f.m {
			score += f.v
		}
	}
	score *= n
	return
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
	won            bool
}

type field struct {
	m bool
	v int
	x int
	y int
}
