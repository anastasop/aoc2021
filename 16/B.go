package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

type stream struct {
	data []byte
}

func newStream(b []byte) *stream {
	return &stream{data: b}
}

func (s *stream) subStream(n int) *stream {
	b := s.data[0:n]
	s.data = s.data[n:]
	return newStream(b)
}

func (s *stream) eof() bool {
	return len(s.data) == 0
}

func (s *stream) read(n int) int {
	if len(s.data) < n {
		fmt.Fprintf(os.Stderr, "requested %d got %d\n", n, len(s.data))
		os.Exit(2)
	}

	v := 0
	for _, b := range s.data[0:n] {
		v = v*2 + int(b-'0')
	}
	s.data = s.data[n:]
	return v
}

type parser struct {
	ist   *stream
	limit int
}

func newParser(s *stream, l int) *parser {
	return &parser{ist: s, limit: l}
}

func (p *parser) read(n int) int {
	return p.ist.read(n)
}

func (p *parser) parserOfLength(n int) *parser {
	return newParser(p.ist.subStream(n), -1)
}

func (p *parser) parserOfPackets(n int) *parser {
	return newParser(p.ist, n)
}

func (p *parser) eof() bool {
	return p.ist.eof()
}

type Emitter interface {
	Emit(v int)
}

type Accumulator struct {
	values []int
}

func (a *Accumulator) Emit(v int) {
	a.values = append(a.values, v)
}

func (a *Accumulator) sum() int {
	s := 0
	for _, v := range a.values {
		s += v
	}
	return s
}

func (a *Accumulator) mul() int {
	s := 1
	for _, v := range a.values {
		s *= v
	}
	return s
}

func (a *Accumulator) min() int {
	s := 100000000
	for _, v := range a.values {
		if v < s {
			s = v
		}
	}
	return s
}

func (a *Accumulator) max() int {
	s := 0
	for _, v := range a.values {
		if v > s {
			s = v
		}
	}
	return s
}

func (a *Accumulator) gt() int {
	if a.values[0] > a.values[1] {
		return 1
	}
	return 0
}

func (a *Accumulator) lt() int {
	if a.values[0] < a.values[1] {
		return 1
	}
	return 0
}

func (a *Accumulator) eq() int {
	if a.values[0] == a.values[1] {
		return 1
	}
	return 0
}

func (p *parser) parse(emitter Emitter) {
	nread := 0
	for {
		if nread == p.limit {
			return
		}

		if p.eof() {
			if p.limit == -1 {
				return
			} else {
				fmt.Fprintf(os.Stderr, "unexpected EOF")
				os.Exit(2)
			}
		}

		_ = p.read(3) // version
		typ := p.read(3)
		nread++

		if typ == 4 {
			v := 0
			for {
				i := p.read(5)
				v = v<<4 + i&0x0F
				if i <= 0x0F {
					break
				}
			}
			emitter.Emit(v)
		} else {
			l := p.read(1)
			var acc Accumulator
			if l == 0 {
				p.parserOfLength(int(p.read(15))).parse(&acc)
			} else if l == 1 {
				p.parserOfPackets(int(p.read(11))).parse(&acc)
			}
			switch typ {
			case 0:
				emitter.Emit(acc.sum())
			case 1:
				emitter.Emit(acc.mul())
			case 2:
				emitter.Emit(acc.min())
			case 3:
				emitter.Emit(acc.max())
			case 5:
				emitter.Emit(acc.gt())
			case 6:
				emitter.Emit(acc.lt())
			case 7:
				emitter.Emit(acc.eq())
			}
		}
	}

}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %v\n", err)
		os.Exit(2)
	}

	bin := make([]byte, 0, len(data)*4)
	for _, b := range bytes.TrimSpace(data) {
		v, _ := strconv.ParseInt(string(b), 16, 64)
		bin = append(bin, []byte(fmt.Sprintf("%04b", v))...)
	}

	var acc Accumulator
	parser := newParser(newStream(bin), -1)
	parser.parserOfPackets(1).parse(&acc)
	fmt.Println(acc.values[0])
}
