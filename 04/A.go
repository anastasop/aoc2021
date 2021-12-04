package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type randsrc struct {
	i    int
	nums []int
}

func (r *randsrc) next() int {
	if r.i >= len(r.nums) {
		return -1
	}
	n := r.nums[r.i]
	r.i++
	return n
}

func newRand(nums []int) *randsrc {
	r := &randsrc{}
	r.nums = make([]int, len(nums))
	copy(r.nums, nums)
	return r
}

type square struct {
	num    int
	marked bool
}

type board struct {
	sqs [5][5]square
}

func newBoard(nums []int) *board {
	b := new(board)
	k := 0
	for i := 0; i < len(b.sqs); i++ {
		for j := 0; j < len(b.sqs[i]); j++ {
			b.sqs[i][j].num = nums[k]
			k++
		}
	}
	return b
}

func (b *board) print(w io.Writer) {
	for i := 0; i < len(b.sqs); i++ {
		for j := 0; j < len(b.sqs[i]); j++ {
			fmt.Fprint(w, b.sqs[i][j].num, " ")
		}
		fmt.Fprintln(w, "")
	}
}

func (b *board) check(num int) (won bool, score int) {
	i, j := 0, 0

	for i = 0; i < len(b.sqs); i++ {
		for j = 0; j < len(b.sqs[i]); j++ {
			if b.sqs[i][j].num == num {
				b.sqs[i][j].marked = true
				goto marked
			}
		}
	}

	return

marked:
	if b.sqs[i][j].marked {
		won = true
		for k := 0; k < len(b.sqs[i]); k++ {
			won = won && b.sqs[i][k].marked
		}
		if !won {
			won = true
			for k := 0; k < len(b.sqs); k++ {
				won = won && b.sqs[k][j].marked
			}
		}
		if won {
			for i = 0; i < len(b.sqs); i++ {
				for j = 0; j < len(b.sqs[i]); j++ {
					if !b.sqs[i][j].marked {
						score += b.sqs[i][j].num
					}
				}
			}
			score *= num
		}
	}

	return
}

func main() {
	var rand *randsrc
	var boards []*board
	var err error
	var fields []string
	var nums []int

	data, err := io.ReadAll(os.Stdin)
	chk(err)

	sections := strings.Split(string(data), "\n\n")

	fields = strings.Split(sections[0], ",")
	nums = make([]int, len(fields))
	for i, s := range fields {
		nums[i], err = strconv.Atoi(s)
		chk(err)
	}
	rand = newRand(nums)

	for _, section := range sections[1:] {
		fields = strings.Fields(section)
		nums = make([]int, len(fields))
		for i, s := range fields {
			nums[i], err = strconv.Atoi(s)
			chk(err)
		}
		boards = append(boards, newBoard(nums))
	}

	for r := rand.next(); r >= 0; r = rand.next() {
		for i := range boards {
			if won, score := boards[i].check(r); won {
				fmt.Printf("Bingo! Board %d Score %d\n", i+1, score)
				os.Exit(0)
			}
		}
	}
	fmt.Println("No bingo????")
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
