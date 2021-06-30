package main

import "testing"

/*func TestPutToken01(t *testing.T) {
	b := &Board{
		tokens: []int{0,0,0,0,0,0,0,0,0}
	}
	b.put(1,1, "o")
	if b.get(1,1) != "o" {
		t.Errorf("....")
	}
}*/

func Test01(t *testing.T) {
	result := ""
	expected := "........."
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result += b.get(i, j)
		}
	}
	if result != expected {
		t.Errorf("Test01 Error")
	}

}

// 標準入力+putの確認
func Test02(t *testing.T) {

	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.input("o")
	if b.get(1, 2) != "o" {
		t.Errorf("Test02 Error")
	}

}

func Test03(t *testing.T) {
	expected := 1
	result := 0
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(0, 0, "o")
	b.put(0, 1, "o")
	b.put(0, 2, "o")
	result = b.hantei()
	if result != expected {
		t.Errorf("Test03 Error")
	}

}
