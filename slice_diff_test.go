package slice_diff_test

import (
	slice_diff "github.com/matt9mg/go-slice-diff"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceDiff_String(t *testing.T) {
	_assert := assert.New(t)

	args1 := []string{"test1", "test"}
	args2 := []string{"test1", "Test"}
	diff := slice_diff.SliceDiff[string](args1, args2)

	_assert.EqualValues([]string{"Test"}, diff)
}

func TestSliceDiff_Int(t *testing.T) {
	_assert := assert.New(t)

	args1 := []int{1, 2, 3, 4}
	args2 := []int{0, 1, 2, 3, 4, 5}
	diff := slice_diff.SliceDiff[int](args1, args2)

	_assert.EqualValues([]int{0, 5}, diff)
}

func TestSliceDiff_Float64(t *testing.T) {
	_assert := assert.New(t)

	args1 := []float64{1.1, 2, 3, 4}
	args2 := []float64{1.0, 2, 3, 4.34}
	diff := slice_diff.SliceDiff[float64](args1, args2)

	_assert.EqualValues([]float64{1.0, 4.34}, diff)
}
