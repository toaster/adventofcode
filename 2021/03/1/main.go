package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	io.ReportError("failed to read standard input", err)

	var counts []int
	lineCount := 0
	for _, line := range strings.Split(strings.Trim(string(input), "\n"), "\n") {
		lineCount++
		for i, r := range line {
			var v int
			switch r {
			case '0':
				v = 0
			case '1':
				v = 1
			default:
				io.ReportError("", fmt.Errorf("unexpected bit: %c", r))
			}
			if i < len(counts) {
				counts[i] += v
			} else {
				counts = append(counts, v)
			}
		}
	}
	gamma := 0
	epsilon := 0
	for _, count := range counts {
		gamma = gamma << 1
		epsilon = epsilon << 1
		if count > lineCount/2 {
			gamma++
		} else {
			epsilon++
		}
	}

	fmt.Println(gamma * epsilon)
}
