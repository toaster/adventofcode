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
	mem := map[uint64]uint64{}
	// var mask string
	// var filterMask string
	// var maskFactor uint64
	// for i := len(lines) - 1; i >= 0; i-- {
	// 	cmdAndValue := strings.Split(lines[i], " = ")
	// 	cmd := cmdAndValue[0]
	// 	if cmd == "mask" || mask == "" {
	// 		for j := i; j >= 0; j-- {
	// 			if strings.HasPrefix(lines[j], "mask = ") {
	// 				mask = lines[j][7:]
	// 				if filterMask == "" {
	// 					filterMask = mask
	// 					maskFactor = uint64(strings.Count(filterMask, "X"))
	// 				} else {
	// 					maskFactor = uint64(strings.Count(filterMask, "X")) - maskFactor
	// 				}
	// 				break
	// 			}
	// 		}
	// 		continue
	// 	}
	// 	address, _ := strconv.Atoi(cmd[4 : len(cmd)-1])
	// 	value, _ := strconv.Atoi(cmdAndValue[1])
	// 	mem[address] = applyMask(mask, mem[address], uint64(value))
	// }
	var mask []rune
	for _, line := range lines {
		cmdAndValue := strings.Split(line, " = ")
		cmd := cmdAndValue[0]
		value := cmdAndValue[1]
		if cmd == "mask" {
			mask = []rune(value)
		} else {
			a, _ := strconv.Atoi(cmd[4 : len(cmd)-1])
			v, _ := strconv.Atoi(value)
			for _, address := range applyMask(mask, uint64(a)) {
				mem[address] = uint64(v)
			}
		}
	}
	var result uint64
	for _, f := range mem {
		result += f
	}
	fmt.Println("result:", result)
}

func applyMask(mask []rune, address uint64) []uint64 {
	var template uint64
	var floating []uint64
	for i := 0; i < 36; i++ {
		field := mask[35-i]
		if field == 'X' {
			floating = append(floating, uint64(i))
		} else {
			var num uint64
			num = 1 << i
			if field == '0' {
				num = num & address
			}
			template = template | num
		}
	}
	var addresses []uint64
	variantCount := 1 << len(floating)
	for variant := 0; variant < variantCount; variant++ {
		a := template
		for variantBit, addressBit := range floating {
			bit := uint64(variant) & (1 << variantBit)
			a = a | (bit << (addressBit - uint64(variantBit)))
		}
		addresses = append(addresses, a)
	}
	return addresses
}
