/*
	- Partindo do código baixo, ordene os []user por idade e sobrenome
	- Os valores no campo Sayings devem ser ordenados também, e demonstrados de maneira esteticamente harmoniosa.
*/

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println("Normal:", users)

	sort.Slice(users, func(i, j int) bool {
		return users[i].Age < users[j].Age
	})

	fmt.Println("Age:", users)

	sort.Slice(users, func(i, j int) bool {
		return users[i].Last < users[j].Last
	})

	fmt.Println("Last:", users)

	for _, user := range users {
		sayings := user.Sayings
		sort.Slice(sayings, func(i, j int) bool {
			return len(sayings[i]) < len(sayings[j])
		})
	}

	for _, user := range users {
		fmt.Println("User:", user.First, user.Last, user.Age)
		for _, phrase := range user.Sayings {
			fmt.Println(phrase)
		}
	}
}
