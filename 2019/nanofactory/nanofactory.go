package nanofactory

import (
	"strconv"
	"strings"

	"github.com/toaster/advent_of_code/internal/math"
)

// Refinery can convert ORE into FUEL.
type Refinery struct {
	reactions map[string]reactDef
	inv       map[string]int
}

type reactDef struct {
	amount      int
	ingredients map[string]int
}

// Parse creates a Refinery with the reactions given by input.
func Parse(input string) *Refinery {
	r := &Refinery{reactions: map[string]reactDef{}}
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		ingredientsAndProduct := strings.Split(line, "=>")
		amountAndProduct := strings.Split(strings.TrimSpace(ingredientsAndProduct[1]), " ")

		a, _ := strconv.Atoi(amountAndProduct[0])
		i := map[string]int{}
		r.reactions[amountAndProduct[1]] = reactDef{amount: a, ingredients: i}

		for _, ai := range strings.Split(strings.TrimSpace(ingredientsAndProduct[0]), ",") {
			amountAndIngredient := strings.Split(strings.TrimSpace(ai), " ")
			i[amountAndIngredient[1]], _ = strconv.Atoi(amountAndIngredient[0])
		}
	}
	return r
}

// ComputeOreRequirementForOneFuel computes how much ore is needed to create exactly one FUEL.
func (r *Refinery) ComputeOreRequirementForOneFuel() int {
	r.inv = map[string]int{}
	return r.computeOre("FUEL", 1)
}

// ComputeFuelProducableWithOreAmount computes how much FUEL can be produced with the given amount of ORE.
func (r *Refinery) ComputeFuelProducableWithOreAmount(oreAmount int) int {
	o1 := r.ComputeOreRequirementForOneFuel()
	start := oreAmount / o1
	ok := true
	fa := start
	// search bad
	for ; ok; fa *= 2 {
		r.inv = map[string]int{}
		oa := r.computeOre("FUEL", fa)
		ok = oa < oreAmount
	}
	step := -fa / 2
	fa += step
	// bisect to max
	for ; ; fa += step {
		r.inv = map[string]int{}
		oa := r.computeOre("FUEL", fa)
		ok = oa < oreAmount
		if math.AbsInt(step) < 2 {
			if ok {
				break
			} else {
				step = -1
			}
		} else {
			if ok {
				step = math.AbsInt(step) / 2
			} else {
				step = -math.AbsInt(step) / 2
			}
		}
	}
	return fa
}

func (r *Refinery) computeOre(requirement string, requiredAmount int) int {
	requiredAmount -= r.inv[requirement]
	if requiredAmount <= 0 {
		r.inv[requirement] = -requiredAmount
		return 0
	}
	r.inv[requirement] = 0
	ore := 0
	react := r.reactions[requirement]
	reactCount := requiredAmount / react.amount
	if requiredAmount%react.amount > 0 {
		reactCount++
	}
	for i, a := range react.ingredients {
		ia := a * reactCount
		if i == "ORE" {
			ore += ia
			continue
		}
		ore += r.computeOre(i, ia)
	}
	r.inv[requirement] += reactCount*react.amount - requiredAmount
	return ore
}
