package slice_diff

import "golang.org/x/exp/constraints"

// SliceDiff takes 2 slices of the same type and returns the difference
// The difference returned is what listB does not match within listA
func SliceDiff[T constraints.Ordered](listA []T, listB []T) []T {
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
