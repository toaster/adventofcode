package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/toaster/advent_of_code/2019/icc"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	program := icc.Parse(string(input))
	in := make(chan int, 200)
	out := make(chan int)

	c := icc.New(in, out)
	c.Load(program)
	c.Patch(0, 2)
	go func() {
		for {
			i := <-out
			if i < 255 {
				fmt.Print(string(rune(i)))
			} else {
				fmt.Println(i)
				break
			}
		}
	}()

	mvP := "A,B,A,C,B,C,B,C,A,C\n"
	mvA := "R,12,L,6,R,12\n"
	mvB := "L,8,L,6,L,10\n"
	mvC := "R,12,L,10,L,6,R,10\n"
	feed := "n\n"
	for _, i := range mvP {
		in <- int(i)
	}
	for _, i := range mvA {
		in <- int(i)
	}
	for _, i := range mvB {
		in <- int(i)
	}
	for _, i := range mvC {
		in <- int(i)
	}
	for _, i := range feed {
		in <- int(i)
	}
	c.Run()
}
