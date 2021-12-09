package main

import (
	"bufio"
	"fmt"
	"os"
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

	risk := 0
	for _, p := range lowPoints(nums) {
		risk += int(p.v) + 1
	}
	fmt.Println(risk)

}
