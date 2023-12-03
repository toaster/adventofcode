package gondolaengine

import "github.com/toaster/advent_of_code/internal/math"

// ParseSchematic parses a Schematic from the given input.
func ParseSchematic(lines []string) *Schematic {
	s := &Schematic{
		numRefs: map[math.Point2D]*number{},
		symbols: map[math.Point2D]*byte{},
	}
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			c := line[x]

			{
				num := 0
				numLen := 0
				for {
					v, isNum := math.DetectDigit(c)
					if !isNum {
						break
					}

					x++
					if x < len(line) {
						c = line[x]
					} else {
						c = '.'
					}
					numLen++
					num *= 10
					num += v
				}
				if numLen > 0 {
					n := &number{
						len:   numLen,
						loc:   math.Point2D{X: x - numLen, Y: y},
						value: num,
					}
					s.nums = append(s.nums, n)
					for i := 0; i < numLen; i++ {
						s.numRefs[n.loc.AddXY(i, 0)] = n
					}
				}
			}

			if c == '.' {
				continue
			}

			s.symbols[math.Point2D{X: x, Y: y}] = &c
		}
	}
	return s
}

// Schematic describes an engine schematic.
type Schematic struct {
	nums    []*number
	numRefs map[math.Point2D]*number
	symbols map[math.Point2D]*byte
}

// SumOfGearRatios computes the sum of all gear rations of the Schematic.
func (s *Schematic) SumOfGearRatios() (sum int) {
	for p, sym := range s.symbols {
		if *sym != '*' {
			continue
		}

		nums := map[*number]bool{}
		for _, a := range p.Adjacents() {
			if s.numRefs[a] != nil {
				nums[s.numRefs[a]] = true
			}
		}
		if len(nums) == 2 {
			ratio := 1
			for n := range nums {
				ratio *= n.value
			}
			sum += ratio
		}
	}
	return sum
}

// SumOfPartNumbers computes the sum of all part numbers of the Schematic.
func (s *Schematic) SumOfPartNumbers() (sum int) {
	for _, n := range s.nums {
		isPartNum := false
		if s.symbols[n.loc.SubtractXY(1, 0)] != nil || s.symbols[n.loc.AddXY(n.len, 0)] != nil {
			isPartNum = true
		}
		for y := -1; !isPartNum && y < 2; y += 2 {
			for x := -1; !isPartNum && x < n.len+1; x++ {
				if s.symbols[n.loc.AddXY(x, y)] != nil {
					isPartNum = true
				}
			}
		}
		if isPartNum {
			sum += n.value
		}
	}
	return sum
}

type number struct {
	len   int
	loc   math.Point2D
	value int
}
