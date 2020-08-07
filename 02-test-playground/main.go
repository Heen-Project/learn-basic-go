package main

import "fmt"

func main() {
	printOddEven(generateNumberFromZero(10))
}

func printOddEven(number []int) {
	for _, num := range number {
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}

func generateNumberFromZero(upTo int) []int {
	var numberSlice []int
	for i := 0; i <= upTo; i++ {
		numberSlice = append(numberSlice, i)
	}
	return numberSlice
}
