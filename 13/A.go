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

func uniqCount(points []point) int {
	m := make(map[point]byte)
	for _, p := range points {
		m[p] = 1
	}
	return len(m)
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
				fmt.Printf("fold along %s=%s %d\n", m[1], m[2], uniqCount(points))
			} else if m[1] == "y" {
				points = foldY(points, pos)
				fmt.Printf("fold along %s=%s %d\n", m[1], m[2], uniqCount(points))
			} else {
				panic("fold")
			}
		} else {
			fmt.Fprintf(os.Stderr, "parse error: %s\n", line)
			os.Exit(2)
		}
	}
}
