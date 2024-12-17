package cc

import (
	"math"

	"github.com/toaster/advent_of_code/internal/io"
)

// Parse initializes a new CC from the given input.
func Parse(lines []string) *CC {
	c := &CC{
		mem: io.ParseInts(lines[4][9:], ","),
		regs: []int{
			io.ParseInt(lines[0][12:]),
			io.ParseInt(lines[1][12:]),
			io.ParseInt(lines[2][12:]),
		},
	}
	c.cmds = map[int]executor{
		0: c.adv,
		1: c.bxl,
		2: c.bst,
		3: c.jnz,
		4: c.bxc,
		5: c.out,
		6: c.bdv,
		7: c.cdv,
	}
	return c
}

// CC is a Chronospatial Computer (https://adventofcode.com/2024/day/17).
// TODO: merge it with 2019’s ICC into a generic computer :D.
type CC struct {
	mem    []int
	ip     int
	cmds   map[int]executor
	output chan<- int
	regs   []int
}

// PatchRegisterA changes the value of register A.
func (c *CC) PatchRegisterA(value int) {
	c.regs[0] = value
}

// Run runs the computer writing to the given output channel.
func (c *CC) Run(out chan<- int) {
	c.output = out
	c.ip = 0
	for {
		opcode, params, halt := c.decodeInstruction()
		if halt {
			break
		}
		c.ip += c.cmds[opcode](params)
	}
	close(out)
}

func (c *CC) adv(params []int) int {
	c.regs[0] = c.divide(params)
	return 2
}

func (c *CC) bdv(params []int) int {
	c.regs[1] = c.divide(params)
	return 2
}

func (c *CC) bst(params []int) int {
	c.regs[1] = c.combo(params[0]) % 8 // or c.combo(params[0]) & 7 :)
	return 2
}

func (c *CC) bxc(_ []int) int {
	c.regs[1] ^= c.regs[2]
	return 2
}

func (c *CC) bxl(params []int) int {
	c.regs[1] ^= params[0]
	return 2
}

func (c *CC) cdv(params []int) int {
	c.regs[2] = c.divide(params)
	return 2
}

func (c *CC) combo(arg int) int {
	switch arg {
	case 0, 1, 2, 3:
		return arg
	case 4, 5, 6:
		return c.regs[arg-4]
	case 7:
		panic("reserved combo")
	default:
		panic("unexpected combo")
	}
}

func (c *CC) decodeInstruction() (int, []int, bool) {
	if c.ip >= len(c.mem) {
		return 0, nil, true
	}
	if c.ip == len(c.mem)-1 {
		panic("unexpected IP")
	}

	return c.mem[c.ip], []int{c.mem[c.ip+1]}, false
}

func (c *CC) divide(params []int) int {
	numerator := c.regs[0]
	denonimator := int(math.Pow(2, float64(c.combo(params[0]))))
	return numerator / denonimator
}

func (c *CC) jnz(params []int) int {
	if c.regs[0] != 0 {
		return params[0] - c.ip
	}

	return 2
}

func (c *CC) out(params []int) int {
	c.output <- c.combo(params[0]) % 8
	return 2
}

// Note: Using a slice of params is not necessary for the task but already prepares the generic computer.
type executor func(params []int) int
