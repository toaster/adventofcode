package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/util"
)

const expansionFactor = 2

func main() {
	sum := 0
	for _, line := range io.ReadLines() {
		conditionRecordsAndDamagedCounts := strings.Split(line, " ")
		damagedCounts := io.ParseInts(conditionRecordsAndDamagedCounts[1], ",")
		sum += countPossibleArrangements(conditionRecordsAndDamagedCounts[0], damagedCounts)
		//  computePossibleFoldedArrangements(conditionRecordsAndDamagedCounts[0], damagedCounts, )
	}
	fmt.Println(sum)
}

func countPossibleArrangements(conditionRecords string, damagedCounts []int) int {
	possibleArrangements := map[string]bool{}
	groupsCandidates := []*damagedOrUnclearGroups{parseDamagedOrUnclearGroups(conditionRecords)}
	computePossibleArrangementsStartingWithGroup(groupsCandidates, damagedCounts, 0, 0, possibleArrangements)
	fmt.Println(conditionRecords, "+", damagedCounts, "=>", len(possibleArrangements))
	return len(possibleArrangements)
}

func countPossibleFoldedArrangements(foldedConditionRecords string, foldedDamagedCounts []int) int {
	possibleArrangements := map[string]bool{}
	conditionRecords := ""
	var damagedCounts []int
	for i := 0; i < 5; i++ {
		if i > 0 {
			conditionRecords += "?"
		}
		conditionRecords += foldedConditionRecords
		damagedCounts = append(damagedCounts, foldedDamagedCounts...)
	}
	fmt.Println("records:", foldedConditionRecords, "=>", conditionRecords, "expected", damagedCounts)

	groupsCandidates := []*damagedOrUnclearGroups{parseDamagedOrUnclearGroups(conditionRecords)}
	computePossibleArrangementsStartingWithGroup(groupsCandidates, damagedCounts, 0, 0, possibleArrangements)
	fmt.Println(conditionRecords, "=>", len(possibleArrangements))
	return len(possibleArrangements)
}

func computePossibleArrangementsStartingWithGroup(groupsCandidates []*damagedOrUnclearGroups, damagedCounts []int, startGroup, level int, possibleArrangements map[string]bool) {
	prefix := ""
	debug := false
	if debug {
		prefix := fmt.Sprintf(fmt.Sprintf("%% %ds", level*2), "")
		fmt.Print(prefix)
		fmt.Println("START AT", startGroup)
		for _, candidate := range groupsCandidates {
			fmt.Print(prefix)
			fmt.Printf("  - %s\n", candidate.display())
		}
	}
	for _, groups := range groupsCandidates {
		nextDamagedCountsIndex := startGroup
		if debug {
			fmt.Print(prefix)
			fmt.Printf("  => %s\n", groups.display())
		}
		if startGroup == len(damagedCounts) && groups.groupCount == startGroup {
			possibleArrangements[groups.string()] = true
			continue
		}
		for i := startGroup; i < groups.groupCount; i++ {
			if nextDamagedCountsIndex == len(damagedCounts) {
				if groups.damaged[i] {
					// no damaged group expected
					break
				} else {
					newCandidates := []*damagedOrUnclearGroups{groups.newGroupsByIgnoring(i)}
					computePossibleArrangementsStartingWithGroup(newCandidates, damagedCounts, i, level+1, possibleArrangements)
					break
				}
			}

			expectedDamagedGroupSize := damagedCounts[nextDamagedCountsIndex]
			if groups.damaged[i] {
				groupSize := groups.groupSizes[i]
				// fmt.Println("look at damaged", i, "of size", groupSize, "expected", expectedDamagedGroupSize)
				if groupSize == expectedDamagedGroupSize {
					if groups.groupCount > i+1 && groups.groupPositions[i+1] == groups.groupPositions[i]+groupSize {
						if groups.groupSizes[i+1] == 1 {
							// cut group of single '?' which has to be '.'
							groups.groupPositions = append(groups.groupPositions[:i+1], groups.groupPositions[i+2:]...)
							groups.groupSizes = append(groups.groupSizes[:i+1], groups.groupSizes[i+2:]...)
							groups.groupCount--
							for j := i + 1; j < groups.groupCount; j++ {
								groups.damaged[j] = groups.damaged[j+1]
							}
						} else {
							// shorten following group of '?' by one
							groups.groupPositions[i+1]++
							groups.groupSizes[i+1]--
						}
						groups.s = ""
						if debug {
							fmt.Printf("optimized %#v\n", groups.string())
						}
					}
					nextDamagedCountsIndex++
				} else if groupSize < expectedDamagedGroupSize {
					newCandidates := groups.computeCandidatesForDamagedGroup(i, expectedDamagedGroupSize)
					computePossibleArrangementsStartingWithGroup(newCandidates, damagedCounts, i, level+1, possibleArrangements)
					// fmt.Println("  damaged group size too low")
					break
				} else {
					// fmt.Println("  damaged group size too high")
					break
				}
				// fmt.Println("  group", i, "matched")
				if i == groups.groupCount-1 && nextDamagedCountsIndex == len(damagedCounts) {
					// fmt.Println("considered matching", groups.string())
					possibleArrangements[groups.string()] = true
				}
			} else {
				newCandidates := groups.computeCandidatesForUnsureGroup(i, expectedDamagedGroupSize)
				computePossibleArrangementsStartingWithGroup(newCandidates, damagedCounts, i, level+1, possibleArrangements)
				break
			}
		}
	}
}

func parseDamagedOrUnclearGroups(conditionRecords string) *damagedOrUnclearGroups {
	g := &damagedOrUnclearGroups{damaged: map[int]bool{}}
	last := '.'
	for i, r := range conditionRecords {
		switch r {
		case '#', '?':
			if last != r {
				g.groupPositions = append(g.groupPositions, i)
				g.groupSizes = append(g.groupSizes, 1)
				if r == '#' {
					g.damaged[g.groupCount] = true
				}
				g.groupCount++
			} else {
				g.groupSizes[g.groupCount-1]++
			}
		}
		last = r
	}
	return g
}

type damagedOrUnclearGroups struct {
	damaged        map[int]bool
	groupCount     int
	groupPositions []int
	groupSizes     []int
	s              string
}

func (g *damagedOrUnclearGroups) computeCandidatesForDamagedGroup(groupIndex int, expectedDamagedGroupSize int) []*damagedOrUnclearGroups {
	if groupIndex+1 == g.groupCount {
		return nil
	}

	groupSize := g.groupSizes[groupIndex]
	groupPos := g.groupPositions[groupIndex]
	if g.groupPositions[groupIndex+1] > groupPos+groupSize {
		return nil
	}

	s := g.string()
	if len(s) < groupPos+expectedDamagedGroupSize {
		return nil
	}

	for i := 0; i < expectedDamagedGroupSize; i++ {
		if s[groupPos+i] == '.' {
			return nil
		}
	}
	mustBeGoodPos := groupPos + expectedDamagedGroupSize
	if len(s) > mustBeGoodPos && s[mustBeGoodPos] == '#' {
		return nil
	}

	n := &damagedOrUnclearGroups{
		damaged: map[int]bool{},
	}

	for i := 0; i < groupIndex; i++ {
		n.groupCount++
		n.groupSizes = append(n.groupSizes, g.groupSizes[i])
		n.groupPositions = append(n.groupPositions, g.groupPositions[i])
		n.damaged[i] = g.damaged[i]
	}
	n.groupPositions = append(n.groupPositions, g.groupPositions[groupIndex])
	n.groupSizes = append(n.groupSizes, expectedDamagedGroupSize)
	n.damaged[groupIndex] = true
	n.groupCount++
	if len(s) > mustBeGoodPos {
		i := groupIndex
		for ; g.groupPositions[i]+g.groupSizes[i] < mustBeGoodPos && i < g.groupCount; i++ {
		}
		afterCurrentGroupPos := g.groupPositions[i] + g.groupSizes[i]
		if afterCurrentGroupPos > mustBeGoodPos+1 {
			n.groupPositions = append(n.groupPositions, mustBeGoodPos+1)
			n.groupSizes = append(n.groupSizes, afterCurrentGroupPos-mustBeGoodPos-1)
			n.groupCount++
		}
		i++
		for ; i < g.groupCount; i++ {
			n.groupSizes = append(n.groupSizes, g.groupSizes[i])
			n.groupPositions = append(n.groupPositions, g.groupPositions[i])
			n.damaged[n.groupCount] = g.damaged[i]
			n.groupCount++
		}
	}
	return []*damagedOrUnclearGroups{n}
}

func (g *damagedOrUnclearGroups) computeCandidatesForUnsureGroup(groupIndex, minSize int) []*damagedOrUnclearGroups {
	var candidates []*damagedOrUnclearGroups
	groupSize := g.groupSizes[groupIndex]
	// fmt.Println("unsure", groupIndex, "-", minSize)
	candidates = append(candidates, g.newGroupsByIgnoring(groupIndex))
	candidates = append(candidates, g.newGroupsByMarkingDamaged(groupIndex, minSize)...)
	if groupSize > minSize {
		for i := 0; i < groupSize-minSize; i++ {
			candidate := g.newGroupsBySplitting(groupIndex, minSize+i)
			if i == 0 {
				candidate.damaged[groupIndex] = true
			}
			candidates = append(candidates, candidate)
		}
	}
	// for _, candidate := range candidates {
	// 	fmt.Printf("    => %s %#v\n", candidate.string(), candidate)
	// }
	return candidates
}

func (g *damagedOrUnclearGroups) copy() *damagedOrUnclearGroups {
	return &damagedOrUnclearGroups{
		damaged:        util.CopyMap(g.damaged),
		groupCount:     g.groupCount,
		groupPositions: util.CopySlice(g.groupPositions),
		groupSizes:     util.CopySlice(g.groupSizes),
	}
}

func (g *damagedOrUnclearGroups) mergeIntoPrevious(groupIndex int) {
	if groupIndex < 1 {
		panic("invalid merge")
	}

	g.groupSizes[groupIndex-1] += g.groupSizes[groupIndex]
	g.groupSizes = append(g.groupSizes[:groupIndex], g.groupSizes[groupIndex+1:]...)
	g.groupPositions = append(g.groupPositions[:groupIndex], g.groupPositions[groupIndex+1:]...)
	g.groupCount--
	for i := groupIndex; i <= g.groupCount; i++ {
		g.damaged[i] = g.damaged[i+1]
	}
}

func (g *damagedOrUnclearGroups) newGroupsByIgnoring(groupIndex int) *damagedOrUnclearGroups {
	n := g.copy()
	n.groupCount--
	n.groupPositions = append(n.groupPositions[:groupIndex], n.groupPositions[groupIndex+1:]...)
	n.groupSizes = append(n.groupSizes[:groupIndex], n.groupSizes[groupIndex+1:]...)
	for i := groupIndex; i < n.groupCount; i++ {
		n.damaged[i] = n.damaged[i+1]
	}
	n.damaged[n.groupCount] = false
	return n
}

func (g *damagedOrUnclearGroups) newGroupsByMarkingDamaged(groupIndex, damagedSize int) []*damagedOrUnclearGroups {
	groupSize := g.groupSizes[groupIndex]
	damagedFollows := 0
	afterGroupIndex := g.groupPositions[groupIndex] + g.groupSizes[groupIndex]
	if g.groupCount > groupIndex+1 && g.groupPositions[groupIndex+1] == afterGroupIndex {
		damagedFollows = g.groupSizes[groupIndex+1]
	}
	var newGroups []*damagedOrUnclearGroups
	if damagedFollows > 0 && damagedFollows <= damagedSize {
		maxDamaged := min(damagedSize-damagedFollows, groupSize)
		for i := 1; i <= maxDamaged; i++ {
			n := g.newGroupsByIgnoring(groupIndex)
			n.groupPositions[groupIndex] -= i
			n.groupSizes[groupIndex] += i
			newGroups = append(newGroups, n)
		}
	}
	maxOffset := groupSize - damagedSize
	if damagedFollows > 0 {
		maxOffset--
	}
	for i := 0; i <= maxOffset; i++ {
		n := g.copy()
		n.damaged[groupIndex] = true
		n.groupPositions[groupIndex] += i
		n.groupSizes[groupIndex] = damagedSize
		newGroups = append(newGroups, n)
	}
	return newGroups
}

func (g *damagedOrUnclearGroups) newGroupsBySplitting(groupIndex, firstPartSize int) *damagedOrUnclearGroups {
	n := &damagedOrUnclearGroups{
		damaged:    util.CopyMap(g.damaged),
		groupCount: g.groupCount + 1,
	}
	for i := 0; i < g.groupCount; i++ {
		pos := g.groupPositions[i]
		size := g.groupSizes[i]
		if i != groupIndex {
			n.groupPositions = append(n.groupPositions, pos)
			n.groupSizes = append(n.groupSizes, size)
		} else if i == groupIndex {
			n.groupPositions = append(n.groupPositions, pos)
			n.groupSizes = append(n.groupSizes, firstPartSize)
			n.groupPositions = append(n.groupPositions, pos+firstPartSize+1)
			n.groupSizes = append(n.groupSizes, size-firstPartSize-1)
		}
	}
	for i := g.groupCount; i > groupIndex; i-- {
		n.damaged[i] = n.damaged[i-1]
	}
	return n
}

func (g *damagedOrUnclearGroups) display() string {
	s := strings.Builder{}
	cur := 0
	for groupIndex := 0; groupIndex < g.groupCount; groupIndex++ {
		groupPos := g.groupPositions[groupIndex]
		groupSize := g.groupSizes[groupIndex]
		if cur < groupPos && cur > 0 {
			s.WriteRune('|')
		}
		for ; cur < groupPos; cur++ {
			s.WriteRune('.')
		}
		s.WriteRune('|')
		r := '?'
		if g.damaged[groupIndex] {
			r = '#'
		}
		for ; cur < groupPos+groupSize; cur++ {
			s.WriteRune(r)
		}
	}
	return s.String()
}

func (g *damagedOrUnclearGroups) string() string {
	if g.s == "" {
		s := strings.Builder{}
		cur := 0
		for groupIndex := 0; groupIndex < g.groupCount; groupIndex++ {
			groupPos := g.groupPositions[groupIndex]
			groupSize := g.groupSizes[groupIndex]
			for ; cur < groupPos; cur++ {
				s.WriteRune('.')
			}
			r := '?'
			if g.damaged[groupIndex] {
				r = '#'
			}
			for ; cur < groupPos+groupSize; cur++ {
				s.WriteRune(r)
			}
		}
		g.s = s.String()
	}
	return g.s
}
