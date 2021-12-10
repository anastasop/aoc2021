package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var scores []int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if s := score(scanner.Text()); s >= 0 {
			scores = append(scores, s)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %v\n", err)
		os.Exit(2)
	}
	sort.Sort(sort.IntSlice(scores))
	fmt.Println(scores[len(scores)/2])
}

var scores = map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
var match = map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}
var rmatch = map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}
var rscores = map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}

func score(line string) int {
	s := make([]rune, 0, len(line))
	for _, c := range line {
		switch c {
		case '(', '[', '{', '<':
			s = append(s, c)
		case ')', ']', '}', '>':
			if top, exp := s[len(s)-1], match[c]; exp == top {
				s = s[:len(s)-1]
			} else {
				return -1 // corrupted line
			}
		default:
			panic("!")
		}
	}

	compl := ""
	sc := 0
	for i := len(s) - 1; i >= 0; i-- {
		m := rmatch[s[i]]
		compl += string(m)
		sc = sc*5 + rscores[m]
	}

	return sc
}
