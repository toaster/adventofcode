package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	channels := map[string]chan int{}
	mutex := &sync.RWMutex{}
	lookup := func(name string) chan int {
		mutex.Lock()
		defer mutex.Unlock()
		if c := channels[name]; c != nil {
			return c
		}

		c := make(chan int)
		channels[name] = c
		return c
	}

	for _, line := range io.ReadLines() {
		name := line[0:4]
		remainder := line[6:]
		if len(line) > 10 && line[10] == ' ' {
			parts := strings.Split(remainder, " ")
			go func(name, op, key1, key2 string) {
				c := lookup(name)
				var value int
				v1 := <-lookup(key1)
				v2 := <-lookup(key2)
				switch op {
				case "+":
					value = v1 + v2
				case "-":
					value = v1 - v2
				case "*":
					value = v1 * v2
				case "/":
					value = v1 / v2
				}
				c <- value
			}(name, parts[1], parts[0], parts[2])
		} else {
			go func(name string, value int) {
				c := lookup(name)
				c <- value
			}(name, io.ParseInt(remainder))
		}
	}
	fmt.Println(<-lookup("root"))
}
