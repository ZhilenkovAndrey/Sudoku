package main

import (
	."fmt"
	"math/rand"
)

const (height = 9
       width = 9)

type cell struct {
    number int
	x, y int
	fits bool
}

type greed [height - 1][width - 1] int

func cellChecking(c cell) bool{
    a := c.x / 3
	b := c.y / 3
    for i := 1; i <= 3; i++{
		for j := 1; i <= 3; i++{
	if c.number != g[a * 3 + i][b * 3 + j] {
        return true
	}
		return false
	}
}

func(c cell) cellFiling(i, j int, g *greed) {
    c.number = rand.Intn(10)
		    if cellChecking(c.number) {
                g[i][j] = c.number
			} else {
				cellFiling
			}
}

func(g *greed) sudokuBuilding() {
    for i := 1; i <= width; i++ {
        for j := 1; j <= height; j++ {
		     cellFiling
		}
	}
}