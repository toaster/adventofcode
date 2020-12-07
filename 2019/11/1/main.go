package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/toaster/advent_of_code/2019/icc"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	p := icc.Parse(string(input))
	in := make(chan int)
	out := make(chan int)
	tok1 := make(chan bool)
	tok2 := make(chan bool)
	c := icc.New(in, out)
	c.Load(p)

	m := map[pos]int{}
	cp := pos{}
	dir := 0

	go func() {
		for {
			<-tok1
			m[cp] = <-out
			dir = (4 + dir - 1 + 2*<-out) % 4
			switch dir {
			case 0:
				cp.y--
			case 1:
				cp.x++
			case 2:
				cp.y++
			case 3:
				cp.x--
			default:
				panic("unexpected direction")
			}
			tok2 <- true
		}
	}()

	go func() {
		for {
			in <- m[cp]
			tok1 <- true
			<-tok2
		}
	}()

	c.Run()

	fmt.Println(len(m))
}

type pos struct {
	x int
	y int
}
