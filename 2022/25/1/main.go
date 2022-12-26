package main

import (
	"fmt"
	"math"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	sum := 0
	for _, line := range io.ReadLines() {
		sum += parseSnafu(line)
	}
	printSnafu(sum)
}

func parseSnafu(line string) (num int) {
	for i := 0; i < len(line); i++ {
		n := math.Pow(5, float64(i))
		switch line[(len(line) - 1 - i)] {
		case '2':
			num += int(n) * 2
		case '1':
			num += int(n)
		case '-':
			num -= int(n)
		case '=':
			num -= int(n) * 2
		}
	}
	return
}

func printSnafu(num int) {
	var runes []rune
	addOne := false
	for num > 0 {
		if addOne {
			num++
		}
		n := num % 5
		switch n {
		case 0:
			runes = append(runes, '0')
			addOne = false
		case 1:
			runes = append(runes, '1')
			addOne = false
		case 2:
			runes = append(runes, '2')
			addOne = false
		case 3:
			runes = append(runes, '=')
			addOne = true
		case 4:
			runes = append(runes, '-')
			addOne = true
		}
		num = num / 5
	}
	for i := len(runes) - 1; i >= 0; i-- {
		fmt.Printf("%c", runes[i])
	}
	fmt.Println()
}
