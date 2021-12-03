package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %v", err)
		os.Exit(2)
	}
	sort.Sort(sort.StringSlice(lines))

	o2s := search(lines, true, '1', 0)
	if len(o2s) != 1 {
		fmt.Fprintf(os.Stderr, "O2 found %d values\n", len(o2s))
		os.Exit(2)
	}
	o2, _ := strconv.ParseInt(o2s[0], 2, 64)

	co2s := search(lines, false, '0', 0)
	if len(co2s) != 1 {
		fmt.Fprintf(os.Stderr, "CO2 found %d values\n", len(co2s))
		os.Exit(2)
	}
	co2, _ := strconv.ParseInt(co2s[0], 2, 64)

	fmt.Printf("O2 %d CO2 %d P %d\n", o2, co2, o2*co2)
}

func search(nums []string, maj bool, tie byte, pos int) []string {
	if len(nums) <= 1 || pos >= len(nums[0]) {
		return nums
	}

	i := 0
	for i = 0; i < len(nums) && nums[i][pos] == '0'; i++ {
	}

	var next []string
	if i > len(nums)-i {
		if maj {
			next = nums[0:i]
		} else {
			next = nums[i:]
		}
	} else if nums[0][pos] == tie {
		next = nums[0:i]
	} else {
		next = nums[i:]
	}

	return search(next, maj, tie, pos+1)
}
