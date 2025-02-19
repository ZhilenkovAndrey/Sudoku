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
	ErrDigitFix    = errors.New("Эта кцифра неизменяема")
)

type cell struct {
	x, y, number int
	fix          bool
}

type greed [height][width]cell

func newSudoku(greed1 [width][height]int) *greed { //Заполнение поля сверху массивом greed1
	var greed2 greed
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			greed2[i][j].number = greed1[i][j]
			greed2[i][j].x = i
			greed2[i][j].y = j
		}
	}
	return &greed2
}

func fieldInitial() *greed { // Заполнение поля вручную массивом инициализированным композитным литералом

	return newSudoku([width][height]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})
}

func (g greed) fixedField() *greed { //Делает все значения не равные нулю неизменяемыми

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if g[i][j].number != 0 {
				g[i][j].fix = true
			}
		}
	}
	return &g
}

func (g *greed) printField() {
	for i := 0; i < width; i++ {
		Println(g[i])
	}
}

func (g *greed) cellDigitChecking(c cell) bool { //Проверка квадрата, в котором ставится цифра на отсутствие совпадений
	var a, b int
	var d bool
	a = c.x / 3
	b = c.y / 3
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if c.number != g[a*3+i][b*3+j].number && c.x != i && c.y != j {
				d = true
			} else {
				Println(ErrDigitSquare)
				d = false
			}
		}
	}
	return d
}

func (g *greed) cellLineChecking(c cell) bool { //Проверка гор линии, в которую ставится цифра, на несовпадение
	var d bool
	for i := 0; i < width; i++ {
		if c.number != g[i][c.y].number {
			d = true
		} else {
			d = false
			Println(ErrDigitLine)
		}
	}
	return d
}

func (g *greed) cellColumnChecking(c cell) bool { //Проверка вертикальной линии, в которой ставится цифра, на несовппадение
	var d bool
	for i := 0; i < height; i++ {
		if c.number != g[c.x][i].number {
			d = true
		} else {
			d = false
			Println(ErrDigitСolumn)
		}
	}
	return d
}

func (g *greed) cellNotFix(c cell) bool { //Проверка заполняемой клетки на изменяемость
	var d bool
	if g[c.x][c.y].fix == true {
		d = false
		Println(ErrDigitFix)
	} else {
		d = true
	}
	return d
}

func (g *greed) addDigit(c cell) { //Проверка возможности установки новой цифры в клетке
	if g.cellDigitChecking(c) && g.cellLineChecking(c) && g.cellColumnChecking(c) && g.cellNotFix(c) {
		g[c.x][c.y].number = c.number
	}

}

func main() {
	fieldInitial().fixedField().printField()
}