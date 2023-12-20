package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

const startWorfkflow = "in"
const accepted = "A"
const rejected = "R"

func main() {
	workflows := map[string]workflow{}
	var parts []*part
	workflowsDone := false
	for _, line := range io.ReadLines() {
		if line == "" {
			workflowsDone = true
			continue
		}
		if workflowsDone {
			defs := strings.Split(line[1:len(line)-1], ",")
			p := &part{}
			for _, def := range defs {
				p.set(def[0], io.ParseInt(def[2:]))
			}
			parts = append(parts, p)
		} else {
			nameAndSpec := strings.Split(line, "{")
			defs := strings.Split(nameAndSpec[1][:len(nameAndSpec[1])-1], ",")
			wf := workflow{}
			for _, def := range defs {
				condAndTarget := strings.Split(def, ":")
				if len(condAndTarget) > 1 {
					var fn func(*part) string
					switch condAndTarget[0][1] {
					case '>':
						fn = func(p *part) string {
							if p.get(condAndTarget[0][0]) > io.ParseInt(condAndTarget[0][2:]) {
								return condAndTarget[1]
							}
							return ""
						}
					case '<':
						fn = func(p *part) string {
							if p.get(condAndTarget[0][0]) < io.ParseInt(condAndTarget[0][2:]) {
								return condAndTarget[1]
							}
							return ""
						}
					}
					wf.steps = append(wf.steps, fn)
				} else {
					wf.steps = append(wf.steps, func(_ *part) string { return condAndTarget[0] })
				}
			}
			workflows[nameAndSpec[0]] = wf
		}
	}

	sum := 0
	for _, p := range parts {
		fmt.Printf("%#v: ", p)
		workflowName := startWorfkflow
		fmt.Print(workflowName)
		for {
			outcome := workflows[workflowName].process(p)
			fmt.Printf(" -> %s", outcome)
			if outcome == accepted {
				sum += p.value()
				break
			} else if outcome == rejected {
				break
			}
			workflowName = outcome
		}
		fmt.Println()
	}
	fmt.Println(sum)
}

type part struct {
	x int
	m int
	a int
	s int
}

func (p *part) get(category byte) int {
	switch category {
	case 'x':
		return p.x
	case 'm':
		return p.m
	case 'a':
		return p.a
	case 's':
		return p.s
	}
	return 0
}

func (p *part) set(category byte, value int) {
	switch category {
	case 'x':
		p.x = value
	case 'm':
		p.m = value
	case 'a':
		p.a = value
	case 's':
		p.s = value
	}
}

func (p *part) value() int {
	return p.x + p.m + p.a + p.s
}

type workflow struct {
	steps []func(*part) string
}

func (w workflow) process(p *part) string {
	for _, step := range w.steps {
		out := step(p)
		if out != "" {
			return out
		}
	}
	return "" // should never be reached
}
