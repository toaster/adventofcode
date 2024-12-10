package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2024/09/amphipod"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	disk := amphipod.InitializeDisk(io.ReadLines()[0])
	// disk.Print()
	disk.CompactFiles()
	// disk.Print()
	fmt.Println(disk.Checksum())
}
