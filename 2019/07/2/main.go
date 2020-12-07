package main

import (
	"fmt"
	"io/ioutil"
	"os"

	icc2 "github.com/toaster/advent_of_code/2019/icc"
)

func main() {
	program, _ := ioutil.ReadAll(os.Stdin)
	s, _ := icc2.ComputeOptimalLoopedThrusterConfig(string(program))
	fmt.Println(s)
}
