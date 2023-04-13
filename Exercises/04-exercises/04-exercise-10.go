/**
	- Utilizando o exercício anterior, remova uma entrada do
	map e demonstre o map inteiro utilizando range.
**/

package main

import "fmt"

func main() {
	users := map[string][]string{
		"doe_Jonas": []string{"Mock", "testing"},
	}

	users["modiz"] = []string{"eat", "sleep"}

	delete(users, "doe_Jonas")

	for name, hobbies := range users {
		fmt.Println("User:", name)
		for index, hobby := range hobbies {
			fmt.Printf("%v - Hobbie: %v\n", index, hobby)
		}
	}
}
