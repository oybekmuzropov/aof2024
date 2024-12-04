package main

import (
	"fmt"
	"github.com/oybekmuzropov/aof2024/util"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	strs := util.Read("./day3/input.txt")
	str := ""
	for j := 0; j < len(strs); j++ {
		str += string(strs[j])
	}

	reg := "mul\\(\\d{1,3},\\d{1,3}\\)|(do\\(\\))|(don't\\(\\))"
	r, _ := regexp.Compile(reg)

	ss := r.FindAllStringSubmatch(str, -1)

	res1 := 0
	res2 := 0
	ok := true

	for _, s := range ss {
		s1 := s[0]

		if s1 == "don't()" {
			ok = false
			continue
		} else if s1 == "do()" {
			ok = true
			continue
		}

		s1 = s1[4:]
		s1 = s1[:len(s1)-1]

		sss := strings.Split(s1, ",")

		val1, _ := strconv.Atoi(sss[0])
		val2, _ := strconv.Atoi(sss[1])

		res1 += val1 * val2
		if ok {
			res2 += val1 * val2
		}
	}
	fmt.Println(res1, res2)
}
