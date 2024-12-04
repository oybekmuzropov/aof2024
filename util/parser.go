package util

import (
	"strconv"
	"strings"
)

func ParseToInts(str string) []int {
	ss := strings.Fields(str)

	var arr []int

	for i := 0; i < len(ss); i++ {
		val, _ := strconv.Atoi(ss[i])

		arr = append(arr, val)
	}

	return arr
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
