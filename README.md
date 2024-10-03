[![Go Reference](https://pkg.go.dev/badge/github.com/matt9mg/go-slice-diff.svg)](https://pkg.go.dev/github.com/matt9mg/go-slice-diff)

# Slice Diff
Returns the difference between two slices

### Installation
```go
go get github.com/mtdevs28080617/go-slice-diff
```

### Examples
```go
listA := []string{"A", "B", "C"}
listB := []string{"A", "B", "D"}

diff := slice_diff.SliceDiff[string](listA, listB)

log.Println(diff) // []string{"D"}

listA := []int{1, 2, 3}
listB := []int{1, 2, 4}

diff := slice_diff.SliceDiff[int](listA, listB)

log.Println(diff) // []int{4}
```