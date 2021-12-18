package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	LBRACKET = math.MaxInt
	RBRACKET = math.MaxInt - 1
	COMMA    = math.MaxInt - 2
)

var (
	scanRe = regexp.MustCompile(`\d+|\[|\]|,`)
)

type elem struct {
	v    int
	prev *elem
	next *elem
}

type number struct {
	head *elem
	tail *elem
}

func (n *number) appendElem(v int) {
	t := &elem{v, nil, nil}
	if n.head == nil {
		n.head = t
		n.tail = t
	} else if n.tail != nil {
		n.tail.next = t
		t.prev = n.tail
		n.tail = t
	}
}

func (n *number) appendNumber(o *number) {
	n.tail.next = o.head
	o.head.prev = n.tail
	n.tail = o.tail
}

func Parse(s string) *number {
	n := &number{}

	toks := scanRe.FindAllString(s, -1)
	for _, tok := range toks {
		var v int
		switch tok {
		case "[":
			v = LBRACKET
		case "]":
			v = RBRACKET
		case ",":
			v = COMMA
		default:
			v, _ = strconv.Atoi(tok)
		}
		n.appendElem(v)
	}
	return n
}

func (n *number) String() string {
	var b strings.Builder
	for t := n.head; t != nil; t = t.next {
		switch t.v {
		case LBRACKET:
			b.WriteString("[")
		case RBRACKET:
			b.WriteString("]")
		case COMMA:
			b.WriteString(",")
		default:
			b.WriteString(fmt.Sprintf("%d", t.v))
		}
	}
	return b.String()
}

func (n *number) StringNums() string {
	var b strings.Builder
	digits := "0123456789"
	open := 0
	for t := n.head; t != nil; t = t.next {
		switch t.v {
		case LBRACKET:
			open++
			b.WriteString(digits[open : open+1])
		case RBRACKET:
			b.WriteString(digits[open : open+1])
			open--
		case COMMA:
			b.WriteString(",")
		default:
			b.WriteString(fmt.Sprintf("(%d)", t.v))
		}
	}
	return b.String()
}

func magAux(e *elem, stop *elem) int {
	if e.v < COMMA {
		return e.v
	}

	var t *elem
	var open int
loop:
	for t = e; t != stop; t = t.next {
		switch t.v {
		case LBRACKET:
			open++
		case RBRACKET:
			open--
		case COMMA:
			if open == 1 {
				break loop
			}
		}
	}
	if t == stop {
		panic("stop")
	}

	if stop != nil {
		stop = stop.prev
	}

	return 3*magAux(e.next, t) + 2*magAux(t.next, stop)
}

func (n *number) mag() int {
	return magAux(n.head, nil)
}

func add(n1, n2 *number) *number {
	h := &number{}
	h.appendElem(LBRACKET)
	h.appendNumber(n1)
	h.appendElem(COMMA)
	h.appendNumber(n2)
	h.appendElem(RBRACKET)
	return h
}

func explode(n *number, e *elem) {
	p0 := e.v
	p1 := e.next.v
	p2 := e.next.next.v
	p3 := e.next.next.next.v
	p4 := e.next.next.next.next.v

	b := p0 == LBRACKET && p1 < COMMA && p2 == COMMA && p3 < COMMA && p4 == RBRACKET
	if !b {
		panic("explode error")
	}

	ne := &elem{0, e.prev, e.next.next.next.next.next}
	e.prev.next = ne
	e.next.next.next.next.next.prev = ne

	for p := ne.prev; p != nil; p = p.prev {
		if p.v < COMMA {
			p.v += p1
			break
		}
	}
	for p := ne.next; p != nil; p = p.next {
		if p.v < COMMA {
			p.v += p3
			break
		}
	}
}

func split(nn *number, e *elem) {
	var a, b int
	if e.v%2 == 1 {
		a = e.v / 2
		b = a + 1
	} else {
		a = e.v / 2
		b = a
	}

	n := &number{}
	n.appendElem(LBRACKET)
	n.appendElem(a)
	n.appendElem(COMMA)
	n.appendElem(b)
	n.appendElem(RBRACKET)

	e.prev.next = n.head
	n.head.prev = e.prev
	e.next.prev = n.tail
	n.tail.next = e.next
}

func (n *number) reduce() {
	for {
		open := 0
		var t *elem

		var forExplode, forSplit *elem = nil, nil
		for t = n.head; t != nil; t = t.next {
			switch t.v {
			case LBRACKET:
				open++
				if open > 4 {
					if forExplode == nil {
						forExplode = t
					}
				}
			case RBRACKET:
				open--
			case COMMA:
				// ignore
			default:
				if t.v >= 10 {
					if forSplit == nil {
						forSplit = t
					}
				}
			}
		}
		if forExplode == nil && forSplit == nil {
			return
		}
		if forExplode != nil {
			explode(n, forExplode)
		} else if forSplit != nil {
			split(n, forSplit)
		}
	}
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %v\n", err)
		os.Exit(2)
	}

	lines := strings.Split(string(data), "\n")
	lines = lines[0 : len(lines)-1]

	var max int
	for x, l1 := range lines {
		for y, l2 := range lines {
			if x != y {
				nn1 := Parse(l1)
				nn2 := Parse(l2)
				n := add(nn1, nn2)
				n.reduce()
				if m := n.mag(); m > max {
					max = m
				}

				nn1 = Parse(l1)
				nn2 = Parse(l2)
				n = add(nn2, nn1)
				n.reduce()
				if m := n.mag(); m > max {
					max = m
				}
			}
		}
	}
	fmt.Println(max)
}
