package main

import "fmt"

type sp interface{ Speak() }

type Dog struct {
	name string
}

func (d *Dog) Speak() {
	if d == nil {
		fmt.Println("Empty")
	} else {
		fmt.Println(d.name)
	}

}

func (d Dog) Checking() {
	fmt.Println("checking")
}

func (d *Dog) set(val string) {
	// fmt.Println(d.)
	d.name = val
}

func main() {
	var s1 sp
	var d1 *Dog
	d2 := Dog{"ramu"}
	d2.Speak()
	// d1.set("july")
	// d2 := d1{"ramu"}
	s1 = d1
	s1.Speak()
	// s1.checking()
}
