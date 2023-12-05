package farm

import (
	math2 "math"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

// ParseAlmanac parses an Almanac.
func ParseAlmanac(lines []string) *Almanac {
	almanac := &Almanac{}
	almanac.seedToSoil, lines = parseAlmanacMap(lines[3:])
	almanac.soilToFertilizer, lines = parseAlmanacMap(lines[1:])
	almanac.fertilizerToWater, lines = parseAlmanacMap(lines[1:])
	almanac.waterToLight, lines = parseAlmanacMap(lines[1:])
	almanac.lightToTemperature, lines = parseAlmanacMap(lines[1:])
	almanac.temperatureToHumidity, lines = parseAlmanacMap(lines[1:])
	almanac.humidityToLocation, lines = parseAlmanacMap(lines[1:])
	return almanac
}

// Almanac describes a farm almanac which contains various mappings.
type Almanac struct {
	seedToSoil            *almanacMap
	soilToFertilizer      *almanacMap
	fertilizerToWater     *almanacMap
	waterToLight          *almanacMap
	lightToTemperature    *almanacMap
	temperatureToHumidity *almanacMap
	humidityToLocation    *almanacMap
}

// NearestLocationForSeeds computes the location ranges for the given seed ranges and returns the nearest (smallest) location.
func (a *Almanac) NearestLocationForSeeds(seeds []*math.Range) int {
	soils := a.seedToSoil.lookup(seeds)
	fertilizers := a.soilToFertilizer.lookup(soils)
	waters := a.fertilizerToWater.lookup(fertilizers)
	lights := a.waterToLight.lookup(waters)
	temperatures := a.lightToTemperature.lookup(lights)
	humidities := a.temperatureToHumidity.lookup(temperatures)
	locations := a.humidityToLocation.lookup(humidities)
	return math.SortRanges(locations)[0].Start
}

func parseAlmanacMap(lines []string) (*almanacMap, []string) {
	var explicitlyMapped []*math.Range
	m := map[*math.Range]*math.Range{}
	for len(lines) > 0 && lines[0] != "" {
		nums := io.ParseInts(lines[0], " ")
		l := nums[2]
		src := &math.Range{Start: nums[1], End: nums[1] + l - 1}
		m[src] = &math.Range{Start: nums[0], End: nums[0] + l - 1}
		explicitlyMapped = append(explicitlyMapped, src)
		lines = lines[1:]
	}
	if len(lines) > 0 {
		lines = lines[1:]
	}
	// Now fill in identity mappings for the uncovered ranges.
	// This simplifies lookups.
	explicitlyMapped = math.SortRanges(explicitlyMapped)
	var mapped []*math.Range
	cur := 0
	for _, r := range explicitlyMapped {
		if cur < r.Start {
			identity := &math.Range{Start: cur, End: r.Start - 1}
			m[identity] = identity
			mapped = append(mapped, identity)
		}
		mapped = append(mapped, r)
		cur = r.End + 1
	}
	if cur < math2.MaxInt {
		identity := &math.Range{Start: cur, End: math2.MaxInt}
		m[identity] = identity
		mapped = append(mapped, identity)
	}
	return &almanacMap{m: m, mapped: mapped}, lines
}

type almanacMap struct {
	m      map[*math.Range]*math.Range
	mapped []*math.Range
}

func (m *almanacMap) lookup(ranges []*math.Range) (result []*math.Range) {
	for _, i := range ranges {
		for _, src := range m.mapped {
			if src.Start > i.End {
				break
			}
			if src.Includes(i) {
				dest := m.m[src]
				offset := dest.Start - src.Start
				result = append(result, &math.Range{Start: i.Start + offset, End: i.End + offset})
			} else if i.Includes(src) {
				result = append(result, m.m[src])
			} else if src.Overlaps(i) {
				dest := m.m[src]
				offset := dest.Start - src.Start
				if src.Start < i.Start {
					result = append(result, &math.Range{Start: i.Start + offset, End: dest.End})
				} else {
					result = append(result, &math.Range{Start: dest.Start, End: i.End + offset})
				}
			}
		}
	}
	return
}
