package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/harris/test.txt")
	fmt.Println(err)
	fmt.Println(f)
}
