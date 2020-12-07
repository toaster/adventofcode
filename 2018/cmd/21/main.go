package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var registers []int
var instructions map[string]fnType
var ip reference

type command struct {
	name      string
	i1, i2, o int
}

type fnInput interface {
	value() int
}

type reference int

func (r reference) value() int {
	if int(r) < 0 || len(registers) <= int(r) {
		panic(fmt.Sprintf("invalid register reference: %d", r))
	}
	return registers[int(r)]
}

func (r reference) setValue(v int) {
	if int(r) < 0 || len(registers) <= int(r) {
		panic(fmt.Sprintf("invalid register reference: %d", r))
	}
	registers[int(r)] = v
}

type immediate int

func (i immediate) value() int {
	return int(i)
}

type fnType func(int, int, int)

func init() {
	registers = make([]int, 6)
	instructions = map[string]fnType{}
	instructions["addr"] = func(i1, i2, o int) {
		a := reference(i1).value()
		b := reference(i2).value()
		c := reference(o)
		c.setValue(a + b)
	}
	instructions["addi"] = func(i1, i2, o int) {
		a := reference(i1).value()
		b := immediate(i2).value()
		c := reference(o)
		c.setValue(a + b)
	}
	instructions["mulr"] = func(i1, i2, o int) {
		a := reference(i1).value()
		b := reference(i2).value()
		c := reference(o)
		c.setValue(a * b)
	}
	instructions["muli"] = func(i1, i2, o int) {
		a := reference(i1).value()
		b := immediate(i2).value()
		c := reference(o)
		c.setValue(a * b)
	}
	instructions["banr"] = func(i1, i2, o int) {
		a := reference(i1).value()
		b := reference(i2).value()
		c := reference(o)
		c.setValue(a & b)
	}
	instructions["bani"] = func(i1, i2, o int) {
		a := reference(i1).value()
		b := immediate(i2).value()
		c := reference(o)
		c.setValue(a & b)
	}
	instructions["borr"] = func(i1, i2, o int) {
		a := reference(i1).value()
		b := reference(i2).value()
		c := reference(o)
		c.setValue(a | b)
	}
	instructions["bori"] = func(i1, i2, o int) {
		a := reference(i1).value()
		b := immediate(i2).value()
		c := reference(o)
		c.setValue(a | b)
	}
	instructions["setr"] = func(i, _, o int) {
		a := reference(i).value()
		c := reference(o)
		c.setValue(a)
	}
	instructions["seti"] = func(i, _, o int) {
		a := immediate(i).value()
		c := reference(o)
		c.setValue(a)
	}
	instructions["gtir"] = func(i1, i2, o int) {
		a := immediate(i1).value()
		b := reference(i2).value()
		c := reference(o)
		c.setValue(btoi(a > b))
	}
	instructions["gtri"] = func(i1, i2, o int) {
		a := reference(i1).value()
		b := immediate(i2).value()
		c := reference(o)
		c.setValue(btoi(a > b))
	}
	instructions["gtrr"] = func(i1, i2, o int) {
		a := reference(i1).value()
		b := reference(i2).value()
		c := reference(o)
		c.setValue(btoi(a > b))
	}
	instructions["eqir"] = func(i1, i2, o int) {
		a := immediate(i1).value()
		b := reference(i2).value()
		c := reference(o)
		c.setValue(btoi(a == b))
	}
	instructions["eqri"] = func(i1, i2, o int) {
		a := reference(i1).value()
		b := immediate(i2).value()
		c := reference(o)
		c.setValue(btoi(a == b))
	}
	instructions["eqrr"] = func(i1, i2, o int) {
		a := reference(i1).value()
		b := reference(i2).value()
		c := reference(o)
		c.setValue(btoi(a == b))
	}
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() {
	inputFile := os.Args[1]
	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSuffix(string(b), "\n"), "\n")

	ipRegex := regexp.MustCompile("#ip (\\d+)")
	m := ipRegex.FindStringSubmatch(lines[0])
	ip = reference(atoi(m[1]))

	arg := 0
	if len(os.Args) > 2 {
		arg = atoi(os.Args[2])
	}
	cmds := []command{}
	cmdRegex := regexp.MustCompile("([a-z]+) (\\d+) (\\d+) (\\d+)")
	copy(registers, []int{arg, 0, 0, 0, 0, 0})
	for i := 1; i < len(lines); i++ {
		m := cmdRegex.FindStringSubmatch(lines[i])
		cmds = append(cmds, command{m[1], atoi(m[2]), atoi(m[3]), atoi(m[4])})
	}

	count := 0
	values := map[int]int{}
	lastValue := 0
	for {
		i := ip.value()
		if i < 0 || i >= len(cmds) {
			break
		}
		cmd := cmds[i]
		count++
		print := false //count%10000000 == 0 || ip.value() > 27 && registers[4] < 0
		if ip.value() == 28 {
			v := registers[4]
			if len(values) == 0 {
				fmt.Println("first stop", v)
			}
			if values[v] > 0 {
				fmt.Println("last stop", lastValue)
				break
			}
			values[v]++
			lastValue = v
		}
		if print {
			fmt.Printf(
				"ip=%d [%d, %d, %d, %d, %d, %d] %s %d %d %d",
				i,
				registers[0],
				registers[1],
				registers[2],
				registers[3],
				registers[4],
				registers[5],
				cmd.name,
				cmd.i1,
				cmd.i2,
				cmd.o)
		}
		instructions[cmd.name](cmd.i1, cmd.i2, cmd.o)
		if print {
			fmt.Printf(
				" [%d, %d, %d, %d, %d, %d]\n",
				registers[0],
				registers[1],
				registers[2],
				registers[3],
				registers[4],
				registers[5])
		}
		//time.Sleep(1 * time.Second)
		ip.setValue(ip.value() + 1)
	}

	fmt.Println("Instructions:", count)
	fmt.Println("Registers:", registers)
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func min(values map[int]int) int {
	m := -1
	for v := range values {
		if v < 0 {
			continue
		}
		if m < 0 || v < m {
			m = v
		}
	}
	return m
}
