package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	list := make([]int, 0)
	for i := 0; i < 10; i++ {
		var input string
		fmt.Println("Enter an integer number, type letter x to stop: ")
		fmt.Scan(&input)
		input = strings.ToLower(input)
		if input == "x" {
			break
		}
		num, _ := strconv.Atoi(input)
		list = append(list, num)
	}
	bubbleSort(list)
	fmt.Print(list)
}

func bubbleSort(slice []int) {
	for {
		swapped := false
		for i := 0; i < len(slice)-1; i++ {
			if slice[i+1] < slice[i] {
				swap(slice, i)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func swap(slice []int, index int) {
	value1 := slice[index]
	slice[index] = slice[index+1]
	slice[index+1] = value1
}
