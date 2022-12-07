package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2022/elfdevice"
)

func main() {
	fsSize := 70000000
	neededSpace := 30000000
	root := elfdevice.ParseShellSession()
	freeSpace := fsSize - root.Size
	minToBeCleared := neededSpace - freeSpace
	elfdevice.PrintNode(root, 0)
	fmt.Println(elfdevice.SmallestDirWithSizeAtLeast(root, minToBeCleared).Size)
}
