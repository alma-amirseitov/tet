package main

import "fmt"

type Board struct {
	square [][]rune
	size  int
}


func BoardConstructor(n int) Board {
	var square [][]rune
	for i := 0; i < n; i++ {
		temp := ""
		for j := 0; j < n; j++ {
			temp += "."
		}
		square = append(square, []rune(temp))
	}
	return Board{square, n,}
}

func (b *Board) Place(x, y int, t Tetris) {
	for _, v := range t.coordinates {
		b.square[x+v[0]][y+v[1]] = t.label
	}
}

func (b *Board) CanPlace(x, y int, t Tetris) bool {
	if b.square[x][y] != '.' {
		return false
	}
	l := len(b.square)
	for _, v := range t.coordinates {
		if x+v[0] < 0 || y+v[1] < 0 {
			return false
		}
		if x+v[0] >= l || y+v[1] >= l {
			return false
		}
		if (b.square)[x+v[0]][y+v[1]] != '.' {
			return false
		}
	}
	return true
}

func (b *Board) display(){
	for i := 0; i < len(b.square); i++ {
		for j := 0; j < len(b.square[i]); j++ {
			fmt.Print(string(b.square[i][j]), " ")
		}
		fmt.Println()
	}
}
func (b *Board) Remove(t Tetris) {
	for i := 0; i < len(b.square); i++ {
		for j := 0; j < len(b.square[i]); j++ {
			if b.square[i][j] == t.label {
				b.square[i][j] = '.'
			}
		}
	}
}