package main

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	channels := map[string]chan float64{}
	mutex := &sync.RWMutex{}
	lookup := func(name string) chan float64 {
		mutex.Lock()
		defer mutex.Unlock()
		if c := channels[name]; c != nil {
			return c
		}

		c := make(chan float64)
		channels[name] = c
		return c
	}

	var root1, root2 chan float64
	for _, line := range io.ReadLines() {
		name := line[0:4]
		if name == "humn" {
			continue
		}

		remainder := line[6:]
		c := lookup(name)
		if len(line) > 10 && line[10] == ' ' {
			parts := strings.Split(remainder, " ")
			c1 := lookup(parts[0])
			c2 := lookup(parts[2])
			if name == "root" {
				root1 = c1
				root2 = c2
			} else {
				switch parts[1] {
				case "+":
					go func(out, in1, in2 chan float64) {
						for {
							c <- <-c1 + <-c2
						}
					}(c, c1, c2)
				case "-":
					go func(out, in1, in2 chan float64) {
						for {
							c <- <-c1 - <-c2
						}
					}(c, c1, c2)
				case "*":
					go func(out, in1, in2 chan float64) {
						for {
							c <- <-c1 * <-c2
						}
					}(c, c1, c2)
				case "/":
					go func(out, in1, in2 chan float64) {
						for {
							c <- <-c1 / <-c2
						}
					}(c, c1, c2)
				}
			}
		} else {
			go func(out chan float64, value float64) {
				for {
					out <- value
				}
			}(c, float64(io.ParseInt(remainder)))
		}
	}
	me := lookup("humn")
	me <- 0
	diff0 := math.Abs(<-root1 - <-root2)
	if diff0 == 0 {
		fmt.Println(0)
		return
	}

	me <- 1
	diff1 := math.Abs(<-root1 - <-root2)
	if diff1 == 0 {
		fmt.Println(1)
		return
	}

	step := diff1 - diff0
	var x float64
	if step > 0 {
		x = diff1/step + 1
	} else {
		x = -(diff0 / step)
	}
	fmt.Printf("check %20f\n", x)
	me <- x
	fmt.Println("=>", <-root1, "==", <-root2)
}
