package monkeymarket

// New creates a new Market
func New(initialSecrets []int) Market {
	buyers := make([]*Buyer, len(initialSecrets))
	for i, secret := range initialSecrets {
		buyers[i] = NewBuyer(secret)
	}
	return Market{buyers: buyers}
}

// Market describes a monkey market (https://adventofcode.com/2024/day/22).
type Market struct {
	buyers []*Buyer
}

// ComputeMaxAmountOfBananas computes the maximum amount of bananas that can be bought using the
// translating monkey at the Market within the given amount of iterations.
func (m Market) ComputeMaxAmountOfBananas(iterations int) int {
	buyerSequences := map[*Buyer]map[[4]int]int{}
	sequences := map[[4]int]bool{}
	for _, buyer := range m.buyers {
		changes := make([]int, 0, iterations)
		buyerSequences[buyer] = map[[4]int]int{}
		for i := 0; i < iterations; i++ {
			change := buyer.computeNext()
			changes = append(changes, change)
			if i > 3 {
				sequence := [4]int{changes[i-3], changes[i-2], changes[i-1], changes[i]}
				if _, ok := buyerSequences[buyer][sequence]; !ok {
					buyerSequences[buyer][sequence] = buyer.lastPrice
					sequences[sequence] = true
				}
			}
		}
	}
	maxAmount := 0
	for sequence := range sequences {
		amount := 0
		for _, buyer := range m.buyers {
			amount += buyerSequences[buyer][sequence]
		}
		if maxAmount < amount {
			maxAmount = amount
		}
	}
	return maxAmount
}
