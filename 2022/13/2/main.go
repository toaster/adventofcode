package main

import (
	"fmt"
	"sort"

	"github.com/toaster/advent_of_code/2022/elfdevice"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	divider1 := elfdevice.ParsePacket("[[2]]")
	divider2 := elfdevice.ParsePacket("[[6]]")
	packets := []elfdevice.Packet{divider1, divider2}
	for _, line := range io.ReadLines() {
		if line == "" {
			continue
		}
		packets = append(packets, elfdevice.ParsePacket(line))
	}
	sort.Slice(packets, func(i, j int) bool {
		return packets[i].Compare(packets[j]) == -1
	})
	decoderKey := 1
	for i, packet := range packets {
		if packet.Compare(divider1) == 0 || packet.Compare(divider2) == 0 {
			decoderKey *= i + 1
		}
	}
	fmt.Println(decoderKey)
}
