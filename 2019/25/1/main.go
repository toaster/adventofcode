package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/toaster/advent_of_code/2019/icc"
)

func main() {
	input, _ := ioutil.ReadFile(os.Args[1])
	program := icc.Parse(string(input))
	in := make(chan int, 1000)
	out := make(chan int, 1000)
	c := icc.New(in, out)
	c.Load(program)

	rd := bufio.NewReader(os.Stdin)

	go c.Run()

	var buf []rune
	for {
		select {
		case x := <-out:
			if x == '\n' {
				fmt.Println(string(buf))
				if string(buf) == "Command?" {
					i, _ := rd.ReadString('\n')
					for _, r := range i {
						in <- int(r)
					}
				}
				buf = buf[:0]
			} else {
				buf = append(buf, rune(x))
			}
		}
	}
}
