package main

import (
	"fmt"
	"io/ioutil"
	"os"

	space2 "github.com/toaster/advent_of_code/2019/space"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	s := space2.ParseSystem(string(input))
	fmt.Println(s.ComputePeriod())
}
