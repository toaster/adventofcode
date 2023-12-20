package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

const startWorfkflow = "in"
const accepted = "A"
const rejected = "R"

func main() {
	workflows := map[string]workflow{}
	for _, line := range io.ReadLines() {
		if line == "" {
			break
		}
		nameAndSpec := strings.Split(line, "{")
		defs := strings.Split(nameAndSpec[1][:len(nameAndSpec[1])-1], ",")
		wf := workflow{}
		for _, def := range defs {
			condAndTarget := strings.Split(def, ":")
			if len(condAndTarget) > 1 {
				var fn processor
				switch condAndTarget[0][1] {
				case '>':
					fn = func(p part) (part, string, part) {
						lower, higher := p.split(condAndTarget[0][0], io.ParseInt(condAndTarget[0][2:]), true)
						return higher, condAndTarget[1], lower
					}
				case '<':
					fn = func(p part) (part, string, part) {
						lower, higher := p.split(condAndTarget[0][0], io.ParseInt(condAndTarget[0][2:]), false)
						return lower, condAndTarget[1], higher
					}
				}
				wf.steps = append(wf.steps, fn)
			} else {
				wf.steps = append(wf.steps, func(p part) (part, string, part) { return p, condAndTarget[0], part{} })
			}
		}
		workflows[nameAndSpec[0]] = wf
	}

	fmt.Println(process(part{
		x: math.Range{Start: 1, End: 4000},
		m: math.Range{Start: 1, End: 4000},
		a: math.Range{Start: 1, End: 4000},
		s: math.Range{Start: 1, End: 4000},
	}, startWorfkflow, workflows))
}

type part struct {
	x math.Range
	m math.Range
	a math.Range
	s math.Range
}

func (p part) split(category byte, value int, equalToLower bool) (part, part) {
	lower := p
	higher := p
	switch category {
	case 'x':
		if equalToLower {
			if p.x.Start > value {
				lower = part{}
			} else if p.x.End <= value {
				higher = part{}
			} else {
				lower.x.End = value
				higher.x.Start = value + 1
			}
		} else {
			if p.x.Start >= value {
				lower = part{}
			} else if p.x.End < value {
				higher = part{}
			} else {
				lower.x.End = value - 1
				higher.x.Start = value
			}
		}
	case 'm':
		if equalToLower {
			if p.m.Start > value {
				lower = part{}
			} else if p.m.End <= value {
				higher = part{}
			} else {
				lower.m.End = value
				higher.m.Start = value + 1
			}
		} else {
			if p.m.Start >= value {
				lower = part{}
			} else if p.m.End < value {
				higher = part{}
			} else {
				lower.m.End = value - 1
				higher.m.Start = value
			}
		}
	case 'a':
		if equalToLower {
			if p.a.Start > value {
				lower = part{}
			} else if p.a.End <= value {
				higher = part{}
			} else {
				lower.a.End = value
				higher.a.Start = value + 1
			}
		} else {
			if p.a.Start >= value {
				lower = part{}
			} else if p.a.End < value {
				higher = part{}
			} else {
				lower.a.End = value - 1
				higher.a.Start = value
			}
		}
	case 's':
		if equalToLower {
			if p.s.Start > value {
				lower = part{}
			} else if p.s.End <= value {
				higher = part{}
			} else {
				lower.s.End = value
				higher.s.Start = value + 1
			}
		} else {
			if p.s.Start >= value {
				lower = part{}
			} else if p.s.End < value {
				higher = part{}
			} else {
				lower.s.End = value - 1
				higher.s.Start = value
			}
		}
	}
	return lower, higher
}

func (p part) value() int {
	return p.x.Length() * p.m.Length() * p.a.Length() * p.s.Length()
}

type processor func(part) (part, string, part)

type workflow struct {
	steps []processor
}

func process(p part, workflowName string, workflows map[string]workflow) int {
	sum := 0
	wf := workflows[workflowName]
	for _, step := range wf.steps {
		var outcome string
		var np part
		p, outcome, np = step(p)
		if outcome == accepted {
			sum += p.value()
		} else if outcome != rejected {
			sum += process(p, outcome, workflows)
		}
		p = np
	}
	return sum
}
