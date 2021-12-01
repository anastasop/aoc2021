package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var depths []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		d, err := strconv.Atoi(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "data error: %v: %s", err, line)
			os.Exit(2)
		}
		depths = append(depths, d)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %v", err)
		os.Exit(2)
	}

	incrsCount := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			incrsCount++
		}
	}
	fmt.Println(incrsCount)
}
