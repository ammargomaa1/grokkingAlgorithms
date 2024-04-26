package main

import "fmt"


func main() {
	fmt.Println(sortElements([]int{20, 40, 641, 41, 658, 9468}))
}

func getSmallestElementIndex(arr []int) int{
	smallestIndex := 0
	for index, element := range arr{
		if arr[smallestIndex] > element {
			smallestIndex = index
		}
	}

	return smallestIndex
}

func sortElements(arr []int) ([]int){
	sortedArr := [] int{}
	length := len(arr)
	for i := 0; i < length; i++ {
		smallestIndex := getSmallestElementIndex(arr)

		sortedArr = append(sortedArr, arr[smallestIndex])

		arr = append(arr[:smallestIndex], arr[smallestIndex+1:]...)
	}

	return sortedArr
	
}