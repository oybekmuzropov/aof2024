package main

import (
	"bufio"
	"fmt"
	"github.com/oybekmuzropov/aof2024/util"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func isGood(arr []int) bool {
	if util.IsSorted(arr) {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1]-arr[i] > 3 || arr[i] == arr[i+1] {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func read() []string {
	f, err := os.Open("./day2/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	strs := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		strs = append(strs, scanner.Text())
	}

	return strs
}

func main() {
	strs := read()
	count1 := 0
	count2 := 0

	for j := 0; j < len(strs); j++ {
		ss := strings.Split(strs[j], " ")

		arr := []int{}

		for i := 0; i < len(ss); i++ {
			val, _ := strconv.Atoi(ss[i])

			arr = append(arr, val)
		}

		if isGood(arr) {
			count1++
		}
		for i := 0; i < len(arr); i++ {
			na := slices.Clone(arr[:i])
			na = append(na, arr[i+1:]...)
			if isGood(na) {
				count2++
				break
			}
		}
	}
	fmt.Println(count1, count2)
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}
