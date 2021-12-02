package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	x, y := 0, 0
	aim := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		toks := strings.Fields(line)

		n, err := strconv.Atoi(toks[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "data error: %v: %s", err, line)
			os.Exit(2)
		}

		switch toks[0][0] {
		case 'u':
			aim -= n
		case 'd':
			aim += n
		case 'f':
			x += n
			y += n * aim
		default:
			fmt.Fprintf(os.Stderr, "data error: %v: %s", err, line)
			os.Exit(2)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %v", err)
		os.Exit(2)
	}

	fmt.Printf("x(%d) y(%d) %d\n", x, y, x*y)
}
