package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type point struct {
	x, y int
	dist int
}

var (
	grid [][]byte
	N    int
	D    int
	T    int = 5
)

func main() {
	inp, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %v\n", err)
		os.Exit(2)
	}

	lines := strings.Split(string(inp), "\n")
	lines = lines[:len(lines)-1]
	grid = make([][]byte, len(lines))
	for y, line := range lines {
		grid[y] = make([]byte, len(line))
		for x, c := range line {
			grid[y][x] = byte(c - '0')
		}
	}
	D = len(grid)
	N = T * D

	search()
}

func inGrid(x, y int) bool {
	return 0 <= y && y < N && 0 <= x && x < N
}

func val(x, y int) int {
	n := int(grid[y%D][x%D]) + x/D + y/D
	if n > 9 {
		n -= 9
	}
	return n
}

func search() {
	var q []*point
	var e []*point

	q = append(q, &point{0, 0, 0})
	for len(q) > 0 {
		var top *point

		q, top = pointMinDist(q)
		e = add(e, top.x, top.y, top.dist)

		if top.x == N-1 && top.y == N-1 {
			break
		}

		check := func(x, y int) {
			if inGrid(x, y) {
				d := top.dist + val(x, y)
				t0, t1 := find(q, x, y), find(e, x, y)
				if t0 == nil && t1 == nil {
					q = add(q, x, y, d)
				} else if t0 != nil && d < t0.dist {
					t0.dist = d
				}
			}
		}

		x, y := top.x, top.y
		check(x, y-1)
		check(x, y+1)
		check(x-1, y)
		check(x+1, y)
	}

	fmt.Println(e[len(e)-1].dist)
}

func pointMinDist(q []*point) ([]*point, *point) {
	i := 0
	p := q[0]
	for k, t := range q {
		if t.dist < p.dist {
			i = k
			p = t
		}
	}
	q[i] = q[len(q)-1]
	q = q[:len(q)-1]
	return q, p
}

func find(q []*point, x, y int) *point {
	for _, t := range q {
		if t.x == x && t.y == y {
			return t
		}
	}
	return nil
}

func add(q []*point, x, y, dist int) []*point {
	return append(q, &point{x, y, dist})
}
