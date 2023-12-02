package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	valves, reachable := parseTunnelScan(io.ReadLines())
	// distances := map[string]map[string]int{}
	// paths := map[string]map[string][]string{}
	// for name := range reachable {
	// distances[name] = map[string]int{name: 1}
	// paths[name] = map[string][]string{}
	// computeDistancesFromNode(distances[name], paths[name], reachable, name, 1, nil)
	// }
	// start := "AA"
	// remainingTime := 30
	fmt.Printf("valves: %#v\n", valves)
	// printAllDistances(distances)
	// node := "AA"
	// printDistances(node, distances[node], paths[node], valves)
	// node = "NQ"
	// printDistances(node, distances[node], paths[node], valves)
	// node = "SX"
	// printDistances(node, distances[node], paths[node], valves)
	// openValves := map[string]bool{}
	// fmt.Println(computeMaxRate(nodes, "AA", 30, openValves))
	paths := []*path{{
		node:          "AA",
		openValves:    map[string]bool{},
		rate:          0,
		remainingTime: 30,
	}}
	for {
		fmt.Println(len(paths))
		var newPaths []*path
		for _, p := range paths {
			if p.remainingTime > 0 {
				if valves[p.node] > 0 && !p.openValves[p.node] {
					newPath := advancePath(p)
					newPath.openValves[p.node] = true
					newPath.rate += valves[p.node]
					newPaths = append(newPaths, newPath)
				}
				for _, node := range reachable[p.node] {
					newPath := advancePath(p)
					newPath.node = node
					newPaths = append(newPaths, newPath)
				}
			}
		}
		if newPaths == nil {
			break
		}
		time.Sleep(1 * time.Second)
		paths = newPaths
	}
	for _, p := range paths {
		fmt.Println(p.amount)
	}
}

func advancePath(p *path) *path {
	newPath := &path{}
	*newPath = *p
	newPath.openValves = map[string]bool{}
	for node, open := range p.openValves {
		newPath.openValves[node] = open
	}
	newPath.remainingTime--
	newPath.amount += newPath.rate
	return newPath
}

func computeDistancesFromNode(distances map[string]int, paths map[string][]string, reachable map[string][]string, cur string, offset int, path []string) {
	for _, next := range reachable[cur] {
		d := offset + 1
		if distances[next] == 0 || distances[next] > d {
			distances[next] = d
			newPath := make([]string, len(path))
			copy(newPath, path)
			newPath = append(newPath, next)
			paths[next] = newPath
			computeDistancesFromNode(distances, paths, reachable, next, d, newPath)
		}
	}
}

func parseTunnelScan(lines []string) (map[string]int, map[string][]string) {
	rates := map[string]int{}
	reachable := map[string][]string{}
	for _, line := range lines {
		name := line[6:8]
		sepIndex := strings.IndexRune(line, ';')
		rates[name] = io.ParseInt(line[23:sepIndex])
		subIndex := sepIndex + 24
		if line[subIndex] == ' ' {
			subIndex++
		}
		reachable[name] = strings.Split(line[subIndex:], ", ")
	}
	return rates, reachable
}

// func computeMaxRate(nodes map[string]*math.Node, startName string, remainingTime int, openValves map[string]bool) (maxRate int, path []string) {
// 	// fmt.Println("compute for", startName, "time remaining", remainingTime)
// 	node := nodes[startName]
// 	path = append(path, startName)
// 	if !openValves[startName] {
// 		if rate := node.Data.(*nodeData).Rate; rate > 0 {
// 			remainingTime--
// 			maxRate += rate * remainingTime
// 			openValves[startName] = true
// 		}
// 	}
// 	if remainingTime > 1 {
// 		remainingTime--
// 		maxSubrate := 0
// 		var selectedSubpath []string
// 		for _, edge := range node.Edges {
// 			name := edge.B.Data.(*nodeData).Name
// 			if nodes[name] == nil {
// 				continue
// 			}
//
// 			rate, subpath := computeMaxRate(nodes, name, remainingTime, openValves)
// 			if rate > maxSubrate {
// 				// fmt.Println("prefer", subpath, "with", rate, "over", selectedSubpath, "with", maxSubrate)
// 				maxSubrate = rate
// 				selectedSubpath = subpath
// 			} else {
// 				// fmt.Println("ignore", subpath, "with", rate, "in favour of", selectedSubpath, "with", maxSubrate)
// 			}
// 		}
// 		maxRate += maxSubrate
// 		path = append(path, selectedSubpath...)
// 	}
// 	return
// }

// func printAllDistances(distances map[string]map[string]int) {
// 	for src, dists := range distances {
// 		printDistances(src, dists)
// 	}
// }

func printDistances(src string, dists map[string]int, paths map[string][]string, valves map[string]int) {
	fmt.Printf("%s:\n", src)
	for tgt, dist := range dists {
		fmt.Printf("  => %s: %d - ", tgt, dist)
		for i, hop := range paths[tgt] {
			if i > 0 {
				fmt.Print(" -> ")
			}
			fmt.Printf("%s (%d)", hop, valves[hop])
		}
		fmt.Println()
	}
}

type path struct {
	amount        int
	node          string
	openValves    map[string]bool
	rate          int
	remainingTime int
}
