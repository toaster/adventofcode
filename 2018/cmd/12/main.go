package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type pixel struct {
	x, y, dx, dy int
}

func main() {
	initialRegex := regexp.MustCompile("initial state: ([.#]+)")
	ruleRegex := regexp.MustCompile("([.#]{5}) => ([.#])")
	inputFile := os.Args[1]
	gens := 20
	offset := 4
	if len(os.Args) > 2 {
		var err error
		gens, err = strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}
	}
	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")

	initial := initialRegex.FindStringSubmatch(lines[0])[1]
	pots := make([]byte, len(initial)+8)
	copy(pots[4:], initial)
	for i := 0; i < 4; i++ {
		pots[i] = '.'
		pots[len(pots)-i-1] = '.'
	}

	rules := map[string]byte{}
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		rule := ruleRegex.FindStringSubmatch(line)
		rules[rule[1]] = []byte(rule[2])[0]
	}
	fmt.Println(string(pots))
	patternFound := false
	for g := 0; g < gens; g++ {
		np := make([]byte, len(pots)+8)
		oldOffset := offset
		offset += 4
		copy(np[4:], pots)
		for i := 0; i < 4; i++ {
			np[i] = '.'
			np[len(np)-i-1] = '.'
		}
		for i := 2; i < len(pots)-2; i++ {
			np[i+4] = rules[string(pots[i-2:i+3])]
			if np[i+4] == 0 {
				np[i+4] = '.'
			}
		}
		for i, p := range np {
			if p == '#' {
				np = np[i-4:]
				offset -= i - 4
				break
			}
		}
		for i := len(np) - 1; i > 0; i-- {
			if np[i] == '#' {
				np = np[:i+5]
				break
			}
		}
		if string(pots) == string(np) {
			if !patternFound {
				patternFound = true
				fmt.Println("found pattern", g, offset-oldOffset)
			}
			offset += (gens - g - 1) * (offset - oldOffset)
			break
		}
		pots = np
		if (g+1)%1000000 == 0 {
			fmt.Println(offset, string(pots))
		}
	}
	fmt.Println(offset, string(pots))
	result := 0
	for i, p := range pots {
		if p == '#' {
			result += i - offset
		}
	}
	fmt.Println(result)
}
