package orbit

import (
	"strings"
)

type tree struct {
	id       string
	orbiters []tree
}

// Count returns the count of transitions needed for YOU to get to SAN.
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
	var py, ps []string
	walkTree(t, []string{}, func(n tree, p []string) {
		switch n.id {
		case "YOU":
			py = p
		case "SAN":
			ps = p
		}
	})
	for i := range py {
		if py[i] != ps[i] {
			py = py[i:]
			ps = ps[i:]
			break
		}
	}
	return len(py) + len(ps)
}

func buildTree(defs map[string][]string, key string) tree {
	t := tree{id: key}
	for _, k := range defs[key] {
		t.orbiters = append(t.orbiters, buildTree(defs, k))
	}
	return t
}

func walkTree(t tree, p []string, cb func(tree, []string)) {
	cb(t, p)
	p = append(p, t.id)
	for _, o := range t.orbiters {
		walkTree(o, p, cb)
	}
}

func countDepths(t tree, rd int) int {
	d := rd
	for _, o := range t.orbiters {
		d += countDepths(o, rd+1)
	}
	return d
}
