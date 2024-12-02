package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func read() []string {
	f, err := os.Open("./day1/input.txt")
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

	arr1 := []int{}
	arr2 := []int{}

	m := map[int]int{}

	for _, str := range strs {
		ss := strings.Fields(str)

		val1, _ := strconv.Atoi(ss[0])
		val2, _ := strconv.Atoi(ss[1])

		m[val2]++

		arr1 = append(arr1, val1)
		arr2 = append(arr2, val2)
	}

	sort.Ints(arr1)
	sort.Ints(arr2)
	s := 0
	for i := 0; i < len(arr1); i++ {
		s += abs(arr1[i] - arr2[i])
	}

	s1 := 0
	for i := 0; i < len(arr1); i++ {
		s1 += m[arr1[i]] * arr1[i]
	}
	fmt.Println(s)
	fmt.Println(s1)
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}
