package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	start := make(chan bool)
	wires := map[string]chan bool{}
	input := strings.Split(io.ReadAll(), "\n\n")
	for _, line := range strings.Split(input[0], "\n") {
		nameAndValue := strings.Split(line, ": ")
		wires[nameAndValue[0]] = make(chan bool)
		go func() {
			<-start
			value := nameAndValue[1] == "1"
			for {
				wires[nameAndValue[0]] <- value
			}
		}()
	}
	for _, line := range strings.Split(strings.TrimSpace(input[1]), "\n") {
		inAndOut := strings.Split(line, " -> ")
		in := strings.Split(inAndOut[0], " ")
		if wires[in[0]] == nil {
			wires[in[0]] = make(chan bool)
		}
		if wires[in[2]] == nil {
			wires[in[2]] = make(chan bool)
		}
		if wires[inAndOut[1]] == nil {
			wires[inAndOut[1]] = make(chan bool)
		}
		switch in[1] {
		case "AND":
			go func() {
				<-start
				a := <-wires[in[0]]
				b := <-wires[in[2]]
				out := a && b
				for {
					wires[inAndOut[1]] <- out
				}
			}()
		case "OR":
			go func() {
				<-start
				a := <-wires[in[0]]
				b := <-wires[in[2]]
				out := a || b
				for {
					wires[inAndOut[1]] <- out
				}
			}()
		case "XOR":
			go func() {
				<-start
				a := <-wires[in[0]]
				b := <-wires[in[2]]
				out := a != b
				for {
					wires[inAndOut[1]] <- out
				}
			}()
		default:
			panic("invalid input")
		}
	}

	var outNames []string
	for name := range wires {
		if strings.HasPrefix(name, "z") {
			outNames = append(outNames, name)
		}
	}
	slices.Sort(outNames)
	slices.Reverse(outNames)
	go func() {
		for {
			start <- true
		}
	}()
	out := 0
	for _, name := range outNames {
		out <<= 1
		if <-wires[name] {
			out |= 1
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
	fmt.Println()
	fmt.Println(out)
}
