package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/toaster/advent_of_code/2019/icc"
)

func main() {
	input, _ := ioutil.ReadFile(os.Args[1])
	joyInput, _ := ioutil.ReadFile(os.Args[2])
	var knownMoves []int
	for _, s := range strings.Split(strings.TrimSpace(string(joyInput)), "\n") {
		mv, _ := strconv.Atoi(s)
		knownMoves = append(knownMoves, mv)
	}
	p := icc.Parse(string(input))
	in := make(chan int, 2)
	out := make(chan int)
	done := make(chan bool)
	ack := make(chan bool)
	c := icc.New(in, out)
	c.Load(p)

	m := map[pos]int{}
	rd := bufio.NewReader(os.Stdin)
	var step, lastMove int

	// joystick
	// var joy int
	// go func() {
	// 	for {
	// 		fmt.Println("write", joy)
	// 		in <- joy
	// 	}
	// }()

	// output
	go func() {
		for {
			// var bx int
			select {
			case <-done:
				ack <- true
				break
			case x := <-out:
				y := <-out
				tile := <-out
				// if tile == 4 {
				// 	bx = x
				// }
				m[pos{x, y}] = tile
				// move joystick with the ball :)
				if tile == 4 {
					// if x < bx {
					// 	in <- 1
					// } else if x > bx {
					// 	in <- -1
					// } else {
					// 	in <- 0
					// }
					printScreen(m, step, lastMove)
					if step < len(knownMoves) {
						lastMove = knownMoves[step]
					} else {
						fmt.Print("MOVE!> ")
						i, _ := rd.ReadString('\n')
						lastMove, _ = strconv.Atoi(strings.TrimSpace(i))
					}
					in <- lastMove
					step++
				}
			}
		}
	}()

	// insert coins
	c.Patch(0, 2)

	c.Run()

	done <- true
	<-ack

	printScreen(m, step, lastMove)
}

func printScreen(m map[pos]int, step, move int) {
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
	fmt.Println("score:", m[pos{-1, 0}], "blocks:", bc, "step:", step, "last move:", move)
}

type pos struct {
	x int
	y int
}
