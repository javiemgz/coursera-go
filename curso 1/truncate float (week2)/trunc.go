package main

import "fmt"

func main() {
	var number float32
	fmt.Print("Enter a float number in order to truncate it: ")
	fmt.Scan(&number)
	truncatedNumber := int(number)
	fmt.Print("Truncated value is: ", truncatedNumber)
}
