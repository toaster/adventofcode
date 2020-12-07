package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/toaster/advent_of_code/2019/07/1/icc"
)

func main() {
	program, _ := ioutil.ReadAll(os.Stdin)
	s, _ := icc.ComputeOptimalThrusterConfig(string(program))
	fmt.Println(s)
}
