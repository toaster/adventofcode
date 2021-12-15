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
	var polymer *linkedList
	rules := map[pair]byte{}
	for i, line := range io.ReadLines() {
		if i == 0 {
			var cur *linkedList
			for _, element := range []byte(line) {
				next := &linkedList{value: element}
				if cur == nil {
					polymer = next
				} else {
					cur.next = next
				}
				cur = next
			}
			continue
		} else if line == "" {
			continue
		}

		raw := strings.Split(line, " -> ")
		rules[pair{raw[0][0], raw[0][1]}] = raw[1][0]
	}
	fmt.Println("polymer?", polymer != nil)

	if len(os.Args) != 2 {
		io.ReportError("", errors.New("you have to specify the step count"))
	}
	stepCount, err := strconv.Atoi(os.Args[1])
	io.ReportError("cannot parse step count", err)

	printPolymer(polymer)
	for i := 0; i < stepCount; i++ {
		cur := polymer
		for cur != nil && cur.next != nil {
			next := cur.next
			cur.next = &linkedList{next: next, value: rules[pair{cur.value, next.value}]}
			cur = next
		}
	}
	e := countElements(polymer)
	// printPolymer(polymer)
	fmt.Println(e)
	fmt.Println(mostCommonCount(e) - leastCommonCount(e))
}

func printPolymer(polymer *linkedList) {
	for polymer != nil {
		fmt.Printf("%c", polymer.value)
		polymer = polymer.next
	}
	fmt.Println()
}

func leastCommonCount(counts map[byte]int) (min int) {
	for _, count := range counts {
		if count < min || min == 0 {
			min = count
		}
	}
	return
}

func mostCommonCount(counts map[byte]int) (max int) {
	for _, count := range counts {
		if count > max {
			max = count
		}
	}
	return
}

func countElements(polymer *linkedList) (counts map[byte]int) {
	counts = map[byte]int{}
	for polymer != nil {
		counts[polymer.value]++
		polymer = polymer.next
	}
	return
}

type linkedList struct {
	next  *linkedList
	value byte
}

type pair struct {
	a byte
	b byte
}
