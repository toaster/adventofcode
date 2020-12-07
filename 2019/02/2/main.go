package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			m := load(string(input))
			enter(m, noun, verb)
			run(m, 0)
			if m[0] == 19690720 {
				fmt.Println(noun*100 + verb)
				return
			}
		}
	}
}

func load(input string) []int {
	var m []int
	for _, s := range strings.Split(strings.TrimSpace(input), ",") {
		i, _ := strconv.Atoi(s)
		m = append(m, i)
	}
	return m
}

func enter(m []int, noun, verb int) {
	m[1] = noun
	m[2] = verb
}

func run(m []int, ip int) {
	switch m[ip] {
	case 1:
		performAdd(m, ip)
	case 2:
		performMult(m, ip)
	case 99:
		return
	default:
		panic("burst to flames")
	}
	run(m, ip+4)
}

func performAdd(m []int, ip int) {
	m[m[ip+3]] = m[m[ip+1]] + m[m[ip+2]]
}
func performMult(m []int, ip int) {
	m[m[ip+3]] = m[m[ip+1]] * m[m[ip+2]]
}
