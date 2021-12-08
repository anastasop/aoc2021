package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// brute force search. Couldn't think of anything else

var bits = []uint{64, 32, 16, 8, 4, 2, 1}

var correct = []uint{
	0: 0b1110111,
	1: 0b0010010,
	2: 0b1011101,
	3: 0b1011011,
	4: 0b0111010,
	5: 0b1101011,
	6: 0b1101111,
	7: 0b1010010,
	8: 0b1111111,
	9: 0b1111011,
}

func valueOf(n uint) int {
	for i, v := range correct {
		if v == n {
			return i
		}
	}
	panic("number not found")
}

func d2i(s string, perm []uint) uint {
	var i uint
	for _, d := range s {
		i |= perm[d-rune('a')]
	}
	return i
}

// from https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
func permutations(arr []uint) [][]uint {
	var helper func([]uint, int)
	res := [][]uint{}

	helper = func(arr []uint, n int) {
		if n == 1 {
			tmp := make([]uint, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func solveEntry(digits, display []string, perms [][]uint) int {
	var mem [10]uint
	var use []uint
	c := 0

	for i := range perms {
		nums := mem[:0]
		for _, d := range digits {
			nums = append(nums, d2i(d, perms[i]))
		}

		m := make(map[uint]int)
		for _, n := range nums {
			m[n]++
		}

		found := true
		for _, v := range correct {
			if m[v] != 1 {
				found = false
			}
		}

		if found {
			use = perms[i]
			c++
		}
	}

	if c > 1 {
		fmt.Fprintln(os.Stderr, "no unique entry")
		os.Exit(2)
	}

	n := 0
	for _, d := range display {
		n = 10*n + valueOf(d2i(d, use))
	}

	return n
}

func main() {
	perms := permutations(bits)

	n := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		i := 0
		for i = range fields {
			if fields[i] == "|" {
				break
			}
		}
		digits := fields[0:i]
		display := fields[i+1:]

		n += solveEntry(digits, display, perms)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %v", err)
		os.Exit(2)
	}
	fmt.Println(n)
}
