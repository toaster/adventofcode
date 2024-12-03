package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	mem := io.ReadAll()
	result := 0
	for i := 0; i < len(mem)-4; i++ {
		if mem[i:i+4] == "mul(" {
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
			result += v1 * v2
		}
	}
	fmt.Println(result)
}
