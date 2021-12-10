package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	score := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		score += corruptionScore(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %v\n", err)
		os.Exit(2)
	}
	fmt.Println(score)
}

var scores = map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
var match = map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}

func corruptionScore(line string) int {
	s := make([]rune, 0, len(line))
	for _, c := range line {
		switch c {
		case '(', '[', '{', '<':
			s = append(s, c)
		case ')', ']', '}', '>':
			if top, exp := s[len(s)-1], match[c]; exp == top {
				s = s[:len(s)-1]
			} else {
				return scores[c]
			}
		default:
			panic("!")
		}
	}

	return 0
}
