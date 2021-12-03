package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	type count struct {
		ones  int
		zeros int
	}
	var counts []count

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if counts == nil {
			counts = make([]count, len(line))
		}
		for i, b := range line {
			if b == '0' {
				counts[i].zeros++
			} else if b == '1' {
				counts[i].ones++
			} else {
				panic("digit")
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %v", err)
		os.Exit(2)
	}

	g, e := 0, 0
	for _, c := range counts {
		g *= 2
		e *= 2
		if c.ones > c.zeros {
			g++
		} else {
			e++
		}
	}

	fmt.Printf("g = %d e = %d g*e = %d\n", g, e, g*e)
}
