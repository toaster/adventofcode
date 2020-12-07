package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var fuel int
	for scanner.Scan() {
		weight, _ := strconv.Atoi(scanner.Text())
		fuel += int(math.Floor(float64(weight)/3.0)) - 2
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println("Needed fuel:", fuel)
}
