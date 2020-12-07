package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"

	"github.com/toaster/advent_of_code/2019/icc"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	program := icc.Parse(string(input))
	in := make(chan int)
	out := make(chan int)
	tok1 := make(chan bool)
	tok2 := make(chan bool)
	c := icc.New(in, out)
	c.Load(program)

	m := map[pos]int{}
	cp := pos{}
	dir := 0
	m[cp] = 1

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

	sx := math.MaxInt64
	sy := math.MaxInt64
	mx := math.MinInt64
	my := math.MinInt64

	for p := range m {
		if p.x < sx {
			sx = p.x
		}
		if p.x > mx {
			mx = p.x
		}
		if p.y < sy {
			sy = p.y
		}
		if p.y > my {
			my = p.y
		}
	}

	for y := sy; y <= my; y++ {
		for x := sx; x <= mx; x++ {
			if m[pos{x, y}] > 0 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	fmt.Println(sx, sy, mx, my, len(m))
}

type pos struct {
	x int
	y int
}
