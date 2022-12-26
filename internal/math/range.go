package math

// Range represents a range.
type Range struct {
	Start int
	End   int
}

// AdjacentTo returns whether the Range is adjacent to another Range (without overlapping).
func (r *Range) AdjacentTo(other *Range) bool {
	return r.Start == other.End+1 || other.Start == r.End+1
}

// Covers returns whether the Range covers the given value.
func (r *Range) Covers(i int) bool {
	return i >= r.Start && i <= r.End
}

// Includes returns whether the Range completely covers (includes) another Range.
func (r *Range) Includes(other *Range) bool {
	return r.Start <= other.Start && r.End >= other.End
}

// Merge merges the Range with another one if possible and returns the new Range as well as whether the operation was performed.
func (r *Range) Merge(other *Range) (*Range, bool) {
	if !r.Overlaps(other) && !r.AdjacentTo(other) {
		return nil, false
	}

	return &Range{Start: MinInt(r.Start, other.Start), End: MaxInt(r.End, other.End)}, true
}

// Overlaps returns whether the Range covers a part of (overlaps with) another Range.
func (r *Range) Overlaps(other *Range) bool {
	return r.Includes(other) ||
		(r.Start >= other.Start && r.Start <= other.End) ||
		(r.End >= other.Start && r.End <= other.End)
}
