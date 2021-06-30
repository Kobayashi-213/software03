package main

import (
	"fmt"
	"os"
)

type Board struct {
	tokens []int // tokens[0] -> (0,0), tokens[1] -> (0,1), ...
}

func (b *Board) put(x, y int, u string) {

	if b.get(x, y) != "." {
		fmt.Println("Cannot put here!")
	} else if u == "o" {
		b.tokens[x+3*y] = 1
	} else if u == "x" {
		b.tokens[x+3*y] = 2
	}
}

func (b *Board) get(x, y int) string {
	if b.tokens[x+3*y] == 1 {
		return "o"
	}
	if b.tokens[x+3*y] == 2 {
		return "x"
	} else {
		return "."
	}
}

func (b *Board) input(u string) {
	//Player1 がxとする
	var x, y int
	if u == "x" {
		fmt.Printf("Player 2: Input (x,y):")
	} else if u == "o" {
		fmt.Printf("Player 1: Input (x,y):")
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

func (b *Board) hantei() int {
	for i := 0; i < 3; i++ {
		s := ""
		for j := 0; j < 3; j++ {
			s += b.get(i, j)
		}
		if s == "ooo" {
			return 1
		} else if s == "xxx" {
			return 2
		}
	}

	for i := 0; i < 3; i++ {
		s := ""
		for j := 0; j < 3; j++ {
			s += b.get(j, i)
		}
		if s == "ooo" {
			return 1
		} else if s == "xxx" {
			return 2
		}
	}
	s := ""
	s = b.get(0, 0) + b.get(1, 1) + b.get(2, 2)
	if s == "ooo" {
		return 1
	} else if s == "xxx" {
		return 2
	}

	s = ""
	s = b.get(0, 2) + b.get(1, 1) + b.get(2, 0)
	if s == "ooo" {
		return 1
	} else if s == "xxx" {
		return 2
	}
	return 0
}

func main() {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	result := ""

	for i := 0; i < 9; i++ {
		if i%2 == 0 {
			b.input("o")
			for i := 0; i < 3; i++ {
				result = ""
				for j := 0; j < 3; j++ {
					result += b.get(i, j)
				}
				fmt.Println(result)
			}
			if b.hantei() == 1 {
				fmt.Println("Player 1 won")
				os.Exit(3)
			}
		} else if i%2 != 0 {
			b.input("x")
			for i := 0; i < 3; i++ {
				result = ""
				for j := 0; j < 3; j++ {
					result += b.get(i, j)
				}
				fmt.Println(result)
			}
			if b.hantei() == 2 {
				fmt.Println("Player 2 won")
				os.Exit(3)
			}
		}
	}
	fmt.Println()
}
