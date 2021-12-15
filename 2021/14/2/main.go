package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	polymer := map[pair]int64{}
	rules := map[pair]byte{}
	for i, line := range io.ReadLines() {
		if i == 0 {
			prev := byte(' ')
			for _, element := range []byte(line) {
				polymer[pair{prev, element}]++
				prev = element
			}
			polymer[pair{prev, ' '}]++
			continue
		} else if line == "" {
			continue
		}

		raw := strings.Split(line, " -> ")
		rules[pair{raw[0][0], raw[0][1]}] = raw[1][0]
	}

	if len(os.Args) != 2 {
		io.ReportError("", errors.New("you have to specify the step count"))
	}
	stepCount, err := strconv.Atoi(os.Args[1])
	io.ReportError("cannot parse step count", err)

	for i := 0; i < stepCount; i++ {
		// printPolymer(polymer)
		changes := map[pair]int64{}
		for p, element := range rules {
			newPA := pair{p.a, element}
			newPB := pair{element, p.b}
			changes[p] -= polymer[p]
			changes[newPA] += polymer[p]
			changes[newPB] += polymer[p]
		}
		for p, c := range changes {
			polymer[p] += c
		}
	}
	e := elementCounts(polymer)
	fmt.Println(e)
	fmt.Println(mostCommonCount(e) - leastCommonCount(e))
}

func printPolymer(polymer map[pair]int64) {
	for p, c := range polymer {
		fmt.Printf("%c%c: %d, ", p.a, p.b, c)
	}
	fmt.Println()
}

func elementCounts(polymer map[pair]int64) map[byte]int64 {
	counts := map[byte]int64{}
	for p, c := range polymer {
		if p.a == ' ' {
			continue
		}
		counts[p.a] += c
	}
	return counts
}

func leastCommonCount(counts map[byte]int64) (min int64) {
	for _, count := range counts {
		if count < min || min == 0 {
			min = count
		}
	}
	return
}

func mostCommonCount(counts map[byte]int64) (max int64) {
	for _, count := range counts {
		if count > max {
			max = count
		}
	}
	return
}

type pair struct {
	a byte
	b byte
}
