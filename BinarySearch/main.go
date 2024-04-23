package main

import (
	"fmt"
)

func main() {

	fmt.Println(BinarySearch([]int{1, 2, 3, 4, 5, 6}, 10))
}

func BinarySearch(slice []int, number int) (index int) {
	lowIndex := 0
	sliceLen := len(slice) - 1
	highIndex := sliceLen

	for lowIndex <= highIndex {

		checkingIndex := (lowIndex + highIndex) / 2

		if slice[checkingIndex] < number {
			lowIndex = checkingIndex + 1
		} else if slice[checkingIndex] > number {
			highIndex = checkingIndex - 1
		}

		if slice[checkingIndex] == number {
			return checkingIndex
		}

	}

	return
}
