package slice_diff

// Diff interface describes the types that SliceDiff can check against
type Diff interface {
	string | int | int8 | int16 | int32 | int64 | float32 | float64
}

// SliceDiff[T Diff] takes 2 slices of the same type and returns the difference
// The difference returned it what is not in listA from listB
func SliceDiff[T Diff](listA []T, listB []T) []T {
	ma := make(map[T]struct{}, len(listA))
	var diffs []T

	for _, ka := range listA {
		ma[ka] = struct{}{}
	}

	for _, kb := range listB {
		if _, ok := ma[kb]; !ok {
			diffs = append(diffs, kb)
		}
	}

	return diffs
}
