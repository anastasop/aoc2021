package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type node struct {
	name      string
	next      [2]string
	currStep  int
	nextStep  int
	currCount int
	nextCount int
}

func main() {
	template := ""
	graph := make(map[string]*node)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "->") {
			fields := strings.Fields(line)
			graph[fields[0]] = &node{
				name:      fields[0],
				next:      [2]string{fields[0][0:1] + fields[2], fields[2] + fields[0][1:]},
				currStep:  -1,
				nextStep:  -1,
				currCount: 0,
				nextCount: 0,
			}
		} else {
			template = line
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %v\n", err)
		os.Exit(2)
	}

	for i := 1; i < len(template); i++ {
		pair := template[i-1:i] + template[i:i+1]
		graph[pair].currStep = 0
		graph[pair].currCount++
	}

	stepsToRun := 40
	var step int
	for step = 0; step < stepsToRun; step++ {
		for _, v := range graph {
			if v.currStep == step {
				for _, z := range v.next {
					w := graph[z]
					if w.nextStep == step+1 {
						w.nextCount += v.currCount
					} else {
						w.nextCount = v.currCount
						w.nextStep = step + 1
					}
				}
			}
		}

		len := 1
		for _, v := range graph {
			v.currStep = v.nextStep
			if v.nextStep == step+1 {
				v.currCount = v.nextCount
				len += v.currCount
			}
		}
	}

	var freqs [26]uint64
	for _, v := range graph {
		if v.currStep == step {
			freqs[v.name[0]-'A'] += uint64(v.currCount)
			freqs[v.name[1]-'A'] += uint64(v.currCount)
		}
	}
	// Each string overlaps so we have counted it twice.
	// There are 2 odd numbers because of the first and last
	// letter of the sequence
	min, max := uint64(math.MaxUint64), uint64(0)
	for _, f := range freqs {
		if f%2 == 1 {
			f++
		}
		f /= 2
		if f > max {
			max = f
		}
		if 0 < f && f < min {
			min = f
		}
	}
	fmt.Println(max - min)
}
