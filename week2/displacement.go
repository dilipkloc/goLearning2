package main

import "fmt"

func main() {
	var a, v, t, s float64
	fmt.Println("Enter the acceleration:")
	fmt.Scan(&a)
	fmt.Println("Enter the initial velocity:")
	fmt.Scan(&v)
	fmt.Println("Enter the initial displacement:")
	fmt.Scan(&s)
	fn := GenDisplaceFn(a, v, s)
	fmt.Println("Enter the time:")
	fmt.Scan(&t)
	fmt.Println(fn(t))
}

func GenDisplaceFn(a float64, v float64, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v*t + s
	}
}
