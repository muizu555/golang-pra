package main

import "fmt"

//配列型  後から要素数を変更できない

func main() {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1) //型を調べる時はPrintf

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]int{1, 2} //自動で要素数を決めてくれる
	fmt.Println(arr4)

	fmt.Println(arr2[0])

	fmt.Println(arr4[0])

	arr2[2] = "c"
	fmt.Println(arr2)

	fmt.Println(len(arr1)) //要素数を出せる

}
