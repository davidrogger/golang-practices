package main

import "fmt"

func main() {
	t := justOdd(sumOfNumbers, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60)
	fmt.Println(t)
}

func sumOfNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func justOdd(callback func(x ...int) int, numbers ...int) int {
	var numbersOdd []int
	for _, number := range numbers {
		if number%2 != 0 {
			numbersOdd = append(numbersOdd, number)
		}
	}
	total := callback(numbersOdd...)
	return total
}
