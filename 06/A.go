package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var nums [9]uint64

	toks := strings.Split(strings.TrimSpace(string(data)), ",")
	for _, t := range toks {
		n, _ := strconv.Atoi(t)
		nums[n]++
	}

	const days = 80
	for d := 0; d < days; d++ {
		z := nums[0]
		nums[0] = nums[1]
		nums[1] = nums[2]
		nums[2] = nums[3]
		nums[3] = nums[4]
		nums[4] = nums[5]
		nums[5] = nums[6]
		nums[6] = nums[7] + z
		nums[7] = nums[8]
		nums[8] = z
	}

	sum := uint64(0)
	for _, v := range nums {
		sum += v
	}
	fmt.Println(sum)
}
