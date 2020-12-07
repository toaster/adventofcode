package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/toaster/advent_of_code/2019/icc"
)

func main() {
	count := 50
	input, _ := ioutil.ReadAll(os.Stdin)
	program := icc.Parse(string(input))
	cin := make([]chan int, 0, count)
	cout := make([]chan int, 0, count)
	hub := make(chan msg, 1000)
	comps := make([]*icc.ICC, 0, count)
	// setup
	for i := 0; i < count; i++ {
		cin = append(cin, make(chan int, 1000))
		cout = append(cout, make(chan int, 1000))
		c := icc.New(cin[i], cout[i])
		c.Load(program)
		comps = append(comps, c)
		go c.Run()
		go handle(cout[i], hub)
	}
	// boot
	for i := 0; i < count; i++ {
		cin[i] <- i
	}
	// hub/nat
	// curiously only works with the output enabled (probably because of race condition) and delivers 19641
	var natMsg msg
	var idle int
	var lastNatY int
	for {
		select {
		case m := <-hub:
			idle = 0
			if m.addr == 255 {
				natMsg = m
				fmt.Println("NAT", m)
			}
			for i := 0; i < count; i++ {
				if i == m.addr {
					cin[i] <- m.x
					cin[i] <- m.y
				} else {
					cin[i] <- -1
				}
			}
			fmt.Println("TX")
		default:
			fmt.Print(".")
			idle++
			if (idle == 100 && natMsg != msg{}) {
				fmt.Println()
				if lastNatY == natMsg.y {
					fmt.Println(lastNatY)
					return
				}
				fmt.Println("PING", natMsg)
				lastNatY = natMsg.y
				cin[0] <- natMsg.x
				cin[0] <- natMsg.y
			}
			for i := 0; i < count; i++ {
				cin[i] <- -1
			}
		}
	}
}

func handle(out chan int, hub chan msg) {
	for {
		select {
		case a := <-out:
			hub <- msg{addr: a, x: <-out, y: <-out}
		}
	}
}

type msg struct {
	addr int
	x    int
	y    int
}
