package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	mem := io.ReadAll()
	result := 0
	enabled := true
	fmt.Println(mem)
	for i := 0; i < len(mem)-7; i++ {
		fmt.Println(mem[i : i+7])
		if mem[i:i+4] == "do()" {
			enabled = true
			i += 3
		} else if mem[i:i+7] == "don't()" {
			enabled = false
			i += 6
		} else if mem[i:i+4] == "mul(" {
			i += 4
			j := i
			v1 := 0
			for mem[i] >= '0' && mem[i] <= '9' {
				v1 = v1*10 + int(mem[i]-'0')
				i++
			}
			if j == i || mem[i] != ',' {
				i--
				continue
			}
			i++
			j = i
			v2 := 0
			for mem[i] >= '0' && mem[i] <= '9' {
				v2 = v2*10 + int(mem[i]-'0')
				i++
			}
			if j == i || mem[i] != ')' {
				i--
				continue
			}
			if enabled {
				result += v1 * v2
			}
		}
	}
	fmt.Println(result)
}
