package elfdevice

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

// VideoSystem represents a graphics processing unit of the elf device.
type VideoSystem struct {
	crt            []bool
	height         int
	program        []command
	signalStrength int
	spritePosition int
	width          int
	x              int
}

// LoadVideoProgram parses a set of VideoSystem instructions and load them into a freshly initialized VideoSystem.
func LoadVideoProgram(input []string) *VideoSystem {
	var program []command
	for _, line := range input {
		components := strings.Split(line, " ")
		arg := 0
		if len(components) > 1 {
			arg = io.ParseInt(components[1])
		}
		program = append(program, command{
			oc:  commandNameToOpcode[components[0]],
			arg: arg,
		})
	}
	return &VideoSystem{
		height:  6,
		program: program,
		width:   40,
	}
}

// PrintCRT prints the image of the CRT.
func (v *VideoSystem) PrintCRT() {
	for y := 0; y < v.height; y++ {
		for x := 0; x < v.width; x++ {
			crtPos := y*v.width + x
			if v.crt[crtPos] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

// Run runs the program loaded into the VideoSystem.
func (v *VideoSystem) Run() {
	offset := 20
	step := 40
	ticker := make(chan int)
	done := make(chan bool)
	v.crt = make([]bool, 40*6)
	v.x = 1
	v.signalStrength = 0
	go func() {
		for cycle := 1; ; cycle++ {
			if (cycle-offset)%step == 0 {
				v.signalStrength += cycle * v.x
			}
			ticker <- cycle
			crtPos := cycle - 1
			if crtPos%step >= v.x-1 && crtPos%step <= v.x+1 {
				v.crt[crtPos] = true
			}
			<-done
		}
	}()
	for _, cmd := range v.program {
		<-ticker
		done <- true
		switch cmd.oc {
		case addx:
			<-ticker
			v.x += cmd.arg
			done <- true
		}
	}
	close(done)
}

// SignalStrength returns the sum of the signal strengths.
func (v *VideoSystem) SignalStrength() int {
	return v.signalStrength
}

const (
	noop opcode = iota
	addx
)

type command struct {
	oc  opcode
	arg int
}

type opcode int

var commandNameToOpcode map[string]opcode

func init() {
	commandNameToOpcode = map[string]opcode{
		"addx": addx,
		"noop": noop,
	}
}
