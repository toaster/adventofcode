package main

import (
	"io/ioutil"
	"os"

	"github.com/toaster/advent_of_code/2019/05/2/icc"
)

func main() {
	program, _ := ioutil.ReadFile(os.Args[1])
	c := icc.New(os.Stdin, os.Stdout)
	c.Load(string(program))
	c.Run()
}
