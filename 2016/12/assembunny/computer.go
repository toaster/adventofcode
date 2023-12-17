package assembunny

import (
	"strings"
	"unicode"

	"github.com/toaster/advent_of_code/internal/io"
)

// Computer is a simple computer which executes Assembunny programs.
type Computer struct {
	registers map[string]int
	ip        int
	program   []*command
}

// NewComputer creates a new Assembunny computer.
func NewComputer() *Computer {
	return &Computer{registers: map[string]int{}}
}

// Load loads the given program into the computer's memory and resets the instruction pointer.
func (c *Computer) Load(input []string) {
	for _, cmd := range input {
		components := strings.Split(cmd, " ")
		var fn executor
		fn = parseFunction(components, fn, c)
		params := parseParams(components[1:])
		c.program = append(c.program, &command{fn, params})
	}
	c.ip = 0
}

// ReadRegister returns the contents of a register.
func (c *Computer) ReadRegister(register string) int {
	return c.registers[register]
}

// Run runs the loaded program.
func (c *Computer) Run() {
	for c.ip < len(c.program) {
		cmd := c.program[c.ip]
		cmd.executor(cmd.params)
	}
}

// WriteRegister writes a value into a register.
func (c *Computer) WriteRegister(register string, value int) {
	c.registers[register] = value
}

func (c *Computer) cpy(params []any) {
	c.registers[params[1].(string)] = c.getValue(params[0])
	c.ip++
}

func (c *Computer) dec(params []any) {
	c.registers[params[0].(string)]--
	c.ip++
}

func (c *Computer) getValue(reference any) int {
	value := 0
	if name, ok := reference.(string); ok {
		value = c.ReadRegister(name)
	} else {
		value = reference.(int)
	}
	return value
}

func (c *Computer) inc(params []any) {
	c.registers[params[0].(string)]++
	c.ip++
}

func (c *Computer) jnz(params []any) {
	if c.getValue(params[0]) != 0 {
		c.ip += params[1].(int)
	} else {
		c.ip++
	}
}

type command struct {
	executor
	params []any
}

type executor func(params []any)

func parseFunction(components []string, fn executor, c *Computer) executor {
	switch components[0] {
	case "cpy":
		fn = c.cpy
	case "inc":
		fn = c.inc
	case "dec":
		fn = c.dec
	case "jnz":
		fn = c.jnz
	}
	return fn
}

func parseParams(input []string) []any {
	var params []any
	for _, p := range input {
		if unicode.IsNumber(rune(p[0])) || p[0] == '-' {
			params = append(params, io.ParseInt(p))
		} else {
			params = append(params, p)
		}
	}
	return params
}
