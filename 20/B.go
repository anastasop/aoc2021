package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	B = 10
	L = 50
	N = 4000
	D = 2000
)

type pixel struct {
	x int
	y int
}

var arrays [2][N][N]byte
var algorithm string
var curr, next int

func getI(x, y int) int {
	return int(arrays[curr][y][x])
}

func set(x, y int, c byte) {
	arrays[next][y][x] = c
}

func enhancePixel(x, y int) {
	i := 0 +
		256*getI(x-1, y-1) + 128*getI(x, y-1) + 64*getI(x+1, y-1) +
		32*getI(x-1, y) + 16*getI(x, y) + 8*getI(x+1, y) +
		4*getI(x-1, y+1) + 2*getI(x, y+1) + 1*getI(x+1, y+1)

	var b byte
	if algorithm[i] == '#' {
		b = 1
	}

	set(x, y, b)
}

func enhanceGridBrute() {
	for x := 0; x < N; x++ {
		for y := 0; y < N; y++ {
			arrays[next][y][x] = 0
		}
	}

	for x := B; x < N-B; x++ {
		for y := B; y < N-B; y++ {
			enhancePixel(x, y)
		}
	}

	curr = (curr + 1) % 2
	next = (next + 1) % 2
}

func numLit() int {
	i := 0
	for x := L; x < N-L; x++ {
		for y := L; y < N-L; y++ {
			i += int(arrays[curr][y][x])
		}
	}

	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	algorithm = scanner.Text()
	scanner.Scan()
	// ignore black line
	y := 0
	curr = 1
	next = 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, c := range line {
			if c == '#' {
				set(x+D, y+D, 1)
			} else {
				set(x+D, y+D, 0)
			}
		}
		y++
	}
	curr = 0
	next = 1
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %v\n", err)
		os.Exit(2)
	}

	for i := 0; i < 50; i++ {
		enhanceGridBrute()
	}
	fmt.Println(numLit())
}
