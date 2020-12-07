package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/toaster/advent_of_code/2019/nanofactory"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	r := nanofactory.Parse(string(input))
	fmt.Println(r.ComputeOreRequirementForOneFuel())
}
