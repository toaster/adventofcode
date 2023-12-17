package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2016/12/assembunny"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	c := assembunny.NewComputer()
	c.Load(io.ReadLines())
	c.WriteRegister("c", 1)
	c.Run()
	fmt.Println(c.ReadRegister("a"))
}
