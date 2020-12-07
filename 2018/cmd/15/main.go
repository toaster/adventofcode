package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	blocked int = -1
	enemy   int = -2
)

type arena struct {
	m         map[pos]tile
	w, h      int
	deadElves int
}

func (a arena) set(x, y int, t tile) {
	a.m[p(x, y)] = t
}

type pos struct {
	x, y int
}

func p(x, y int) pos {
	return pos{x, y}
}

type tile interface {
	c() byte
}

type unit interface {
	tile
	ap() int
	hp() int
	hit(int) bool
	ignored() bool
	ignore(bool)
	isEnemy(unit) bool
}

type structure byte

func (s structure) c() byte {
	return byte(s)
}

const (
	floor structure = '.'
	wall  structure = '#'
)

type elf struct {
	_ap, _hp int
	i        bool
}

func (e *elf) c() byte {
	return 'E'
}

func (e *elf) ap() int {
	return e._ap
}

func (e *elf) hp() int {
	return e._hp
}

func (e *elf) hit(d int) bool {
	e._hp -= d
	return e._hp <= 0
}

func (e *elf) ignored() bool {
	return e.i
}

func (e *elf) ignore(i bool) {
	e.i = i
}

func (e *elf) isEnemy(u unit) bool {
	_, ok := u.(*goblin)
	return ok
}

type goblin struct {
	_ap, _hp int
	i        bool
}

func (g *goblin) c() byte {
	return 'G'
}

func (g *goblin) ap() int {
	return g._ap
}

func (g *goblin) hp() int {
	return g._hp
}

func (g *goblin) hit(d int) bool {
	g._hp -= d
	return g._hp <= 0
}

func (g *goblin) ignored() bool {
	return g.i
}

func (g *goblin) ignore(i bool) {
	g.i = i
}

func (g *goblin) isEnemy(u unit) bool {
	_, ok := u.(*elf)
	return ok
}

func main() {
	inputFile := os.Args[1]
	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSuffix(string(b), "\n"), "\n")

	h := len(lines)
	w := len(lines[0])
	goblinAP := 3
	elfAP := 3
	for elvesWin := false; !elvesWin; {
		a := arena{map[pos]tile{}, w, h, 0}
		for y, l := range lines {
			for x, b := range []byte(l) {
				a.set(x, y, parseField(b, elfAP, goblinAP))
			}
		}

		// printMap(&a, w, h)
		i := 1
		for done := false; !done; i++ {
			withMap(&a, func(p pos, t tile) {
				if u, ok := t.(unit); ok {
					done = processUnit(&a, u, p)
				}
			}, nil)
			// time.Sleep(1000 * time.Millisecond)
			// fmt.Println("After Round:", i)
			// printMap(&a, w, h)
		}
		hp := 0
		withMap(&a, func(_ pos, t tile) {
			if u, ok := t.(unit); ok {
				hp += u.hp()
			}
		}, nil)
		fmt.Printf("elf AP: %d; Outcome: %d * %d = %d; dead elves: %d\n", elfAP, i-2, hp, hp*(i-2), a.deadElves)
		if a.deadElves == 0 {
			elvesWin = true
		} else {
			elfAP++
		}
	}
}

func parseField(b byte, eAP, gAP int) tile {
	switch b {
	case '#':
		return wall
	case 'E':
		return &elf{_ap: eAP, _hp: 200}
	case 'G':
		return &goblin{_ap: gAP, _hp: 200}
	default:
		return floor
	}
}

func printMap(a *arena, w, h int) {
	units := []unit{}
	withMap(a, func(_ pos, t tile) {
		fmt.Print(string(t.c()))
		if u, ok := t.(unit); ok {
			units = append(units, u)
		}
	}, func() {
		for _, u := range units {
			fmt.Printf(" %c(% 3d)", u.c(), u.hp())
		}
		units = []unit{}
		fmt.Println("")
	})
}

func withMap(a *arena, tf func(pos, tile), lf func()) {
	for y := 0; y < a.h; y++ {
		for x := 0; x < a.w; x++ {
			c := p(x, y)
			tf(c, a.m[c])
		}
		if lf != nil {
			lf()
		}
	}
}

func processUnit(a *arena, u unit, c pos) bool {
	if u.ignored() {
		u.ignore(false)
		return false
	}
	distances := map[pos]int{}
	enemies := []pos{}
	computeDistances(&distances, &enemies, a, []pos{c}, 0, u)
	if len(enemies) == 0 {
		if allEnemiesDead(a, u) {
			return true
		}
		return false
	}
	//fmt.Println(enemies)

	// fmt.Println(string(u.c()))
	// withMap(a, func(p pos, _ tile) {
	// 	if p == c {
	// 		fmt.Print("  X")
	// 	} else {
	// 		fmt.Printf("% 3d", distances[p])
	// 	}
	// }, func() { fmt.Println("") })

	target := searchAttackableTarget(a, u, c)

	// head for target (if necessary)
	if (target == pos{}) {
		// map targets to their nearest adjacent position
		eap := map[pos]pos{}
		for _, ep := range enemies {
			for _, ap := range adjacentPositions(ep) {
				if ap == c {
					eap[ep] = c
				} else if distances[ap] > 0 {
					if (eap[ep] == pos{} || distances[eap[ep]] > distances[ap]) {
						eap[ep] = ap
					}
				}
			}
		}
		//fmt.Println(eap)

		// detect target to heading for
		md := 0
		for _, ap := range eap {
			if md == 0 || distances[ap] < md {
				md = distances[ap]
			}
		}
		tc := map[pos]pos{}
		for ep, ap := range eap {
			if distances[ap] == md {
				tc[ep] = ap
			}
		}
		// fmt.Println("select target to head for from:", tc)
		ht := pos{}
		for y := 0; (y < a.h && ht == pos{}); y++ {
			for x := 0; (x < a.w && ht == pos{}); x++ {
				c := p(x, y)
				if (tc[c] != pos{}) {
					ht = c
				}
			}
		}
		// go
		// fmt.Println("HEADING FOR", ht, tc[ht])
		// detect immediate attack position
		nc := pos{}
		for _, ap := range adjacentPositions(c) {
			if tc[ht] == ap {
				nc = ap
				break
			}
		}
		if (nc == pos{}) {
			d := map[pos]int{}
			computeDistances(&d, nil, a, []pos{tc[ht]}, 0, nil)
			// withMap(a, func(p pos, _ tile) {
			// 	if p == tc[ht] {
			// 		fmt.Print("  X")
			// 	} else {
			// 		fmt.Printf("% 3d", d[p])
			// 	}
			// }, func() { fmt.Println("") })
			for _, ap := range adjacentPositions(c) {
				// fmt.Println("have a look at", ap, a.m[ap] == floor, distances[tc[ht]], d[ap])
				if a.m[ap] == floor && distances[tc[ht]]-1 == d[ap] {
					nc = ap
					break
				}
			}
		}
		// fmt.Println("chose", nc)
		a.m[nc] = a.m[c]
		a.m[c] = floor
		u.ignore(nc.x > c.x || nc.y > c.y)

		if (target == pos{}) {
			target = searchAttackableTarget(a, u, nc)
		}
	}

	// attack (if possible)
	if (target != pos{}) {
		// fmt.Println("ATTACK", target)
		if a.m[target].(unit).hit(u.ap()) {
			// fmt.Println("DIED", target)
			if _, ok := a.m[target].(*elf); ok {
				a.deadElves++
			}
			a.m[target] = floor
		}
	}
	return false
}

func allEnemiesDead(a *arena, u unit) bool {
	noEnemies := true
	withMap(a, func(_ pos, t tile) {
		if ou, ok := t.(unit); ok && u.isEnemy(ou) {
			noEnemies = false
		}
	}, nil)
	return noEnemies
}

func searchAttackableTarget(a *arena, u unit, up pos) pos {
	target := pos{}
	// fmt.Println("select target on one of:", adjacentPositions(up))
	for _, ap := range adjacentPositions(up) {
		if ou, ok := a.m[ap].(unit); ok && u.isEnemy(ou) {
			// fmt.Println("consider", ap)
			if (target == pos{} || ou.hp() < a.m[target].(unit).hp()) {
				target = ap
			}
		}
	}
	return target
}

func computeDistances(distances *map[pos]int, enemies *[]pos, a *arena, cs []pos, d int, u unit) {
	ncs := []pos{}
	for _, c := range cs {
		for _, nc := range adjacentPositions(c) {
			if (*distances)[nc] == 0 {
				if a.m[nc] == floor {
					(*distances)[nc] = d + 1
				} else if ou, ok := a.m[nc].(unit); ok && enemies != nil && u.isEnemy(ou) {
					(*distances)[nc] = enemy
					*enemies = append(*enemies, nc)
				} else {
					(*distances)[nc] = blocked
				}
				if (*distances)[nc] > 0 {
					ncs = append(ncs, nc)
				}
			}
		}
	}
	if len(ncs) > 0 {
		computeDistances(distances, enemies, a, ncs, d+1, u)
	}
}

func adjacentPositions(c pos) []pos {
	return []pos{p(c.x, c.y-1), p(c.x-1, c.y), p(c.x+1, c.y), p(c.x, c.y+1)}
}
