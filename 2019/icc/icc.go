package icc

import (
	"fmt"
	"strconv"
	"strings"
)

// ICC is the IntCodeComputer.
type ICC struct {
	in      chan int
	out     chan int
	m       []int
	ip      int
	cmds    map[int]executor
	relBase int
}

// New creates a new IntCodeComputer reading from int and writing to out.
func New(in, out chan int) *ICC {
	c := ICC{in: in, out: out}
	c.cmds = map[int]executor{
		1: c.add,
		2: c.mul,
		3: c.rd,
		4: c.wr,
		5: c.jnz,
		6: c.jz,
		7: c.lt,
		8: c.eq,
		9: c.ab,
	}
	return &c
}

// ComputeOptimalThrusterConfig determines the phase configuration which lead to the highest
// thruster signal for a program.
func ComputeOptimalThrusterConfig(input string) (int, []int) {
	program := Parse(input)
	in := make(chan int, 2)
	o1 := make(chan int, 2)
	o2 := make(chan int, 2)
	o3 := make(chan int, 2)
	o4 := make(chan int, 2)
	out := make(chan int)

	phaseConfig := []int{0, 1, 2, 3, 4}
	maxSig := 0
	bestConf := make([]int, len(phaseConfig))
	perm(phaseConfig, 0, func(pc []int) {
		t1 := New(in, o1)
		t2 := New(o1, o2)
		t3 := New(o2, o3)
		t4 := New(o3, o4)
		t5 := New(o4, out)
		t1.Load(program)
		t2.Load(program)
		t3.Load(program)
		t4.Load(program)
		t5.Load(program)
		go t1.Run()
		go t2.Run()
		go t3.Run()
		go t4.Run()
		go t5.Run()
		in <- pc[0]
		o1 <- pc[1]
		o2 <- pc[2]
		o3 <- pc[3]
		o4 <- pc[4]
		in <- 0
		s := <-out
		if s > maxSig {
			maxSig = s
			copy(bestConf, pc)
		}
	})
	return maxSig, bestConf
}

// ComputeOptimalLoopedThrusterConfig determines the phase configuration which lead to the highest
// thruster signal for a program.
func ComputeOptimalLoopedThrusterConfig(input string) (int, []int) {
	program := Parse(input)
	inOut := make(chan int, 2)
	o1 := make(chan int, 2)
	o2 := make(chan int, 2)
	o3 := make(chan int, 2)
	o4 := make(chan int, 2)

	phaseConfig := []int{5, 6, 7, 8, 9}
	maxSig := 0
	bestConf := make([]int, len(phaseConfig))
	perm(phaseConfig, 0, func(pc []int) {
		t1 := New(inOut, o1)
		t2 := New(o1, o2)
		t3 := New(o2, o3)
		t4 := New(o3, o4)
		t5 := New(o4, inOut)
		t1.Load(program)
		t2.Load(program)
		t3.Load(program)
		t4.Load(program)
		t5.Load(program)

		inOut <- pc[0]
		o1 <- pc[1]
		o2 <- pc[2]
		o3 <- pc[3]
		o4 <- pc[4]
		inOut <- 0
		go t1.Run()
		go t2.Run()
		go t3.Run()
		go t4.Run()
		t5.Run()
		s := <-inOut
		if s > maxSig {
			maxSig = s
			copy(bestConf, pc)
		}
	})
	return maxSig, bestConf
}

func perm(a []int, i int, f func([]int)) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, i+1, f)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, i+1, f)
		a[i], a[j] = a[j], a[i]
	}
}

// Parse converts a string which represents a program into an []int that can be loaded into an ICC.
func Parse(input string) []int {
	program := []int{}
	for _, s := range strings.Split(strings.TrimSpace(input), ",") {
		i, _ := strconv.Atoi(s)
		program = append(program, i)
	}
	return program
}

// Load loads the given program into the computer's memory.
func (c *ICC) Load(program []int) {
	c.m = make([]int, 1024*1024)
	copy(c.m, program)
}

// Patch changes the value stored at the given memory address.
func (c *ICC) Patch(addr, val int) {
	c.m[addr] = val
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
	c.m[c.pa(3, pm)] = c.pv(1, pm) + c.pv(2, pm)
	return 4
}

func (c *ICC) mul(pm []int) int {
	c.m[c.pa(3, pm)] = c.pv(1, pm) * c.pv(2, pm)
	return 4
}

func (c *ICC) rd(pm []int) int {
	x, _ := <-c.in
	c.m[c.pa(1, pm)] = x
	return 2
}

func (c *ICC) wr(pm []int) int {
	c.out <- c.pv(1, pm)
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
	a := c.pa(3, pm)
	if c.pv(1, pm) < c.pv(2, pm) {
		c.m[a] = 1
	} else {
		c.m[a] = 0
	}
	return 4
}

func (c *ICC) eq(pm []int) int {
	a := c.pa(3, pm)
	if c.pv(1, pm) == c.pv(2, pm) {
		c.m[a] = 1
	} else {
		c.m[a] = 0
	}
	return 4
}

func (c *ICC) ab(pm []int) int {
	c.relBase += c.pv(1, pm)
	return 2
}

func (c *ICC) pv(p int, pm []int) int {
	v := c.m[c.ip+p]
	switch pm[p-1] {
	case 0:
		return c.m[v]
	case 1:
		return v
	case 2:
		return c.m[c.relBase+v]
	default:
		panic(fmt.Sprint("invalid parameter mode:", pm[p-1]))
	}
}

func (c *ICC) pa(p int, pm []int) int {
	a := c.m[c.ip+p]
	switch pm[p-1] {
	case 0:
		return a
	case 2:
		return c.relBase + a
	default:
		panic(fmt.Sprint("invalid address parameter mode:", pm[p-1]))
	}
}
