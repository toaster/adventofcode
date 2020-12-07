package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/toaster/advent_of_code/2019/icc"
)

// Solution 1:
// NOT A J
// NOT B T
// OR T J
// NOT C T
// OR T J
// AND D J
// WALK
// -> 1935750
//
// Solution 2:
// Input instructions:
// NOT A T
// NOT B J
// OR J T
// NOT C J
// OR J T
// AND D T
// NOT E J
// NOT J J
// OR H J
// AND T J
// RUN
//
// Running...
//
// damage: 1142830249
func main() {
	input, _ := ioutil.ReadFile(os.Args[1])
	program := icc.Parse(string(input))
	in := make(chan int, 10)
	out := make(chan int, 1000)
	go func() {
		for {
			i := <-out
			if i > 255 {
				fmt.Println("damage:", i)
			} else {
				fmt.Print(string(rune(i)))
			}
		}
	}()
	go func() {
		b := make([]byte, 1)
		for {
			os.Stdin.Read(b)
			in <- int(b[0])
		}
	}()
	c := icc.New(in, out)
	c.Load(program)
	c.Run()
	time.Sleep(1 * time.Second)
}
