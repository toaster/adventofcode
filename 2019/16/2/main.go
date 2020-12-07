package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/toaster/advent_of_code/2019/16/fft"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	var part []int
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "") {
		i, _ := strconv.Atoi(s)
		part = append(part, i)
	}
	fmt.Println("read part of length", len(part))
	fmt.Println(fft.PerformHuge(part, 100))
}
