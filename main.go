package main

import "fmt"

func main() {
	fmt.Print("working", "\n")
	// fmt.Println(test(20))
	x, y := test(20)
	fmt.Print(x, y, "\n")
	fmt.Println(test1())
}

func test(x int) (int, int) {
	return x, x + 10
}

func test1() int {
	return 20
}
