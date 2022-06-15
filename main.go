package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

func measureTime(t time.Time){
	duration := time.Since(t)+time.Second
	fmt.Println("Time spended : ",duration+time.Second)
}

func main(){
	args := os.Args[1:]
	if len(args) == 1{
		arr,err := parseFile(args[0])
		if err != nil{
			fmt.Println("Error",err.Error())
			return
		}
		l := List{}
		count,err := allocate(arr,&l)
		if err != nil {
			fmt.Println("Error",err.Error())
			return
		}
		initialSize := int(math.Round(math.Sqrt(float64(count)*4)+0.4))
		b := solve(initialSize,&l)
		b.display()
	}else {
		fmt.Println("please provide 1 filename")
	}
}

func solve(size int,l *List) Board{
	board := BoardConstructor(size)

	var backtracking func(x,y int,node *Node) bool

	backtracking = func(x int,y int,node *Node) bool{
		if node == nil{
			return true
		}
		for x != board.size-1 && y != board.size-1{
			if board.square[x][y] != '.' || !board.CanPlace(x,y,node.Data){
				y++
				if y == board.size-1{
					x++
					y=0
				}
				continue
			}
			board.Place(x,y,node.Data)
			if backtracking(0,0,node.Next){
				return true
			}
			board.Remove(node.Data)
			y++
			if y == board.size-1{
				x++
				y=0
			}
		}
		return false
	}
	if backtracking(0,0,l.Head){
		return board
	}
	return solve(size+1,l)
}
