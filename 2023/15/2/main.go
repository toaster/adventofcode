package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	initializationSteps := strings.Split(strings.TrimSpace(io.ReadAll()), ",")
	boxes := map[int]*lensBox{}
	for i := 0; i < 256; i++ {
		boxes[i] = &lensBox{number: i}
	}
	for _, step := range initializationSteps {
		if strings.HasSuffix(step, "-") {
			label := step[:len(step)-1]
			box := boxes[computeHash(label)]
			box.removeLensWithLabel(label)
		} else {
			components := strings.Split(step, "=")
			l := &lens{
				label:       components[0],
				focalLength: io.ParseInt(components[1]),
			}
			box := boxes[computeHash(l.label)]
			box.addLens(l)
		}
	}
	sum := 0
	for _, box := range boxes {
		sum += box.focusingPower()
	}
	fmt.Println(sum)
}

func computeHash(s string) int {
	hash := 0
	for _, c := range []byte(s) {
		hash += int(c)
		hash *= 17
		hash &= 0xFF
	}
	return hash
}

type lensBox struct {
	lenses []*lens
	number int
}

func (b *lensBox) addLens(newLens *lens) {
	i := b.indexOfLensWithLabel(newLens.label)
	if i >= 0 {
		b.lenses[i] = newLens
	} else {
		b.lenses = append(b.lenses, newLens)
	}
}

func (b *lensBox) focusingPower() int {
	p := 0
	for i, l := range b.lenses {
		p += (1 + b.number) * (1 + i) * l.focalLength
	}
	return p
}

func (b *lensBox) indexOfLensWithLabel(label string) int {
	return slices.IndexFunc(b.lenses, func(l *lens) bool { return l.label == label })
}

func (b *lensBox) removeLensWithLabel(label string) {
	i := b.indexOfLensWithLabel(label)
	if i >= 0 {
		b.lenses = slices.Delete(b.lenses, i, i+1)
	}
}

func (b *lensBox) string() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("Box %d: ", b.number))
	for i, l := range b.lenses {
		if i > 0 {
			s.WriteRune(' ')
		}
		s.WriteString(fmt.Sprintf("[%s %d]", l.label, l.focalLength))
	}
	return s.String()
}

type lens struct {
	label       string
	focalLength int
}
