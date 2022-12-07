package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2022/elfdevice"
)

func main() {
	root := elfdevice.ParseShellSession()
	elfdevice.PrintNode(root, 0)
	fmt.Println(elfdevice.TreeSize(root, 100000))
}
