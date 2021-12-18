package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	var numbers []snailfishNumber
	for _, line := range io.ReadLines() {
		numbers = append(numbers, parseSnailfishNumber(line))
	}
	max := 0
	for _, number := range numbers {
		fmt.Println(">>>>>>", number)
		for _, other := range numbers {
			if other == number {
				fmt.Println("skip")
				continue
			}
			sum := snailfishAdd(number, other)
			snailfishReduce(sum)
			mag := sum.Magnitude()
			fmt.Println(sum, "=>", sum.Magnitude())
			if mag > max {
				max = mag
			}
		}
	}
	fmt.Println(max)
}

func parseSnailfishNumber(line string) (n snailfishNumber) {
	n, _ = parseSnailfishSubNumber(line, nil)
	return
}

func parseSnailfishSubNumber(line string, parent snailfishNumber) (snailfishNumber, string) {
	rest := line[1:]
	if line[0] == '[' {
		p := &snPair{}
		p.left, rest = parseSnailfishSubNumber(rest, p)
		p.right, rest = parseSnailfishSubNumber(rest[1:], p)
		p.parent = parent
		return p, rest[1:]
	}

	start := line
	var i int
	for ; rest[0] != ',' && rest[0] != ']'; i++ {
		rest = rest[1:]
	}
	num, err := strconv.Atoi(start[:i+1])
	io.ReportError("failed to parse snailfish number", err)
	return &snInt{parent: parent, value: num}, rest
}

func snailfishAdd(a, b snailfishNumber) snailfishNumber {
	sum := &snPair{}
	sum.left = a.Clone(sum)
	sum.right = b.Clone(sum)
	return sum
}

func snailfishReduce(num snailfishNumber) {
	// fmt.Println("to reduce:", num)
	n := num
	for tryExplode(n) || trySplit(n) {
		// fmt.Println("num:", n)
	}
	// fmt.Println("reduced:", num)
}

func tryAction(n snailfishNumber, action func(n snailfishNumber, level int) bool) bool {
	level := 0
	for {
		if n.IsPair() {
			// fmt.Println("dive left:", n, "at:", level)
			n = n.Left()
			level++
			continue
		}

		if action(n, level) {
			return true
		}

		if n.Parent() != nil && n.Parent().Left() == n {
			// fmt.Println("turn right:", n.Parent(), "at:", level-1)
			n = n.Parent().Right()
			continue
		}

		for n.Parent() != nil && n.Parent().Right() == n {
			level--
			n = n.Parent()
			// fmt.Println("went back:", n, "at:", level)
		}
		if n.Parent() != nil {
			// fmt.Println("continue right:", n.Parent(), "at:", level-1)
			n = n.Parent().Right()
			continue
		}
		break
	}
	return false
}

func tryExplode(number snailfishNumber) bool {
	// fmt.Println("try explode", number)
	return tryAction(number, func(n snailfishNumber, level int) bool {
		parent := n.Parent()
		if level > 4 && !parent.Left().IsPair() && !parent.Right().IsPair() {
			parent.Explode()
			// fmt.Println("exploded", parent, "=>", number)
			return true
		}
		return false
	})
}

func trySplit(number snailfishNumber) bool {
	// fmt.Println("try split", number)
	return tryAction(number, func(n snailfishNumber, _ int) bool {
		if n.Value() > 9 {
			n.Split()
			// fmt.Println("split", n, "=>", number)
			return true
		}
		return false
	})
}

type snailfishNumber interface {
	Clone(parent snailfishNumber) snailfishNumber
	Explode()
	IsPair() bool
	Left() snailfishNumber
	Magnitude() int
	Parent() snailfishNumber
	Right() snailfishNumber
	SetParent(number snailfishNumber)
	Split()
	String() string
	Value() int
}

type snInt struct {
	parent snailfishNumber
	value  int
}

var _ snailfishNumber = (*snInt)(nil)

func (i *snInt) Clone(parent snailfishNumber) snailfishNumber {
	return &snInt{
		parent: parent,
		value:  i.value,
	}
}

func (i *snInt) Explode() {
	return
}

func (i *snInt) IsPair() bool {
	return false
}

func (i *snInt) Left() snailfishNumber {
	return nil
}

func (i *snInt) Magnitude() int {
	return i.value
}

func (i *snInt) Parent() snailfishNumber {
	return i.parent
}

func (i *snInt) Right() snailfishNumber {
	return nil
}

func (i *snInt) SetParent(parent snailfishNumber) {
	i.parent = parent
}

func (i *snInt) Split() {
	split := &snPair{parent: i.parent}
	split.left = &snInt{
		parent: split,
		value:  int(math.Floor(float64(i.value) / 2)),
	}
	split.right = &snInt{
		parent: split,
		value:  int(math.Ceil(float64(i.value) / 2)),
	}
	if i.parent.Left() == i {
		i.parent.(*snPair).left = split
	} else {
		i.parent.(*snPair).right = split
	}
}

func (i *snInt) String() string {
	return strconv.Itoa(i.value)
}

func (i *snInt) Value() int {
	return i.value
}

type snPair struct {
	left   snailfishNumber
	parent snailfishNumber
	right  snailfishNumber
}

var _ snailfishNumber = (*snPair)(nil)

func (p *snPair) Clone(parent snailfishNumber) snailfishNumber {
	clone := &snPair{
		parent: parent,
	}
	clone.left = p.left.Clone(clone)
	clone.right = p.right.Clone(clone)
	return clone
}

func (p *snPair) Explode() {
	// fmt.Println("explode", p)
	leftNum := snailfishNumber(p)
	for leftNum.Parent() != nil && leftNum.Parent().Left() == leftNum {
		leftNum = leftNum.Parent()
	}
	leftNum = leftNum.Parent()
	if leftNum != nil {
		leftNum = leftNum.Left()
	}
	if leftNum != nil {
		for leftNum.IsPair() {
			leftNum = leftNum.Right()
		}
	}
	rightNum := snailfishNumber(p)
	for rightNum.Parent() != nil && rightNum.Parent().Right() == rightNum {
		rightNum = rightNum.Parent()
	}
	rightNum = rightNum.Parent()
	if rightNum != nil {
		rightNum = rightNum.Right()
	}
	if rightNum != nil {
		for rightNum.IsPair() {
			rightNum = rightNum.Left()
		}
	}
	// fmt.Println("left:", leftNum, "right:", rightNum)
	if leftNum != nil {
		leftNum.(*snInt).value += p.left.Value()
	}
	if rightNum != nil {
		rightNum.(*snInt).value += p.right.Value()
	}
	if p.parent.Left() == p {
		p.parent.(*snPair).left = &snInt{parent: p.parent, value: 0}
	} else {
		p.parent.(*snPair).right = &snInt{parent: p.parent, value: 0}
	}
}

func (p *snPair) IsPair() bool {
	return true
}

func (p *snPair) Left() snailfishNumber {
	return p.left
}

func (p *snPair) Magnitude() int {
	return 3*p.left.Magnitude() + 2*p.right.Magnitude()
}

func (p *snPair) Parent() snailfishNumber {
	return p.parent
}

func (p *snPair) Right() snailfishNumber {
	return p.right
}

func (p *snPair) SetParent(parent snailfishNumber) {
	p.parent = parent
}

func (p *snPair) Split() {
	return
}

func (p *snPair) String() string {
	return fmt.Sprintf("[%v,%v]", p.left, p.right)
}

func (p *snPair) Value() int {
	return 0
}
