package main

import (
	"fmt"
	"github.com/oybekmuzropov/aof2024/util"
)

func main() {
	strs := util.Read("./day4/input.txt")

	res := 0
	res2 := 0

	m := len(strs)
	n := len(strs[0])

	for i := 0; i < len(strs); i++ {
		s := strs[i]
		for j := 0; j < len(s); j++ {
			c := s[j]

			if c == 'X' {
				if j-3 >= 0 && s[j-1] == 'M' && s[j-2] == 'A' && s[j-3] == 'S' {
					res++
				}
				if j+3 < len(s) && s[j+1] == 'M' && s[j+2] == 'A' && s[j+3] == 'S' {
					res++
				}
				if i+3 < len(strs) && strs[i+1][j] == 'M' && strs[i+2][j] == 'A' && strs[i+3][j] == 'S' {
					res++
				}
				if i-3 >= 0 && j+3 < len(strs) && strs[i-1][j+1] == 'M' && strs[i-2][j+2] == 'A' && strs[i-3][j+3] == 'S' {
					res++
				}
				if i+3 < len(strs) && j-3 >= 0 && strs[i+1][j-1] == 'M' && strs[i+2][j-2] == 'A' && strs[i+3][j-3] == 'S' {
					res++
				}
				if i-3 >= 0 && strs[i-1][j] == 'M' && strs[i-2][j] == 'A' && strs[i-3][j] == 'S' {
					res++
				}
				if i+3 < len(strs) && j+3 < len(s) && strs[i+1][j+1] == 'M' && strs[i+2][j+2] == 'A' && strs[i+3][j+3] == 'S' {
					res++
				}
				if i-3 >= 0 && j-3 >= 0 && strs[i-1][j-1] == 'M' && strs[i-2][j-2] == 'A' && strs[i-3][j-3] == 'S' {
					res++
				}
			}
		}
	}

	for i := 0; i < len(strs); i++ {
		s := strs[i]

		for j := 0; j < len(s); j++ {
			c := s[j]

			if c == 'A' {
				if i-1 >= 0 && j-1 >= 0 &&
					i+1 < m && j+1 < n {
					tl := strs[i-1][j-1]
					tr := strs[i-1][j+1]
					bl := strs[i+1][j-1]
					br := strs[i+1][j+1]

					b1 := (tl == 'M' && br == 'S') || (tl == 'S' && br == 'M')
					b2 := (tr == 'M' && bl == 'S') || (tr == 'S' && bl == 'M')

					if b1 && b2 {
						res2++
					}
				}
			}
		}
	}
	fmt.Println(res)
	fmt.Println(res2)
}
