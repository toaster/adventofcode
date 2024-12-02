package report

import (
	"slices"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

// Report describes a report of the Red-Nosed Reindeer nuclear fusion/fission plant.
type Report []int

// IsSafe returns whether the report is safe or not without the Problem Dampener applied.
func (r Report) IsSafe() bool {
	direction := 0
	for i := 1; i < len(r); i++ {
		diff := r[i] - r[i-1]
		absDiff := math.AbsInt(diff)
		if absDiff < 1 || absDiff > 3 {
			return false
		}

		diffDirection := diff / absDiff
		if i == 1 {
			direction = diffDirection
		} else if diffDirection != direction {
			return false
		}
	}
	return true
}

// IsSafeWhenWithProblemDampenerApplied returns whether the report is safe or not with the Problem Dampener applied.
func (r Report) IsSafeWhenWithProblemDampenerApplied() bool {
	if r.IsSafe() {
		return true
	}

	for i := 0; i < len(r); i++ {
		variant := slices.Clone(r)
		variant = append(variant[:i], variant[i+1:]...)
		if variant.IsSafe() {
			return true
		}
	}
	return false
}

// Parse parses a line of input into a Report.
func Parse(line string) Report {
	return io.ParseInts(line, " ")
}
