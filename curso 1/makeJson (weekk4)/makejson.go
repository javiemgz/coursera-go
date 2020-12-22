package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := make(map[string]string)
	var input string
	fmt.Println("Enter your name: ")
	fmt.Scan(&input)
	data["name"] = input
	fmt.Println("Enter your adress: ")
	fmt.Scan(&input)
	data["adress"] = input

	convertedAsByte, _ := json.Marshal(data)
	convertedAsString := string(convertedAsByte)
	fmt.Println(convertedAsString)
}
