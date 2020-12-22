package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var filename string
	type Person struct {
		fname string
		lname string
	}
	people := make([]Person, 0)

	fmt.Print("Enter filename: ")
	fmt.Scan(&filename)

	file, _ := os.Open(filename)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	file.Close()

	for _, eachline := range lines {
		splitted := strings.Split(eachline, " ")
		people = append(people, Person{splitted[0], splitted[1]})
	}

	for i, person := range people {
		fmt.Printf("person %d \nName: %s Last Name: %s \n", i, person.fname, person.lname)
	}

}
