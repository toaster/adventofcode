package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var registers []int
var instructions map[string]fnType

type fnInput interface {
	value() (int, error)
}

type reference int

func (r reference) value() (int, error) {
	if len(registers) > int(r) {
		return registers[int(r)], nil
	}
	return 0, fmt.Errorf("invalid register reference: %d", r)
}

func (r reference) setValue(v int) error {
	if len(registers) > int(r) {
		registers[int(r)] = v
		return nil
	}
	return fmt.Errorf("invalid register reference: %d", r)
}

type immediate int

func (i immediate) value() (int, error) {
	return int(i), nil
}

type fnType func(int, int, int) error

func init() {
	registers = make([]int, 4)
	instructions = map[string]fnType{}
	instructions["addr"] = func(i1, i2, o int) error {
		a, err := reference(i1).value()
		if err != nil {
			return err
		}
		b, err := reference(i2).value()
		if err != nil {
			return err
		}
		c := reference(o)
		if err := c.setValue(a + b); err != nil {
			return err
		}
		return nil
	}
	instructions["addi"] = func(i1, i2, o int) error {
		a, err := reference(i1).value()
		if err != nil {
			return err
		}
		b, err := immediate(i2).value()
		if err != nil {
			return err
		}
		c := reference(o)
		if err := c.setValue(a + b); err != nil {
			return err
		}
		return nil
	}
	instructions["mulr"] = func(i1, i2, o int) error {
		a, err := reference(i1).value()
		if err != nil {
			return err
		}
		b, err := reference(i2).value()
		if err != nil {
			return err
		}
		c := reference(o)
		if err := c.setValue(a * b); err != nil {
			return err
		}
		return nil
	}
	instructions["muli"] = func(i1, i2, o int) error {
		a, err := reference(i1).value()
		if err != nil {
			return err
		}
		b, err := immediate(i2).value()
		if err != nil {
			return err
		}
		c := reference(o)
		if err := c.setValue(a * b); err != nil {
			return err
		}
		return nil
	}
	instructions["banr"] = func(i1, i2, o int) error {
		a, err := reference(i1).value()
		if err != nil {
			return err
		}
		b, err := reference(i2).value()
		if err != nil {
			return err
		}
		c := reference(o)
		if err := c.setValue(a & b); err != nil {
			return err
		}
		return nil
	}
	instructions["bani"] = func(i1, i2, o int) error {
		a, err := reference(i1).value()
		if err != nil {
			return err
		}
		b, err := immediate(i2).value()
		if err != nil {
			return err
		}
		c := reference(o)
		if err := c.setValue(a & b); err != nil {
			return err
		}
		return nil
	}
	instructions["borr"] = func(i1, i2, o int) error {
		a, err := reference(i1).value()
		if err != nil {
			return err
		}
		b, err := reference(i2).value()
		if err != nil {
			return err
		}
		c := reference(o)
		if err := c.setValue(a | b); err != nil {
			return err
		}
		return nil
	}
	instructions["bori"] = func(i1, i2, o int) error {
		a, err := reference(i1).value()
		if err != nil {
			return err
		}
		b, err := immediate(i2).value()
		if err != nil {
			return err
		}
		c := reference(o)
		if err := c.setValue(a | b); err != nil {
			return err
		}
		return nil
	}
	instructions["setr"] = func(i, _, o int) error {
		a, err := reference(i).value()
		if err != nil {
			return err
		}
		c := reference(o)
		if err := c.setValue(a); err != nil {
			return err
		}
		return nil
	}
	instructions["seti"] = func(i, _, o int) error {
		a, err := immediate(i).value()
		if err != nil {
			return err
		}
		c := reference(o)
		if err := c.setValue(a); err != nil {
			return err
		}
		return nil
	}
	instructions["gtir"] = func(i1, i2, o int) error {
		a, err := immediate(i1).value()
		if err != nil {
			return err
		}
		b, err := reference(i2).value()
		if err != nil {
			return err
		}
		c := reference(o)
		if err := c.setValue(btoi(a > b)); err != nil {
			return err
		}
		return nil
	}
	instructions["gtri"] = func(i1, i2, o int) error {
		a, err := reference(i1).value()
		if err != nil {
			return err
		}
		b, err := immediate(i2).value()
		if err != nil {
			return err
		}
		c := reference(o)
		if err := c.setValue(btoi(a > b)); err != nil {
			return err
		}
		return nil
	}
	instructions["gtrr"] = func(i1, i2, o int) error {
		a, err := reference(i1).value()
		if err != nil {
			return err
		}
		b, err := reference(i2).value()
		if err != nil {
			return err
		}
		c := reference(o)
		if err := c.setValue(btoi(a > b)); err != nil {
			return err
		}
		return nil
	}
	instructions["eqir"] = func(i1, i2, o int) error {
		a, err := immediate(i1).value()
		if err != nil {
			return err
		}
		b, err := reference(i2).value()
		if err != nil {
			return err
		}
		c := reference(o)
		if err := c.setValue(btoi(a == b)); err != nil {
			return err
		}
		return nil
	}
	instructions["eqri"] = func(i1, i2, o int) error {
		a, err := reference(i1).value()
		if err != nil {
			return err
		}
		b, err := immediate(i2).value()
		if err != nil {
			return err
		}
		c := reference(o)
		if err := c.setValue(btoi(a == b)); err != nil {
			return err
		}
		return nil
	}
	instructions["eqrr"] = func(i1, i2, o int) error {
		a, err := reference(i1).value()
		if err != nil {
			return err
		}
		b, err := reference(i2).value()
		if err != nil {
			return err
		}
		c := reference(o)
		if err := c.setValue(btoi(a == b)); err != nil {
			return err
		}
		return nil
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

	beforeRegex := regexp.MustCompile("Before: +\\[(\\d+), (\\d+), (\\d+), (\\d+)\\]")
	cmdRegex := regexp.MustCompile("(\\d+) (\\d+) (\\d+) (\\d+)")
	afterRegex := regexp.MustCompile("After: +\\[(\\d+), (\\d+), (\\d+), (\\d+)\\]")

	multiCount := 0
	ccsByOpcode := map[int][]string{}
	blocks := strings.Split(strings.TrimSuffix(string(b), "\n"), "\n\n\n")
	lines := strings.Split(blocks[0], "\n")
	for i := 0; i < len(lines); i++ {
		if m := beforeRegex.FindStringSubmatch(lines[i]); m != nil {
			before := []int{atoi(m[1]), atoi(m[2]), atoi(m[3]), atoi(m[4])}
			i++
			m = cmdRegex.FindStringSubmatch(lines[i])
			cmd := [4]int{atoi(m[1]), atoi(m[2]), atoi(m[3]), atoi(m[4])}
			i++
			m = afterRegex.FindStringSubmatch(lines[i])
			after := []int{atoi(m[1]), atoi(m[2]), atoi(m[3]), atoi(m[4])}
			possibleCmds := []string{}
			for c, fn := range instructions {
				copy(registers, before)
				fn(cmd[1], cmd[2], cmd[3])
				if equal(registers, after) {
					possibleCmds = append(possibleCmds, c)
				}
			}
			if len(possibleCmds) > 2 {
				multiCount++
			}
			if ccsByOpcode[cmd[0]] == nil {
				ccsByOpcode[cmd[0]] = possibleCmds
			} else {
				ccsByOpcode[cmd[0]] = intersection(possibleCmds, ccsByOpcode[cmd[0]])
			}
		}
	}
	cmdsByOpcode := map[int]string{}
	for done := false; !done; {
		done = true
		for oc, cmds := range ccsByOpcode {
			if len(cmds) == 0 {

			} else if len(cmds) == 1 {
				cmdsByOpcode[oc] = cmds[0]
			} else {
				done = false
				ccsByOpcode[oc] = difference(ccsByOpcode[oc], values(cmdsByOpcode))
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
	lines = strings.Split(blocks[1], "\n")
	copy(registers, []int{0, 0, 0, 0})
	for i := 0; i < len(lines); i++ {
		if m := cmdRegex.FindStringSubmatch(lines[i]); m != nil {
			cmd := [4]int{atoi(m[1]), atoi(m[2]), atoi(m[3]), atoi(m[4])}
			instructions[cmdsByOpcode[cmd[0]]](cmd[1], cmd[2], cmd[3])
		}
	}
	fmt.Println(cmdsByOpcode)
	fmt.Println("Multis:", multiCount)
	fmt.Println("Registers:", registers)
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func equal(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}

func intersection(s1, s2 []string) []string {
	r := []string{}
	for _, v := range s1 {
		for _, w := range s2 {
			if v == w {
				r = append(r, v)
			}
		}
	}
	return r
}

func difference(s1, s2 []string) []string {
	r := []string{}
	for _, v := range s1 {
		keep := true
		for _, w := range s2 {
			if v == w {
				keep = false
				break
			}
		}
		if keep {
			r = append(r, v)
		}
	}
	return r
}

func values(m map[int]string) []string {
	r := []string{}
	for _, v := range m {
		r = append(r, v)
	}
	return r
}
