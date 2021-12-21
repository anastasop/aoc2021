package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
)

const (
	D = 100
	S = 10
	W = 1000
)

type Dice struct {
	curr  int
	rolls int
}

func (d *Dice) roll() int {
	d.curr++
	if d.curr > D {
		d.curr = 1
	}
	d.rolls++

	return d.curr
}

type Player struct {
	pos   int
	score int
}

func (p *Player) move(d *Dice) bool {
	v1 := d.roll()
	v2 := d.roll()
	v3 := d.roll()

	np := p.pos + v1 + v2 + v3
	for np > S {
		np -= S
	}
	p.pos = np
	p.score += np

	return p.score >= W
}

func (p *Player) won() bool {
	return p.score >= W
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %v\n", err)
		os.Exit(2)
	}

	re := regexp.MustCompile(`\d+`)
	nums := re.FindAllString(string(data), -1)
	pos1, _ := strconv.Atoi(nums[1])
	pos2, _ := strconv.Atoi(nums[3])

	p1 := Player{pos1, 0}
	p2 := Player{pos2, 0}

	fmt.Println(444356092776315 / math.MaxInt32)

	var d Dice
	for {
		if p1.move(&d) {
			break
		}

		if p2.move(&d) {
			break
		}
	}

	if p1.won() {
		fmt.Printf("%d x %d = %d\n", p2.score, d.rolls, p2.score*d.rolls)
	} else if p2.won() {
		fmt.Printf("%d x %d = %d\n", p1.score, d.rolls, p2.score*d.rolls)
	}
}
