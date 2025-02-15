package main

import (
	"errors"
	. "fmt"
)

const (
	height = 9
	width  = 9
)

var (
	ErrDigitSquare = errors.New("В квадрате есть такая цифра.")
	ErrDigitLine   = errors.New("В линии есть такая цифра")
	ErrDigitСolumn = errors.New("в столбце есть такая цифра")
)

type cell struct {
	number int
	x, y   int
	fits   bool
}

type greed [height][width]int

func (g greed) cellDigitChecking(c cell) bool {
	var a, b int
	var d bool
	a = c.x / 3
	b = c.y / 3
	for i := 1; i <= 3; i++ {
		for j := 1; i <= 3; i++ {
			if c.number != g[a*3+i][b*3+j] {
				d = true
			} else {
				Println(ErrDigitSquare)
				d = false
			}
		}
	}
	return d
}

func (g greed) cellLineChecking(c cell) bool {
	var d bool
	for i := 0; i < width; i++ {
		if c.number != g[i][c.y] {
			d = true
		} else {
			d = false
			Println(ErrDigitLine)
		}
	}
	return d
}

func (g greed) cellColumnChecking(c cell) bool {
	var d bool
	for i := 0; i < height; i++ {
		if c.number != g[c.x][i] {
			d = true
		} else {
			d = false
			Println(ErrDigitСolumn)
		}
	}
	return d
}

// func(c cell) cellFiling(i, j int, g greed) {
//     c.number = rand.Intn(10)
// 		    if cellChecking(c.number) {
//                 g[i][j] = c.number
// 			} else {
// 				cellFiling
// 			}
// }

// func(g *greed) sudokuBuilding() {
//     for i := 1; i <= width; i++ {
//         for j := 1; j <= height; j++ {
// 		     cellFiling(i, j, g)
// 		}
// 	}
// }
