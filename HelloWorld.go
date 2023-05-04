package main

import (
	"fmt"
)

var y int = 90

func outer() {
	var s4 string = "Hello"
	fmt.Println(s4)
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	outer()
	fmt.Printf("%T\n", y) //型を教えてくれる
}

//go buildでコンパイル
