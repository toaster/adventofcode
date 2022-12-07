package math

// Range represents a range.
type Range struct {
	Start int
	End   int
}

// Includes returns whether the Range completely covers (includes) another Range.
func (r *Range) Includes(other *Range) bool {
	return r.Start <= other.Start && r.End >= other.End
}

// Overlaps returns whether the Range covers a part of (overlaps with) another Range.
func (r *Range) Overlaps(other *Range) bool {
	return r.Includes(other) ||
		(r.Start >= other.Start && r.Start <= other.End) ||
		(r.End >= other.Start && r.End <= other.End)
}
