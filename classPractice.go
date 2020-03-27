package main

import "fmt"

type Point struct {
	x int
	y int
}

func (p Point) updateXY(x1, y1 int) {
	p.x += x1
	p.y += y1
	fmt.Println(p.y)
}

func (p Point) PrintXY() {
	fmt.Println(p.x)
	fmt.Println(p.y)
}

func (p *Point) updateXYPointer(v int) {
	p.x += v
	p.y += v
}

func main() {
	var p Point //Cannot initialize any values this type of declaration
	p.PrintXY()
	p.updateXY(1, 2) // will not update original values
	p.PrintXY()
	p1 := Point{5, 3} // can declare initialization using this methods
	p1.PrintXY()
	p1.updateXYPointer(10) // will update original values cause using call by reference in function * is used
	p1.PrintXY()
}
