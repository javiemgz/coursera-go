package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)

	var cow Animal
	cow.New("grass", "walk", "moo")

	var bird Animal
	bird.New("worms", "fly", "peep")

	var snake Animal
	snake.New("mice", "slither", "hsss")

	for {
		fmt.Print("> ")
		reader.Scan()
		command := reader.Text()
		command = strings.TrimSuffix(command, "\n")
		spplited := strings.Split(command, " ")

		switch spplited[0] {
		case "cow":
			doAction(spplited[1], cow)
		case "bird":
			doAction(spplited[1], bird)
		case "snake":
			doAction(spplited[1], snake)
		default:
			println("El animal no existe bro")
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

//Animal data for the animal
type Animal struct {
	foodEaten  string
	locomotion string
	sound      string
}

//New inits a new animal
func (animal *Animal) New(foodEaten, locomotion, sound string) {
	animal.foodEaten = foodEaten
	animal.locomotion = locomotion
	animal.sound = sound
}

func (animal *Animal) Eat() {
	fmt.Println(animal.foodEaten)
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.sound)
}
