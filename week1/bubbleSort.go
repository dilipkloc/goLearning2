package main

import "fmt"

func main() {
	values := make([]int, 10)
	fmt.Println("Enter the 10 integers numbers:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&values[i])
	}
	BubbleSort(values)
	fmt.Println("Sorted Elements are:")
	for i := 0; i < 10; i++ {
		fmt.Println(values[i])
	}

}

func BubbleSort(values []int) {
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < 10-1; i++ {
			if values[i] > values[i+1] {
				sorted = false
				Swap(values, i)
			}
		}
	}
}

func Swap(values []int, index int) {
	pre := values[index]
	values[index] = values[index+1]
	values[index+1] = pre
}
