package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile(os.Args[1])
	m := load(string(input))
	run(m, 0)
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
	for {
		op, pm := decodeInstruction(m[ip])
		if op == 99 {
			return
		}
		ip += commands[op](m, ip, pm)
	}
}

type executor func(m []int, ip int, pm []int) int

var commands = map[int]executor{
	1: performAdd,
	2: performMult,
	3: performRead,
	4: performWrite,
}

func decodeInstruction(i int) (int, []int) {
	oc := i % 100
	pm := make([]int, 3)
	pi := 0
	for p := i / 100; p > 0; p /= 10 {
		pm[pi] = p % 10
		pi++
	}
	return oc, pm
}

func performAdd(m []int, ip int, pm []int) int {
	m[m[ip+3]] = pv(m, ip+1, pm[0]) + pv(m, ip+2, pm[1])
	return 4
}

func performMult(m []int, ip int, pm []int) int {
	m[m[ip+3]] = pv(m, ip+1, pm[0]) * pv(m, ip+2, pm[1])
	return 4
}

func performRead(m []int, ip int, pm []int) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	in, _ := reader.ReadString('\n')
	m[m[ip+1]], _ = strconv.Atoi(strings.TrimSpace(in))
	return 2
}

func performWrite(m []int, ip int, pm []int) int {
	fmt.Println(pv(m, ip+1, pm[0]))
	return 2
}

func pv(m []int, p, pm int) int {
	v := m[p]
	switch pm {
	case 0:
		return m[v]
	}
	return v
}
