package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	list := make([]int, 3)
	iterationNumber := 0 //control first 3 values
	indexToInsert := 0   // handle negative numbers in first 3 itertions

	for {
		var input string

		fmt.Println("Enter an intergeer number: ")
		fmt.Scan(&input)
		input = strings.ToLower(input)
		if input == "x" {
			break
		}

		num, _ := strconv.Atoi(input)

		if iterationNumber < 3 {
			list[indexToInsert] = num
			if num < 0 {
				indexToInsert++
			}
			iterationNumber++
		} else {
			list = append(list, num)
		}
		sort.Ints(list)

		fmt.Println("The list is ", list)
	}
}
