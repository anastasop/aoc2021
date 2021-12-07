package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func fuel(a, b int) int {
	n := a - b
	if n < 0 {
		n = -n
	}
	return n
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	toks := strings.Split(strings.TrimSpace(string(data)), ",")
	crabs := make([]int, len(toks))
	for i, t := range toks {
		crabs[i], _ = strconv.Atoi(t)
	}

	min, max := 1<<31, -1
	for _, p := range crabs {
		if p < min {
			min = p
		}
		if p > max {
			max = p
		}
	}

	minPos, minFuel := 0, 1<<31
	for p := min; p <= max; p++ {
		sum := 0
		for _, c := range crabs {
			sum += fuel(c, p)
		}
		if sum < minFuel {
			minFuel = sum
			minPos = p
		}
	}

	fmt.Printf("pos %d fuel %d\n", minPos, minFuel)
}
