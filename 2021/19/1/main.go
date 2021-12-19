package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	var scanners []*scanner
	var cur *scanner
	for _, line := range io.ReadLines() {
		if strings.HasPrefix(line, "--- scanner") {
			cur = &scanner{name: line}
			scanners = append(scanners, cur)
			continue
		}

		if line == "" {
			continue
		}

		values := io.ParseInts(line, ",")
		p := math.Point3D{X: values[0], Y: values[1], Z: values[2]}
		cur.beaconPositions = append(cur.beaconPositions, p)
	}

	for _, s := range scanners {
		for _, beaconPos := range s.beaconPositions {
			b := &beacon{}
			s.beacons = append(s.beacons, b)
			for _, pos := range s.beaconPositions {
				b.neighbourPositions = append(b.neighbourPositions, pos.Subtract(beaconPos))
			}
		}
	}

	// p1 := math.Point3D{X: 1, Y: 2, Z: 3}
	// p2 := math.Point3D{X: 2, Y: -3, Z: -1}
	// for i := 0; i < math.OrientationCount3D; i++ {
	// 	fmt.Printf("p1(%02d): %s\tp2(%02d): %s\n", i, math.TransformOrientation(p1, i).String(), i, math.TransformOrientation(p2, i).String())
	// }

	// s := scanners[0]
	// for o := 0; o < math.OrientationCount3D; o++ {
	// 	fmt.Println("Orientation:", o)
	// 	for _, pos := range s.beaconPositions[o] {
	// 		fmt.Println(pos)
	// 	}
	// }

	// commonBeaconPositions(scanners[0], scanners[1])

	final := []*scanner{scanners[0]}
	beaconPositions := map[math.Point3D]bool{}
	for _, position := range final[0].beaconPositions {
		beaconPositions[position] = true
	}
	scanners = scanners[1:]
	for len(scanners) > 0 {
	scan:
		for i, a := range scanners {
			for _, b := range final {
				commonB, commonA, orientation := commonBeaconIndices(b, a)
				if len(commonA) > 11 {
					fmt.Print("add: ", a.name, " … ")

					adjusted := &scanner{
						beaconPositions: transformedPositions(a.beaconPositions, orientation),
						name:            a.name + " adjusted",
						orientation:     orientation,
					}
					adjusted.position = b.position.Add(b.beaconPositions[commonB[0]].Subtract(adjusted.beaconPositions[commonA[0]]))

					fmt.Println("pos: ", adjusted.position)
					// fmt.Println("common from", b.name)
					// for _, x := range commonB {
					// 	fmt.Println(b.beaconPositions[x])
					// }
					// fmt.Println("common from", a.name)
					// for _, x := range commonA {
					// 	fmt.Println(a.beaconPositions[x])
					// }
					// fmt.Println("common adjusted")
					// for _, x := range commonA {
					// 	fmt.Println(adjusted.beaconPositions[x].Add(adjusted.position))
					// }

					for _, bcn := range a.beacons {
						adjusted.beacons = append(adjusted.beacons, &beacon{transformedPositions(bcn.neighbourPositions, orientation)})
					}

					final = append(final, adjusted)
					scanners = append(scanners[:i], scanners[i+1:]...)
					for _, position := range adjusted.beaconPositions {
						beaconPositions[position.Add(adjusted.position)] = true
					}
					fmt.Println("done")
					// os.Exit(1)
					break scan
				}
			}
		}
	}
	fmt.Println(len(beaconPositions))
	var sortable math.Sortable3DPoints
	for pos := range beaconPositions {
		sortable = append(sortable, pos)
	}
	sort.Sort(sortable)
	for _, p := range sortable {
		fmt.Println(p)
	}

	// beaconCount := 0
	// var done []*scanner
	// for i, a := range scanners {
	// 	fmt.Println("add", len(a.beacons), "beacons of scanner", i)
	// 	beaconCount += len(a.beacons)
	// 	for j, sb := range done {
	// 		cc := commonBeaconCount(a, sb)
	// 		fmt.Println("scanner", i, "and", j, "overlap by", cc, "beacons … don’t count these")
	// 		beaconCount -= cc
	// 	}
	// 	done = append(done, a)
	// }
	// fmt.Println(beaconCount)
}

func commonBeaconIndices(a *scanner, b *scanner) (commonA, commonB []int, orientation int) {
	for _, ba := range a.beacons {
		for _, bb := range b.beacons {
			for orientation = 0; orientation < math.OrientationCount3D; orientation++ {
				transformed := transformedPositions(bb.neighbourPositions, orientation)
				commonA, commonB = commonPointIndices(ba.neighbourPositions, transformed)
				if len(commonA) > 11 {
					// for _, index := range commonIndices {
					// 	common = append(common, a.beaconPositions[index])
					// }
					// fmt.Println("common:", len(common), "at:", orientation)
					// for _, pos := range common {
					// 	fmt.Println(pos)
					// }
					return
				}
			}
		}
	}
	return nil, nil, 0
}

func commonCount(a, b []math.Point3D) (count int) {
	for _, pa := range a {
		for _, pb := range b {
			if pa == pb {
				count++
			}
		}
	}
	return
}

// func commonBeaconCount(a, b *scanner) (count int) {
// 	for _, ba := range a.beacons {
// 		for _, bb := range b.beacons {
// 			c := maxCommonNeighbourCount(ba, bb)
// 			// fmt.Println("max neighbour count:", c)
// 			if c > 11 && c > count {
// 				count = c
// 			}
// 		}
// 	}
// 	return
// }

func commonPointIndices(a, b []math.Point3D) (commonA, commonB []int) {
	for i, pa := range a {
		for j, pb := range b {
			if pa == pb {
				commonA = append(commonA, i)
				commonB = append(commonB, j)
			}
		}
	}
	return
}

// func maxCommonNeighbourCount(a, b *beacon) (max int) {
// 	for o := 0; o < math.OrientationCount3D; o++ {
// 		count := commonCount(a.neighbourPositions[0], b.neighbourPositions[o])
// 		if max < count {
// 			max = count
// 		}
// 	}
// 	return
// }

func transformedPositions(positions []math.Point3D, o int) (transformed []math.Point3D) {
	if o == 0 {
		return positions
	}

	for _, pos := range positions {
		transformed = append(transformed, math.TransformOrientation(pos, o))
	}
	return
}

type scanner struct {
	beaconPositions []math.Point3D
	beacons         []*beacon
	name            string
	orientation     int
	position        math.Point3D
}

type beacon struct {
	neighbourPositions []math.Point3D
}
