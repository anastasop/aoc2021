package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var total int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		solveEntry(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %v", err)
		os.Exit(2)
	}
	fmt.Println(total)
}

func solveEntry(entry string) {
	fields := strings.Fields(entry)
	i := 0
	for i = range fields {
		if fields[i] == "|" {
			break
		}
	}
	// digits := fields[0:i]
	display := fields[i+1:]

	for _, d := range display {
		if len(d) == 2 || len(d) == 3 || len(d) == 4 || len(d) == 7 {
			total++
		}
	}
}
