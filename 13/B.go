package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type point struct {
	x, y int
}

func foldX(points []point, p int) []point {
	for i := range points {
		x, y := points[i].x, points[i].y
		if x < p {
			points[i] = point{p - x - 1, y}
		} else {
			points[i] = point{x - p - 1, y}
		}
	}
	return points
}

func foldY(points []point, p int) []point {
	for i := range points {
		x, y := points[i].x, points[i].y
		if y > p {
			points[i] = point{x, 2*p - y}
		}
	}
	return points
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func print(points []point) {
	maxX, maxY := 0, 0
	for _, p := range points {
		if p.y >= maxY {
			maxY = p.y
		}
		if p.x >= maxX {
			maxX = p.x
		}
	}

	grid := make([][]byte, maxY+1)
	for y := range grid {
		grid[y] = make([]byte, maxX+1)
		for x := range grid[y] {
			grid[y][x] = '.'
		}
	}

	for _, p := range points {
		grid[p.y][p.x] = '#'
	}
	for y := range grid {
		fmt.Println(reverse(string(grid[y])))
	}
}

func main() {
	pointRe := regexp.MustCompile(`(\d+),(\d+)`)
	emptyRe := regexp.MustCompile(`^$`)
	foldRe := regexp.MustCompile(`fold along (x|y)=(\d+)`)

	var points []point

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := string(scanner.Text())
		if emptyRe.MatchString(line) {
			// empty line
		} else if pointRe.MatchString(line) {
			m := pointRe.FindStringSubmatch(line)
			x, _ := strconv.Atoi(m[1])
			y, _ := strconv.Atoi(m[2])
			points = append(points, point{x, y})
		} else if foldRe.MatchString(line) {
			m := foldRe.FindStringSubmatch(line)
			pos, _ := strconv.Atoi(m[2])
			if m[1] == "x" {
				points = foldX(points, pos)
			} else if m[1] == "y" {
				points = foldY(points, pos)
			} else {
				panic("fold")
			}
		} else {
			fmt.Fprintf(os.Stderr, "parse error: %s\n", line)
			os.Exit(2)
		}
	}
	print(points)
}
