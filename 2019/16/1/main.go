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
	var signal []int
	for _, s := range strings.Split(string(input), "") {
		i, _ := strconv.Atoi(s)
		signal = append(signal, i)
	}
	fmt.Println(fft.Perform(signal, 100))
}
