package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/toaster/advent_of_code/2019/icc"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	boostProgram := icc.Parse(string(input))
	in := make(chan int, 10)
	out := make(chan int, 50)
	in <- 2
	go func() {
		for {
			fmt.Println(">", <-out)
		}
	}()
	c := icc.New(in, out)
	c.Load(boostProgram)
	c.Run()
	// wait for output
	time.Sleep(100 * time.Millisecond)
}
