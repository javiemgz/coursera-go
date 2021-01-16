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

//Cow data for the animal
type Cow struct {
	name string
}

func (cow *Cow) setName(name string) {
	cow.name = name
}

func (cow *Cow) Eat() {
	fmt.Println("Grass")
}

func (cow *Cow) Move() {
	fmt.Println("Walk")
}

func (cow *Cow) Speak() {
	fmt.Println("Moo")
}

//Cow data for the animal
type Bird struct {
	name string
}

func (bird *Bird) setName(name string) {
	bird.name = name
}

func (bird *Bird) Eat() {
	fmt.Println("Worms")
}

func (bird *Bird) Move() {
	fmt.Println("Fly")
}

func (bird *Bird) Speak() {
	fmt.Println("Peeep")
}

//Cow data for the animal
type Snake struct {
	name string
}

func (snake *Snake) setName(name string) {
	snake.name = name
}

func (snake *Snake) Eat() {
	fmt.Println("Mice")
}

func (snake *Snake) Move() {
	fmt.Println("Slither")
}

func (snake *Snake) Speak() {
	fmt.Println("Hsss")
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	animals := make(map[string]Animal)

	for {
		fmt.Print("> ")
		reader.Scan()
		command := reader.Text()
		command = strings.TrimSuffix(command, "\n")
		spplited := strings.Split(command, " ")

		if spplited[0] == "newanimal" {
			switch spplited[2] {
			case "cow":
				cow := &Cow{}
				cow.setName(spplited[1])
				animals[spplited[1]] = cow
			case "bird":
				var bird Bird
				bird.setName(spplited[1])
				animals[spplited[1]] = &bird
			case "snake":
				var snake Snake
				snake.setName(spplited[1])
				animals[spplited[1]] = &snake
			default:
				println("El tipo de animal no existe bro")
			}
		}

		if spplited[0] == "query" {

			doAction(spplited[2], animals[spplited[1]])

		}
	}
}

func doAction(action string, animal Animal) {
	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		println("No existe la accion breo")
	}
}
