package main

import (
	"bufio"
	"fmt"
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

func check() {
	tests := []struct {
		num string
		mag int
	}{
		{`[[1,2],[[3,4],5]]`, 143},
		{`[[[[0,7],4],[[7,8],[6,0]]],[8,1]]`, 1384},
		{`[[[[1,1],[2,2]],[3,3]],[4,4]]`, 445},
		{`[[[[3,0],[5,3]],[4,4]],[5,5]]`, 791},
		{`[[[[5,0],[7,4]],[5,5]],[6,6]]`, 1137},
		{`[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]`, 3488},
	}

	for _, test := range tests {
		n := Parse(test.num)
		if t := n.String(); t != test.num {
			fmt.Fprintf(os.Stderr, "parse failed:\norig: %s\ncopy: %s\n", test.num, t)
			os.Exit(2)
		}
		if m := n.mag(); m != test.mag {
			fmt.Fprintf(os.Stderr, "mag failed for %s:\norig: %d\ncopy: %d\n", test.num, test.mag, m)
			os.Exit(2)
		}
	}

	nums := []*number{
		Parse(`[1,1]`),
		Parse(`[2,2]`),
		Parse(`[3,3]`),
		Parse(`[4,4]`),
	}
	n := nums[0]
	for _, t := range nums[1:] {
		n = add(n, t)
	}
	if n.String() != `[[[[1,1],[2,2]],[3,3]],[4,4]]` {
		fmt.Fprintf(os.Stderr, "add failed")
		os.Exit(2)
	}
}

func main() {
	check()
	checkReduce()

	var n *number
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nn := Parse(scanner.Text())
		if n == nil {
			n = nn
		} else {
			n = add(n, nn)
			n.reduce()
		}
	}
	fmt.Println(n)
	fmt.Println(n.mag())
}

func checkReduce() {
	tests := [][]string{
		{
			`[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]`,
			`[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]`,
			`[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]`,
		},

		{
			`[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]`,
			`[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]`,
			`[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]`,
		},

		{
			`[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]`,
			`[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]`,
			`[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]`,
		},

		{
			`[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]`,
			`[7,[5,[[3,8],[1,4]]]]`,
			`[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]`,
		},

		{
			`[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]`,
			`[[2,[2,2]],[8,[8,1]]]`,
			`[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]`,
		},

		{
			`[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]`,
			`[2,9]`,
			`[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]`,
		},

		{
			`[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]`,
			`[1,[[[9,3],9],[[9,0],[0,7]]]]`,
			`[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]`,
		},

		{
			`[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]`,
			`[[[5,[7,4]],7],1]`,
			`[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]`,
		},

		{
			`[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]`,
			`[[[[4,2],2],6],[8,7]]`,
			`[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]`,
		},
	}

	for i, test := range tests {
		n1 := Parse(test[0])
		n2 := Parse(test[1])
		addmsg := fmt.Sprintf("%s + %s", n1, n2)
		n := add(n1, n2)
		n.reduce()
		if test[2] != n.String() {
			fmt.Printf("Error: %d\n%s\nGot:\t%s\nIs:\t%s\n", i+1, addmsg, n.String(), test[2])
			break
		}
	}
}
