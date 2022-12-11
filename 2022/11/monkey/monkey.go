package monkey

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

// Monkey describes a monkey :).
type Monkey struct {
	InspectionCount int
	items           []int
	operation       func(int) int
	test            func(int) int
	TestDivisor     int
}

// ParseMonkeys parses a couple of notes and returns a []*Monkey based on them.
func ParseMonkeys(input []string) (monkeys []*Monkey) {
	for i := 0; i < len(input); i += 7 {
		test, testDivisor := parseTest(input[i+3 : i+6])
		monkeys = append(monkeys, &Monkey{
			items:       parseItems(input[i+1]),
			operation:   parseOperation(input[i+2]),
			test:        test,
			TestDivisor: testDivisor,
		})
	}
	return
}

// PlayRound plays a round of Keep Away.
func PlayRound(monkeys []*Monkey, worryLevelModificator func(int) int) {
	for _, m := range monkeys {
		playTurn(m, monkeys, worryLevelModificator)
	}
}

func parseItems(input string) []int {
	return io.ParseInts(strings.TrimSpace(input)[16:], ", ")
}

func parseOperation(input string) func(int) int {
	components := strings.Split(input[23:], " ")
	getArg := func(old int) int {
		if components[1] == "old" {
			return old
		}

		return io.ParseInt(components[1])
	}
	switch components[0] {
	case "+":
		return func(old int) int { return old + getArg(old) }
	case "-":
		return func(old int) int { return old - getArg(old) }
	case "*":
		return func(old int) int { return old * getArg(old) }
	case "/":
		return func(old int) int { return old / getArg(old) }
	}

	io.ReportError("", errors.New("Unexpected operation: "+components[0]))
	return nil
}

func parseTest(input []string) (func(int) int, int) {
	divisor := io.ParseInt(input[0][21:])
	trueTarget := io.ParseInt(input[1][29:])
	falseTarget := io.ParseInt(input[2][30:])
	return func(level int) int {
		if level%divisor == 0 {
			return trueTarget
		}

		return falseTarget
	}, divisor
}

func playTurn(m *Monkey, monkeys []*Monkey, worryLevelModificator func(int) int) {
	for _, oldLevel := range m.items {
		m.InspectionCount++
		newLevel := m.operation(oldLevel)
		if newLevel < oldLevel {
			fmt.Println("OVERFLOW", oldLevel, "=>", newLevel)
			os.Exit(1)
		}
		newLevel = worryLevelModificator(newLevel)
		recipient := m.test(newLevel)
		monkeys[recipient].items = append(monkeys[recipient].items, newLevel)
	}
	m.items = nil
}
