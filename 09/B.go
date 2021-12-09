package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type lowPoint struct {
	x, y int
	v    byte
}

func inMatrix(nums [][]byte, x, y int) bool {
	return 0 <= x && x < len(nums[0]) && 0 <= y && y < len(nums)
}

func lowPoints(nums [][]byte) []lowPoint {
	isLowPoint := func(x, y int, points ...int) bool {
		if len(points)%2 == 1 {
			panic("check points")
		}

		checks, passed := 0, 0
		for i := 0; i < len(points); i += 2 {
			if inMatrix(nums, points[i], points[i+1]) {
				checks++
				if nums[y][x] < nums[points[i+1]][points[i]] {
					passed++
				}
			}
		}
		return checks == passed
	}

	var s []lowPoint
	for y := range nums {
		for x := range nums[y] {
			if isLowPoint(x, y, x+1, y, x-1, y, x, y-1, x, y+1) {
				s = append(s, lowPoint{x, y, nums[y][x]})
			}
		}
	}
	return s
}

func flood(nums [][]byte, q []int) int {
	c := 0

	for i := 0; len(q) > 0; i++ {
		x := q[0]
		y := q[1]
		q = q[2:]
		if inMatrix(nums, x, y) && nums[y][x] < 9 {
			c++
			nums[y][x] += 100
			q = append(q, x+1, y)
			q = append(q, x-1, y)
			q = append(q, x, y+1)
			q = append(q, x, y-1)
		}
	}

	return c
}

func main() {
	var nums [][]byte

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		b := make([]byte, len(scanner.Bytes()))
		copy(b, scanner.Bytes())
		for i := range b {
			b[i] -= '0'
		}
		nums = append(nums, b)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("IO error: ", err)
		os.Exit(2)
	}

	var sizes []int
	for _, p := range lowPoints(nums) {
		s := flood(nums, []int{p.x, p.y})
		sizes = append(sizes, s)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	prod := sizes[0] * sizes[1] * sizes[2]
	fmt.Println(prod)
}
