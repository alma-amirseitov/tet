package main

import "fmt"

type Tetris struct {
	label       rune
	coordinates [4][2]int
	placed bool
}

type Node struct {
	Data Tetris
	Next *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func (l *List) PushBack(data Tetris){
	node := &Node{
		Data: data,
	}
	if l.Head == nil{
		l.Head = node
		l.Tail = node
	}else{
		l.Tail.Next = node
		l.Tail = node
	}
}

func (l *List) Display() {
	list := l.Head
	for list != nil {
		fmt.Printf("%+v ->", string(list.Data.label))
		list = list.Next
	}
}

func (l *List)allPlaced()bool{
	n := l.Head
	for n != nil{
		if !n.Data.placed{
			return false
		}
		n = n.Next
	}
	return true
}