package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a string: ")
	stringValue, _ := reader.ReadString('\n')

	stringValue = strings.TrimSpace(stringValue)
	stringValue = strings.ToLower(stringValue)
	fmt.Printf("value: %s \n", stringValue)

	isFound := strings.HasPrefix(stringValue, "i")
	isFound = isFound && strings.Contains(stringValue, "a")
	isFound = isFound && strings.HasSuffix(stringValue, "n")

	if isFound {
		fmt.Print("Found!")
	} else {
		fmt.Print("Not Found!")
	}
}
