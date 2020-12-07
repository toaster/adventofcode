package icc

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// ICC is the IntCodeComputer.
type ICC struct {
	in   *bufio.Reader
	out  *bufio.Writer
	m    []int
	ip   int
	cmds map[int]executor
}

// New creates a new IntCodeComputer reading from int and writing to out.
func New(in io.Reader, out io.Writer) *ICC {
	c := ICC{in: bufio.NewReader(in), out: bufio.NewWriter(out)}
	c.cmds = map[int]executor{
		1: c.add,
		2: c.mul,
		3: c.rd,
		4: c.wr,
		5: c.jnz,
		6: c.jz,
		7: c.lt,
		8: c.eq,
	}
	return &c
}

// Load loads the given program into the computer's memory.
func (c *ICC) Load(input string) {
	for _, s := range strings.Split(strings.TrimSpace(input), ",") {
		i, _ := strconv.Atoi(s)
		c.m = append(c.m, i)
	}
}

// Run runs the currently loaded program.
func (c *ICC) Run() {
	c.ip = 0
	for {
		op, pm := c.decodeInstruction()
		if op == 99 {
			return
		}
		c.ip += c.cmds[op](pm)
	}
}

type executor func(pm []int) int

func (c *ICC) decodeInstruction() (int, []int) {
	oc := c.m[c.ip] % 100
	pm := make([]int, 3)
	pi := 0
	for p := c.m[c.ip] / 100; p > 0; p /= 10 {
		pm[pi] = p % 10
		pi++
	}
	return oc, pm
}

func (c *ICC) add(pm []int) int {
	c.m[c.m[c.ip+3]] = c.pv(1, pm) + c.pv(2, pm)
	return 4
}

func (c *ICC) mul(pm []int) int {
	c.m[c.m[c.ip+3]] = c.pv(1, pm) * c.pv(2, pm)
	return 4
}

func (c *ICC) rd(_ []int) int {
	_, _ = c.out.WriteString("> ")
	_ = c.out.Flush()
	in, _ := c.in.ReadString('\n')
	c.m[c.m[c.ip+1]], _ = strconv.Atoi(strings.TrimSpace(in))
	return 2
}

func (c *ICC) wr(pm []int) int {
	_, _ = c.out.WriteString(fmt.Sprintln(c.pv(1, pm)))
	_ = c.out.Flush()
	return 2
}

func (c *ICC) jnz(pm []int) int {
	if c.pv(1, pm) != 0 {
		return c.pv(2, pm) - c.ip
	}
	return 3
}

func (c *ICC) jz(pm []int) int {
	if c.pv(1, pm) == 0 {
		return c.pv(2, pm) - c.ip
	}
	return 3
}

func (c *ICC) lt(pm []int) int {
	if c.pv(1, pm) < c.pv(2, pm) {
		c.m[c.m[c.ip+3]] = 1
	} else {
		c.m[c.m[c.ip+3]] = 0
	}
	return 4
}

func (c *ICC) eq(pm []int) int {
	if c.pv(1, pm) == c.pv(2, pm) {
		c.m[c.m[c.ip+3]] = 1
	} else {
		c.m[c.m[c.ip+3]] = 0
	}
	return 4
}

func (c *ICC) pv(p int, pm []int) int {
	v := c.m[c.ip+p]
	switch pm[p-1] {
	case 0:
		return c.m[v]
	}
	return v
}
