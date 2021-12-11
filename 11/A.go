package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	N       = 10
	STEPS   = 100
	FLASHED = 128
)

type octopus uint8

func (o *octopus) level() uint8  { return uint8(*o) & 0x0F }
func (o *octopus) flashed() bool { return uint8(*o)&FLASHED > 0 }
func (o *octopus) levelUp()      { *o++ }
func (o *octopus) flash()        { *o |= octopus(FLASHED) }

var octs [N][N]octopus

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %v\n", err)
		os.Exit(2)
	}
	lines := strings.Fields(string(data))

	for y := 0; y < len(octs); y++ {
		for x, r := range lines[y] {
			octs[y][x] = octopus(uint8(r - '0'))
		}
	}

	flashed := 0
	for step := 1; step <= STEPS; step++ {
		for y := 0; y < len(octs); y++ {
			for x := range octs[y] {
				octs[y][x].levelUp()
			}
		}

		for y := 0; y < len(octs); y++ {
			for x := range octs[y] {
				if octs[y][x].level() > 9 {
					diffuse(y, x)
				}
			}
		}

		for y := 0; y < len(octs); y++ {
			for x := range octs[y] {
				if octs[y][x].flashed() {
					flashed++
					octs[y][x] = octopus(0)
				}
			}
		}
	}
	fmt.Println(flashed)
}

func inGrid(y, x int) bool {
	return 0 <= x && x < len(octs[0]) && 0 <= y && y < len(octs)
}

func diffuse(y, x int) {
	if !inGrid(y, x) || octs[y][x].flashed() {
		return
	}

	if octs[y][x].level() > 9 {
		octs[y][x].flash()

		diffuse(y-1, x-1)
		diffuse(y-1, x)
		diffuse(y-1, x+1)
		diffuse(y, x-1)
		diffuse(y, x+1)
		diffuse(y+1, x-1)
		diffuse(y+1, x)
		diffuse(y+1, x+1)
	} else {
		octs[y][x].levelUp()
		if octs[y][x].level() > 9 {
			diffuse(y, x)
		}
	}
}
