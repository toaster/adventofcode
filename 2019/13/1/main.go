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
	done := make(chan bool)
	ack := make(chan bool)
	c := icc.New(in, out)
	c.Load(p)

	m := map[pos]int{}

	go func() {
		for {
			select {
			case <-done:
				ack <- true
				break
			case x := <-out:
				y := <-out
				tile := <-out
				m[pos{x, y}] = tile
			}
		}
	}()

	c.Run()

	done <- true
	<-ack

	var mx, my int
	for p := range m {
		if p.x > mx {
			mx = p.x
		}
		if p.y > my {
			my = p.y
		}
	}
	var bc int
	for y := 0; y <= my; y++ {
		for x := 0; x <= mx; x++ {
			switch m[pos{x, y}] {
			case 1:
				fmt.Print("#")
			case 2:
				bc++
				fmt.Print("")
			case 3:
				fmt.Print("—")
			case 4:
				fmt.Print("•")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println(bc)
}

type pos struct {
	x int
	y int
}
