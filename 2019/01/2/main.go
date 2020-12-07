package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)
var fuel int
	for scanner.Scan() {
		weight, _ := strconv.Atoi(scanner.Text())
		fuel +=computeFuel(weight)
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println("Needed fuel:", fuel)
}

func computeFuel(mass int) int {
	fuel := int(math.Floor(float64(mass) / 3.0)) - 2
	if fuel < 0 {
		return 0
	}
	return fuel + computeFuel(fuel)
}
