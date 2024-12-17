package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/toaster/advent_of_code/2024/17/cc"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	c := cc.Parse(io.ReadLines())
	out := make(chan int)
	go c.Run(out)
	var output []string
	for value := range out {
		output = append(output, strconv.Itoa(value))
	}
	fmt.Println(strings.Join(output, ","))
}
