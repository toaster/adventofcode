package orbit

import (
	"strings"
)

type tree struct {
	id       string
	orbiters []tree
}

// Count returns the count of orbits.
func Count(input string) int {
	defs := map[string][]string{}
	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			continue
		}
		kv := strings.Split(l, ")")
		defs[kv[0]] = append(defs[kv[0]], kv[1])
	}
	t := buildTree(defs, "COM")
	return countDepths(t, 0)
}

func buildTree(defs map[string][]string, key string) tree {
	t := tree{id: key}
	for _, k := range defs[key] {
		t.orbiters = append(t.orbiters, buildTree(defs, k))
	}
	return t
}

func countDepths(t tree, rd int) int {
	d := rd
	for _, o := range t.orbiters {
		d += countDepths(o, rd+1)
	}
	return d
}
