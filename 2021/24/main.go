package main

import (
	"fmt"
	"os"

	"github.com/toaster/advent_of_code/2021/24/alu"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	computer := alu.New(io.ReadFile(os.Args[1]))
	computer.Run()
	fmt.Println("state:", computer.State())
}
