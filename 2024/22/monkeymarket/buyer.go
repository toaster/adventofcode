package monkeymarket

// NewBuyer creates a new Buyer with the given initial secret.
func NewBuyer(initialSecret int) *Buyer {
	return &Buyer{lastSecret: initialSecret, lastPrice: initialSecret % 10}
}

// Buyer is a buyer on the MonkeyMarket.
type Buyer struct {
	lastSecret int
	lastPrice  int
}

// ComputeSecret computes the secret after the given amount of iterations
func (m *Buyer) ComputeSecret(iterations int) int {
	for i := 0; i < iterations; i++ {
		m.computeNext()
	}
	return m.lastSecret
}

func (m *Buyer) computeNext() int {
	s := m.lastSecret << 6
	s ^= m.lastSecret
	s &= pruneValue
	m.lastSecret = s
	s = m.lastSecret >> 5
	s ^= m.lastSecret
	s &= pruneValue
	m.lastSecret = s
	s = m.lastSecret << 11
	s ^= m.lastSecret
	s &= pruneValue
	p := s % 10
	change := p - m.lastPrice
	m.lastSecret = s
	m.lastPrice = p
	return change
}

const pruneValue = 2<<23 - 1
