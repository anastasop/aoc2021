package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
)

type spec struct {
	vx, vy     int
	xMin, xMax int
	yMin, yMax int
}

func (s spec) probe() (bool, int) {
	altY := math.MinInt
	vx, vy := s.vx, s.vy
	x, y := 0, 0
	for {
		x += vx
		y += vy

		if y > altY {
			altY = y
		}

		vy--
		if vx > 0 {
			vx--
		} else if vx < 0 {
			vx++
		}

		if s.xMin <= x && x <= s.xMax && s.yMin <= y && y <= s.yMax {
			return true, altY
		}

		if y < s.yMin && vy < 0 {
			return false, 0
		}
	}
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %v\n")
		os.Exit(2)
	}

	r := regexp.MustCompile(`-?\d+`)
	strs := r.FindAllString(string(data), -1)
	nums := make([]int, len(strs))
	for i, s := range strs {
		nums[i], _ = strconv.Atoi(s)
	}

	minY := math.MinInt
	for vy := -1000; vy <= 1000; vy++ {
		for vx := 0; vx <= 1000; vx++ {
			s := spec{vx, vy, nums[0], nums[1], nums[2], nums[3]}
			if hit, y := s.probe(); hit {
				if y > minY {
					minY = y
				}
			}
		}
	}
	fmt.Println(minY)
}
