package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/2024/23/ebhq"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	network := ebhq.NewNetwork()
	for _, line := range io.ReadLines() {
		ids := strings.Split(line, "-")
		network.AddConnection(ids[0], ids[1])
	}
	fmt.Println(network.ComputeLANPartyPassword())
}
