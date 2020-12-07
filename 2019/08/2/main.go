package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/toaster/advent_of_code/2019/08/sif"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	fmt.Println(sif.RenderImage(25, 6, strings.TrimSpace(string(input))))
}
