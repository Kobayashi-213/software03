package main

import (
	"fmt"
)

type Board struct {
	tokens []int // tokens[0] -> (0,0), tokens[1] -> (0,1), ...
}

func (b *Board) put(x, y int, u string) {
	if u == "o" {
		b.tokens[x+3*y] = 1
	}
}

func (b *Board) get(x, y int) string {
	if b.tokens[x+3*y] == 1 {
		return "o"
	}
	if b.tokens[x+3*y] == 0 {
		return "."
	}

	return "x"
}

func (b *Board) input(u string) {
	//Player1 がxとする
	var x, y int
	if u == "x" {
		fmt.Printf("Player 1: Input (x,y):")
	} else if u == "o" {
		fmt.Printf("Player 2: Input (x,y):")
	}
	//test用
	x = 1
	y = 2
	//標準入力から受け取る
	fmt.Scan(&x)
	fmt.Scan(&y)

	b.put(x, y, u)
	return
}

func main() {

}
