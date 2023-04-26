/*
	- Utilizando goroutines, crie uma programa incrementador:
		- Tenha uma variável com o valor da contagem
		- Crie um monte de goroutines, onde cada um deve:
			- Ler o valor do contador
			- Salvar este valor
			- Fazer yield da thread com runtime.Goshed()
			- Incrementar o valor salvo
			- Copiar o novo valor para a variável inicial
		- Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
		- Demonstre que há uma condição de corrida utilizando a flag -race
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	const routineAmount = 100
	wg.Add(routineAmount)

	for i := 1; i <= routineAmount; i++ {
		go func() {
			runtime.Gosched()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
}

// go run -race Exercises/09-exercises/09-exercise-03.go
