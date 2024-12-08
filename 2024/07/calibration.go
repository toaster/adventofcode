package calibration

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

// ParseEquation creates a new equation from the input
func ParseEquation(line string) Equation {
	resultAndOperators := strings.Split(line, ":")
	result := io.ParseInt(resultAndOperators[0])
	operands := io.ParseInts(resultAndOperators[1], " ")
	return Equation{
		operands: operands,
		result:   result,
	}
}

// Equation represents the result and the operands of an equation.
type Equation struct {
	operands []int
	result   int
}

// CanBeTrue returns whether the Equation can be solved with the given amount of possible operators.
func (e Equation) CanBeTrue(operatorCount int) bool {
	gen := newTupleGenerator(len(e.operands)-1, operatorCount)
	for gen.generateNext() {
		test := e.operands[0]
		for i := 0; i < gen.length; i++ {
			switch gen.current[i] {
			case 0:
				test += e.operands[i+1]
			case 1:
				test *= e.operands[i+1]
			case 2:
				test = io.ParseInt(fmt.Sprintf("%d%d", test, e.operands[i+1]))
			}
		}
		if test == e.result {
			return true
		}
	}
	return false
}

// Result returns the result of the Equation.
func (e Equation) Result() int {
	return e.result
}

func newTupleGenerator(length, valueCount int) *tupleGenerator {
	return &tupleGenerator{length: length, valueCount: valueCount}
}

type tupleGenerator struct {
	current    []int
	done       bool
	length     int
	valueCount int
}

func (g *tupleGenerator) generateNext() bool {
	if g.done {
		return false
	}

	if g.current == nil {
		g.current = make([]int, g.length)
		return true
	}

	for i := 0; i < g.length; i++ {
		g.current[i]++
		if g.current[i] < g.valueCount {
			break
		} else {
			if i == g.length-1 {
				g.done = true
				break
			}
			g.current[i] = 0
		}
	}
	return g.done == false
}
