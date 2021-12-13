// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var graph = make(map[string][]string)
var path []string
var paths = 0

func canVisit(v string) bool {
	if v == "start" {
		return len(path) == 0
	}

	freqs := make(map[string]int)
	freqs[v]++
	for _, s := range path {
		if strings.Index("abcdefghijklmnopqrstuvwxzy", string(s[0])) >= 0 {
			freqs[s]++
		}
	}
	twoCount := 0
	maxVal := 0
	for _, v := range freqs {
		if v >= 2 {
			twoCount++
		}
		if v > maxVal {
			maxVal = v
		}
	}

	return twoCount <= 1 && maxVal <= 2
}

func dfs(v string) {
	if v == "end" {
		paths++
		return
	}
	path = append(path, v)
	for _, w := range graph[v] {
		check := w == "start" || strings.Index("abcdefghijklmnopqrstuvwxzy", string(w[0])) >= 0
		if !check || canVisit(w) {
			dfs(w)
		}
	}
	path = path[0 : len(path)-1]
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %v\n", err)
		os.Exit(2)
	}
	for _, e := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		s := strings.Split(e, "-")
		v, w := s[0], s[1]
		graph[v] = append(graph[v], w)
		graph[w] = append(graph[w], v)
	}
	dfs("start")
	fmt.Println(paths)
}
