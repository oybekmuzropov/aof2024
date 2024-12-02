package util

import (
	"slices"
)

func IsSorted(arr []int) bool {
	if slices.IsSorted(arr) {
		return true
	}
	slices.Reverse(arr)
	return slices.IsSorted(arr)
}
