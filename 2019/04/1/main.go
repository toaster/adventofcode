package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	v := strings.Split(strings.TrimSpace(string(input)), "-")
	start, _ := strconv.Atoi(v[0])
	end, _ := strconv.Atoi(v[1])
	fmt.Println(countCandidatest(start, end))

}

func countCandidatest(start, end int) int {
	count := 0
	for k := 0; k < 10; k++ {
		for l := k; l < 10; l++ {
			for m := l; m < 10; m++ {
				for n := m; n < 10; n++ {
					for o := n; o < 10; o++ {
						for p := o; p < 10; p++ {
							if k == l || l == m || m == n || n == o || o == p {
								num := k*100000 + l*10000 + m*1000 + n*100 + o*10 + p
								if num > end {
									return count
								}
								if num < start {
									continue
								}
								count++
							}
						}
					}
				}
			}
		}
	}
	return count
}
