package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	action rune
	amount int
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
		os.Exit(1)
	}

	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	mem := map[int]uint64{}
	var mask []rune
	for _, line := range lines {
		cmdAndValue := strings.Split(line, " = ")
		cmd := cmdAndValue[0]
		value := cmdAndValue[1]
		if cmd == "mask" {
			mask = []rune(value)
		} else {
			address, _ := strconv.Atoi(cmd[4 : len(cmd)-1])
			v, _ := strconv.Atoi(value)
			mem[address] = applyMask(mask, mem[address], uint64(v))
		}
	}
	var result uint64
	for _, f := range mem {
		result += f
	}
	fmt.Println("result:", result)
}

func applyMask(mask []rune, cur, new uint64) uint64 {
	var result uint64
	for i := 0; i < 36; i++ {
		field := mask[35-i]
		if field == 'X' {
			result = result | (new & (1 << i))
		} else {
			var bit uint64
			if field == '1' {
				bit = 1
			}
			result = result | (bit << i)
		}
	}
	return result
}
