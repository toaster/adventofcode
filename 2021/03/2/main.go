package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	lines := io.ReadLines()
	var o2Candidates []string
	var co2Candidates []string
	for _, line := range lines {
		o2Candidates = append(o2Candidates, line)
		co2Candidates = append(co2Candidates, line)
	}
	var o2Line string
	var co2Line string
	for i := 0; i < len(lines[0]); i++ {
		o2Candidates = sieve(o2Candidates, i, true)
		co2Candidates = sieve(co2Candidates, i, false)
		if len(o2Candidates) == 1 {
			o2Line = o2Candidates[0]
		}
		if len(co2Candidates) == 1 {
			co2Line = co2Candidates[0]
		}
	}
	o2Rate := bin2Int(o2Line)
	co2Rate := bin2Int(co2Line)
	fmt.Println(o2Rate * co2Rate)
}

func sieve(candidates []string, pos int, keepMostCommon bool) []string {
	var setCount int
	var setLines []string
	var unsetLines []string
	for _, line := range candidates {
		if int(line[pos]-'0') == 1 {
			setCount++
			setLines = append(setLines, line)
		} else {
			unsetLines = append(unsetLines, line)
		}
	}
	if setCount*2 >= len(candidates) {
		if keepMostCommon {
			return setLines
		}
		return unsetLines
	}

	if keepMostCommon {
		return unsetLines
	}
	return setLines
}

func bin2Int(bin string) (res int) {
	for _, r := range bin {
		res = res << 1
		res += int(r - '0')
	}
	return
}
