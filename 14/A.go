package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	template := ""
	rules := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "->") {
			fields := strings.Fields(line)
			rules[fields[0]] = fields[2]
		} else {
			template = line
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %v\n", err)
		os.Exit(2)
	}

	s := template
	for i := 0; i < 10; i++ {
		s = step(s, rules)
	}
	fmt.Println(sig(s))
}

func step(s string, rules map[string]string) string {
	t := make([]byte, 0, len(s)*2)
	t = append(t, s[0])
	for i := 1; i < len(s); i++ {
		rule := rules[s[i-1:i]+s[i:i+1]]
		t = append(t, rule[0], s[i])
	}
	return string(t)
}

func sig(s string) int {
	var freqs [26]int

	for _, c := range s {
		freqs[c-'A']++
	}

	min, max := 1000000, 0
	for _, f := range freqs {
		if f > max {
			max = f
		}
		if 0 < f && f < min {
			min = f
		}
	}
	return max - min
}
