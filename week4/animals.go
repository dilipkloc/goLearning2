package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name  string
	food  string
	move  string
	speak string
}

func (c *Cow) Eat() {
	fmt.Println(c.name, " eats ", c.food)
}

func (c *Cow) Move() {
	fmt.Println(c.name, " Movements ", c.move)
}

func (c *Cow) Speak() {
	fmt.Println(c.name, " eats ", c.speak)
}

type Bird struct {
	name  string
	food  string
	move  string
	speak string
}

func (c *Bird) Eat() {
	fmt.Println(c.name, " eats ", c.food)
}

func (c *Bird) Move() {
	fmt.Println(c.name, " Movements ", c.move)
}

func (c *Bird) Speak() {
	fmt.Println(c.name, " eats ", c.speak)
}

type Snake struct {
	name  string
	food  string
	move  string
	speak string
}

func (c *Snake) Eat() {
	fmt.Println(c.name, " eats ", c.food)
}

func (c *Snake) Move() {
	fmt.Println(c.name, " Movements ", c.move)
}

func (c *Snake) Speak() {
	fmt.Println(c.name, " eats ", c.speak)
}

func main() {
	var c []Cow
	var b []Bird
	var s []Snake
	fmt.Println("Enter your requests")
	fmt.Print(">")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	request := scanner.Text()
	requestSplit := strings.Fields(request)
	// fmt.Println(len(requestSplit))
	for len(requestSplit) == 3 {

		// fmt.Println(len(c))
		switch requestSplit[0] {
		case "newanimal":
			switch requestSplit[2] {
			case "cow":
				found := false
				for _, c1 := range c {
					if c1.name == requestSplit[1] {
						fmt.Println("Name already exists")
						found = true
						break
					}
				}
				if !found {
					c = append(c, Cow{requestSplit[1], "grass", "walk", "moo"})
					fmt.Println("Created it!")
				}

			case "bird":
				found := false
				for _, c1 := range c {
					if c1.name == requestSplit[1] {
						fmt.Println("Name already exists")
						found = true
						break
					}
				}
				if !found {
					b = append(b, Bird{requestSplit[1], "worms", "fly", "peep"})
					fmt.Println("Created it!")
				}
			case "snake":
				found := false
				for _, c1 := range c {
					if c1.name == requestSplit[1] {
						fmt.Println("Name already exists")
						found = true
						break
					}
				}
				if !found {
					s = append(s, Snake{requestSplit[1], "mice", "slither", "hsss"})
					fmt.Println("Created it!")
				}
			}

		case "query":
			found := false
			for _, c1 := range c {
				if c1.name == requestSplit[1] {
					c1.PrintRequestCOW(requestSplit[2])
					found = true
				}
			}
			if !found {
				for _, c1 := range b {
					if c1.name == requestSplit[1] {
						c1.PrintRequestBird(requestSplit[2])
						found = true
					}
				}
			}
			if !found {
				for _, c1 := range s {
					if c1.name == requestSplit[1] {
						c1.PrintRequestSnake(requestSplit[2])
						found = true
					}
				}
			}
		}
		fmt.Println("Enter your requests")
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		request = scanner.Text()
		requestSplit = strings.Fields(request)
	}
}

func (c Cow) PrintRequestCOW(str string) {
	switch str {
	case "eat":
		c.Eat()
	case "move":
		c.Move()
	case "speak":
		c.Speak()
	}
}

func (c Bird) PrintRequestBird(str string) {
	switch str {
	case "eat":
		c.Eat()
	case "move":
		c.Move()
	case "speak":
		c.Speak()
	}
}

func (c Snake) PrintRequestSnake(str string) {
	switch str {
	case "eat":
		c.Eat()
	case "move":
		c.Move()
	case "speak":
		c.Speak()
	}
}
