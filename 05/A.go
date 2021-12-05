package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type point struct {
	x int
	y int
}

var grid = make(map[point]int)

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func mark(x0, y0, x1, y1 int) {
	if x0 == x1 {
		for p0, p1 := minmax(y0, y1); p0 <= p1; p0++ {
			grid[point{x0, p0}]++
		}
	} else if y0 == y1 {
		for p0, p1 := minmax(x0, x1); p0 <= p1; p0++ {
			grid[point{p0, y0}]++
		}
	}
}

func main() {
	num := regexp.MustCompile(`[0-9]+`)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		nums := num.FindAllString(line, -1)
		mark(atoi(nums[0]), atoi(nums[1]), atoi(nums[2]), atoi(nums[3]))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %v", err)
		os.Exit(2)
	}

	c := 0
	for _, v := range grid {
		if v >= 2 {
			c++
		}
	}
	fmt.Println(c)
}
