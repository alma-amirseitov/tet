package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
)

func parseFile(filename string)([]string,error){
	f,err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return nil,err
	}
	scanner := bufio.NewScanner(f)
	str := ""
	arr := make([]string,0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0{
			continue
		}
		if len(line) != 4{
			return nil,errors.New("bad format")
		}
		for _,v := range line{
			if v != '.' && v != '#' {
				return nil,errors.New("bad format")
			}
		}
		str += line
		if len(str) == 16{
			arr = append(arr,str)
			str = ""
		}
	}
	if len(arr) == 0{
		return nil,errors.New("bad format")
	}
	return arr,nil
}

func allocate(arr []string,l *List)(int,error){
	count := 0
	for k,v1 := range arr{
		if len(v1) != 16{
			return 0,errors.New("more than 4 line")
		}
		data := [4][2]int{}
		temp :=-1
		for j,v2 := range v1{
			if temp > 3{
				return 0,errors.New( "more than 4 # = " + string(temp))
			}
			if v2 == '#'{
				temp++
				x := j/4
				y := j%4
				data[temp] = [2]int{x,y}
			}
		}
		x := data[0][0]
		y := data[0][1]
		for i:=0;i<4;i++{
			data[i][0] -= x
			data[i][1] -= y
		}
		if !validate(data){
			return 0,errors.New("invalid data")
		}
		t := Tetris{
			label: rune(65+k),
			coordinates: data,
		}
		l.PushBack(t)
		count++
	}
	return count,nil
}

func validate(data [4][2]int)bool{
	count := 0

	for i:=0;i<4;i++ {
		for j:=0;j<4;j++ {
			if i == j {
				continue
			}
			x := int(math.Abs(float64(data[i][0] - data[j][0])))
			y := int(math.Abs(float64(data[i][1] - data[j][1])))
			if x + y == 1{
				count++
			}
		}
	}

	if count != 8 && count !=6{
		fmt.Println(count,data)
		return false
	}
	return true
}

/*
0  1  2  3
4  5  6  7
8  9  10 11
12 13 11  15

0,0
1,-1
1,0
1,1
2,0



*/