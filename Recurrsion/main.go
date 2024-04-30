package main

import "fmt"

func main(){
	fmt.Println(factorial(10))
}

func factorial(num int) int {
	if num == 2 {
		return num
	}

	return num * factorial(num - 1)
}