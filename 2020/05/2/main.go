package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
	}
	passes := strings.Split(strings.Trim(string(input), "\n"), "\n")
	seatIDs := map[uint]bool{}
	for _, pass := range passes {
		seatIDs[computeSeatID(pass)] = true
	}
	var prev uint
	var id uint
	for ; ; id++ {
		if seatIDs[id] {
			prev = id
			continue
		}
		if prev == 0 {
			continue
		}
		break
	}
	fmt.Println("my seat ID:", id)
}

func computeSeatID(pass string) uint {
	rowPart := pass[:7]
	rowPart = strings.ReplaceAll(rowPart, "F", "0")
	rowPart = strings.ReplaceAll(rowPart, "B", "1")
	row := parseBinary(rowPart)
	colPart := pass[7:]
	colPart = strings.ReplaceAll(colPart, "L", "0")
	colPart = strings.ReplaceAll(colPart, "R", "1")
	col := parseBinary(colPart)
	return row*8 + col
}

func parseBinary(value string) uint {
	val, err := strconv.ParseUint(value, 2, 8)
	if err != nil {
		panic(fmt.Sprint("cannot parse row:", err))
	}
	return uint(val)
}
