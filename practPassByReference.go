package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	foo(a)
	fmt.Print(a[1])
	c := [3]int{4, 5, 6}
	test2(&c)
}

func foo(b []int) {
	b[1] = 20
}

func test2(d *[3]int) {
	fmt.Print(d[1])
}
