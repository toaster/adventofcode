package gameconsole

import (
	"fmt"
	"strconv"
	"strings"
)

// GameConsole represents a virtual game console.
type GameConsole struct {
	acc          int
	instructions []instruction
	ip           int
	lastRun      []int
}

// NewGameConsole returns a new GameConsole.
func NewGameConsole() *GameConsole {
	return &GameConsole{}
}

// Acc returns the value of the console’s accumulator.
func (c *GameConsole) Acc() int {
	return c.acc
}

// FixInfiniteLoopAndRun tries to find the cause for the infinite loop, fixes it and runs the program.
// After running this the loaded program is altered.
func (c *GameConsole) FixInfiniteLoopAndRun() {
	attempts := 0
	patch := -1
	for !c.RunAndStopOnReexecution() {
		if patch > -1 {
			c.swapNopAndJmp(patch)
			patch = -1
			continue
		}

		patchable := 0
		for i := len(c.lastRun) - 1; i >= 0; i-- {
			ip := c.lastRun[i]
			if c.instructions[ip].op == acc {
				continue
			}
			if patchable < attempts {
				patchable++
				continue
			}
			patch = ip
			break
		}
		if patch < 0 {
			break
		}

		attempts++
		c.swapNopAndJmp(patch)
	}
}

// LoadProgram parses the input string as GameConsole program and loads it into the console.
func (c *GameConsole) LoadProgram(program string) error {
	c.instructions = []instruction{}
	for _, line := range strings.Split(program, "\n") {
		if line == "" {
			continue
		}

		cmdAndArgs := strings.Split(line, " ")
		arg, err := strconv.Atoi(cmdAndArgs[1])
		if err != nil {
			return fmt.Errorf("failed to parse argument %s: %w", cmdAndArgs[1], err)
		}

		cmd := instruction{arg: arg}
		switch cmdAndArgs[0] {
		case "nop":
			cmd.op = nop
		case "acc":
			cmd.op = acc
		case "jmp":
			cmd.op = jmp
		default:
			return fmt.Errorf("unrecognised command: %s", line)
		}
		c.instructions = append(c.instructions, cmd)
	}
	return nil
}

// RunAndStopOnReexecution runs the console's program until it stops or tries to re-execute an instruction.
// It returns `true` if the program terminated and `false` if an infinite loop occurred.
func (c *GameConsole) RunAndStopOnReexecution() bool {
	c.acc = 0
	c.ip = 0
	c.lastRun = []int{}
	executed := make([]bool, len(c.instructions))
	for c.ip < len(c.instructions) && !executed[c.ip] {
		executed[c.ip] = true
		c.lastRun = append(c.lastRun, c.ip)
		c.step()
	}
	return c.ip >= len(c.instructions)
}

func (c *GameConsole) step() {
	cmd := c.instructions[c.ip]
	switch cmd.op {
	case nop:
		c.ip++
	case acc:
		c.acc += cmd.arg
		c.ip++
	case jmp:
		c.ip += cmd.arg
	}
}

func (c *GameConsole) swapNopAndJmp(ip int) {
	cmd := &c.instructions[ip]
	switch cmd.op {
	case nop:
		cmd.op = jmp
	case jmp:
		cmd.op = nop
	}
}

type instruction struct {
	op  operator
	arg int
}

type operator int

const (
	nop operator = iota
	acc
	jmp
)
