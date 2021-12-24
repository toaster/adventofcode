package alu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

// New creates a new ALU computer.
func New(code []string) *ALU {
	a := &ALU{reader: bufio.NewReader(os.Stdin)}
	a.load(code)
	return a
}

// ALU is an arithmetic logic unit.
type ALU struct {
	inputCount   int
	instructions []*instruction
	reader       *bufio.Reader
	w            int
	x            int
	y            int
	z            int
}

// Run performs the program loaded into the ALU.
func (a *ALU) Run() {
	a.w = 0
	a.x = 0
	a.y = 0
	a.z = 0
	for _, op := range a.instructions {
		op.fn(op.arg1, op.arg2)
		fmt.Println("performed:", op.raw, "\t=>", a.State())
	}
}

// State returns the state of the ALU as string.
func (a *ALU) State() string {
	return fmt.Sprintf("w: %d, x: %d, y: %d, z: %d", a.w, a.x, a.y, a.z)
}

func (a *ALU) load(lines []string) {
	for _, line := range lines {
		a.instructions = append(a.instructions, a.parseLine(line))
	}
	return
}

func (a *ALU) add(i *int, valOrRef valueOrReference) {
	*i += valOrRef.value()
}

func (a *ALU) div(i *int, valOrRef valueOrReference) {
	*i /= valOrRef.value()
}

func (a *ALU) eql(i *int, valOrRef valueOrReference) {
	if *i == valOrRef.value() {
		*i = 1
	} else {
		*i = 0
	}
}

func (a *ALU) inp(i *int, _ valueOrReference) {
	a.inputCount++
	fmt.Println("state:", a.State())
	fmt.Printf("Enter value: %d >", a.inputCount)
	s, err := a.reader.ReadString('\n')
	io.ReportError("failed to read value from stdin", err)
	s = strings.Trim(s, "\n")
	*i, err = strconv.Atoi(s)
	io.ReportError("failed to parse input", err)
	fmt.Printf("read: %d\n", *i)
}

func (a *ALU) mod(i *int, valOrRef valueOrReference) {
	*i %= valOrRef.value()
}

func (a *ALU) mul(i *int, valOrRef valueOrReference) {
	*i *= valOrRef.value()
}

func (a *ALU) parseLine(line string) (i *instruction) {
	args := strings.Split(strings.Trim(strings.Split(line, "#")[0], " "), " ")
	i = &instruction{raw: strings.Join(args, " "), arg1: a.parseVar(args[1])}
	if len(args) > 2 {
		i.arg2 = a.parseVarOrValue(args[2])
	}
	switch args[0] {
	case "add":
		i.fn = a.add
	case "div":
		i.fn = a.div
	case "eql":
		i.fn = a.eql
	case "inp":
		i.fn = a.inp
	case "mod":
		i.fn = a.mod
	case "mul":
		i.fn = a.mul
	default:
		io.ReportError("", fmt.Errorf("failed to parse instruction: %s", line))
	}
	return
}

func (a *ALU) parseVar(s string) *int {
	switch s {
	case "w":
		return &a.w
	case "x":
		return &a.x
	case "y":
		return &a.y
	case "z":
		return &a.z
	default:
		io.ReportError("", fmt.Errorf("failed to parse variable: %s", s))
	}
	return nil
}

func (a *ALU) parseVarOrValue(s string) valueOrReference {
	if s[0] != '-' && (s[0] < '0' || s[0] > '9') {
		return valueOrReference{
			ref: a.parseVar(s),
		}
	}

	val, err := strconv.Atoi(s)
	io.ReportError("failed to parse value", err)
	return valueOrReference{val: val}
}

type executor func(*int, valueOrReference)

type instruction struct {
	raw  string
	fn   executor
	arg1 *int
	arg2 valueOrReference
}

type valueOrReference struct {
	ref *int
	val int
}

func (v *valueOrReference) value() int {
	if v.ref != nil {
		return *v.ref
	}

	return v.val
}
