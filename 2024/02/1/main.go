package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2024/02/report"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	var reports []report.Report
	for _, line := range io.ReadLines() {
		reports = append(reports, report.Parse(line))
	}

	count := 0
	for _, r := range reports {
		if r.IsSafe() {
			count++
		}
	}
	fmt.Println(count)
}
