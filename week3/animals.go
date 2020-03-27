package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animals struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animals) Eat() {
	fmt.Println("Eats ", a.food, "\n\n")
}

func (a *Animals) Move() {
	fmt.Println("Moves ", a.locomotion, "\n\n")
}

func (a *Animals) Speak() {
	fmt.Println("Speaks ", a.noise, "\n\n")
}

func (a *Animals) CallFunc(str string) {
	switch str {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "noise":
		a.Speak()
	default:
		fmt.Println("Wrong input")
	}
}

func main() {
	defer fmt.Println("Exiting")
	cow := Animals{"grass", "walk", "moo"}
	bird := Animals{"worms", "fly", "peep"}
	snake := Animals{"mice", "slither", "hsss"}
	fmt.Println(`Enter your requests:

	Your requests should be Animal_name attribute			
	Animal_name: cow, bird, snake
	attributes: eat,move,noise
	Enter -1 to exit	
		`)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	request := scanner.Text()
	requestSplit := strings.Fields(request)
	fmt.Println(len(requestSplit))

	for len(requestSplit) == 2 {

		switch requestSplit[0] {
		case "cow":
			cow.CallFunc(requestSplit[1])
		case "bird":
			bird.CallFunc(requestSplit[1])
		case "snake":
			snake.CallFunc(requestSplit[1])
		default:
			fmt.Println("Wrong input")
		}

		fmt.Println(`Enter your requests:

	Your requests should be Animal_name attribute			
	Animal_name: cow, bird, snake
	attributes: eat,move,noise
	Enter -1 to exit	
		`)
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		request = scanner.Text()
		requestSplit = strings.Fields(request)

	}

}
