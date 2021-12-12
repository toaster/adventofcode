package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	m := map[string]*cave{}
	for _, line := range io.ReadLines() {
		c := strings.Split(line, "-")
		s := c[0]
		e := c[1]
		m[s] = addNeighbour(m[s], s, e)
		m[e] = addNeighbour(m[e], e, s)
	}
	startName := "start"
	paths := followPath(m, &path{
		caves:   []string{startName},
		visited: map[string]bool{startName: true},
	})
	for _, p := range paths {
		fmt.Println(strings.Join(p.caves, ","))
	}
	fmt.Println(len(paths))
}

func followPath(m map[string]*cave, p *path) []*path {
	var paths []*path
	cur := m[p.caves[len(p.caves)-1]]
	for _, neighbour := range cur.neighbours {
		doubleVisit := false
		if strings.ToLower(neighbour) == neighbour && p.visited[neighbour] {
			if p.doubleVisitedSmallCave != "" || neighbour == "start" {
				continue
			} else {
				doubleVisit = true
			}
		}

		newPath := &path{
			caves:                  newCaves(p.caves, neighbour),
			visited:                newVisited(p.visited, neighbour),
			doubleVisitedSmallCave: p.doubleVisitedSmallCave,
		}

		if neighbour == "end" {
			paths = append(paths, newPath)
		} else {
			if doubleVisit {
				newPath.doubleVisitedSmallCave = neighbour
			}
			paths = append(paths, followPath(m, newPath)...)
		}
	}
	return paths
}

func newCaves(caves []string, neighbour string) []string {
	c := make([]string, len(caves))
	copy(c, caves)
	c = append(c, neighbour)
	return c
}

func newVisited(visited map[string]bool, neighbour string) map[string]bool {
	n := map[string]bool{}
	for k, v := range visited {
		n[k] = v
	}
	n[neighbour] = true
	return n
}

func addNeighbour(c *cave, name, neighbour string) *cave {
	if c == nil {
		c = &cave{name: name}
	}
	c.neighbours = append(c.neighbours, neighbour)
	return c
}

type cave struct {
	name       string
	neighbours []string
}

type path struct {
	caves                  []string
	visited                map[string]bool
	doubleVisitedSmallCave string
}
