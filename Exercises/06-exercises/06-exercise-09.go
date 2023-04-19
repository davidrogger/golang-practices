/*
	- Passe Uma função como argumento a outra função
*/

package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func sumAllEven(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		if number%2 == 0 {
			total += number
		}
	}
	return total
}

func genericCaller(callback func(...int) int, numbers ...int) int {
	return callback(numbers...)
}

func main() {
	numbers := []int{2, 3, 4, 5, 8, 9, 11, 13, 16}
	total := genericCaller(sumAll, numbers...)
	justEvenTotal := genericCaller(sumAllEven, numbers...)

	fmt.Println(total)
	fmt.Println(justEvenTotal)
}
