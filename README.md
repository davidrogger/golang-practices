# Repositório para anotações e hands-on Golang

## [Curso AprendaGo](https://github.com/vkorbes/aprendago) por [Ellen Korbes](https://github.com/vkorbes)

[![IMAGE_ALT](https://img.youtube.com/vi/WiGU_ZB-u0w/0.jpg)](https://www.youtube.com/watch?v=WiGU_ZB-u0w)

# Golang

Resolvendo o problema de:

- Linguagens com alta performance e complexidade alta
- Linguaguens com complexidade baixa e baixa performance

Com base no C++, sendo uma linguagem com bastante performance, porém complexa e seu tempo de compilação alto, foi então iniciado a criação de uma nova linguagem, que resolvesse esses problemas.
Além de melhor uso para processadores de multi núcloes da época.

Desenvolvido por:
- Ken Thompson(Unix)
- Rob Pike (UTF 8)
- Robert Griesemer(V8)

É uma linguagem:
- Eficiente
- Ótima biblioteca nativa
- Multiplataforma
- Garbage collection (lightning fast)
- Cross-compile
- Tipagem forte e estática
- Fácil de aprender

# Orientação a objetos em Go

O conceito de interface em Go oferece uma abordagem diferente que acredita ser mais fácil de usar de uma forma mais geral. Há também formas de embutir tipos com outros tipos, providos de algo análogo, mais não idêntico, subclasses. Os métodos em go são mais gerais do que em C+ ou Java, eles podem ser definidos para qualquer tipo de informação, até mesmo tipos integrados assim como plain, unboxed, integers. Eles não são restritos como estruturas (classes).

# Operador Curto de declaração

Para declarar uma variavel em go, pode-se usar o `:=` conhecido como gopher, ele realizar a tipagem por inferência e só pode ser usado em blocks de códigos.

Exemplo:

```
package main

import "fmt"

func main() {
  a := "exemplo de operador curto"
  fmt.Println(a)
}
```

# Tipos

## Tipo booleano

- É um tipo binário, que só contem dois valores, true ou false.
- Usados em lógicas condicionais, declarações switch, if, fluxo de controle, etc.

## Tipos Númericos

Tipos númericos são divididos em duas categorias os números inteiros(int) e os números com ponto flutuantes(float64) que são os números com ponto(números racionais ou reais).
Para mistura-los na declaração é necessário usar conversão.
## Tipo string

São sequencias de bytes Imutáveis, slice of bytes, uma cadeia de caracteres.

## Declaração de tipos

Ao declarar um  variavel, a tipagem em go, ocorre por inferência caso não seja uma declaração explicita.
Tipagens em go, são estáticas, após uma variavel ter seu tipo declarado, ele não pode ser alterado.
Para declarar tipos de forma explicita:

1. `var nome string = "isso é uma string"`
2. `var numero int = 1`

Os tipos nativos como bloco básicos são; int, string e bool.
Os tipos compostos, são tipos compostos pelos primitivos; slice, array, struct, map.
A criação de estrutura de tipos composto chama-se composição.

# Valor zero

Valor zero é o valor inicial de uma declaração, quando não atribuido um valor;

```
var numeroInteiro int
var numeroFlutuant float32
var frase string
var resposta bool
```

Quando acessando essas variaveis, seus valores serão o valor zero, que por padrão;

1. int = 0
1. float32 = 0
2. string = É uma string vazia ""
3. bool = false
4. O valor zero de funções, pointers, interface, slices, channels, maps são nil (vazio)

# Criando tipos

"Type is life", grande parte dos aspectos mais avançados de Go dependem quase que exclusivamente de tipos.

```
type hotdog int
var a hotdog = 1
```

Na criação desse tipo hotdog para exemplo, foi definido seu subconjunto como inteiro, logo ele somente aceitará a declaração com elementos do tipo inteiro. Porém ele não pode receber o valor de outra variavel que é do tipo int;
```
var b int = 1
var a = b
```

Ele não aceita esse tipo de declaração, pois ele espera o tipo hotdog, não inteiro por mais que o hotdog seja int.

# Conversão, não coerção

Em Go não se é usado os termos, coerção ou cast e sim Conversão, usando o exemplo anterior;
```
var b int = 1
var a = hotdog(b)
```

É possivel converter o valor de b para hotdog, ja que ele é um int.

# Sistemas númericos

- Base10: decimal, 0~9
- Base2: binário, 0~1
- Base16: hexadecimal, 0~f

# Constantes

Assim como as var, podem ter seus tipos declarados ou recebem seu tipo por inferência, mas somente quando a constante é usada, com var seu tipo é inferido no momento da sua declaração.

```
const hello = "Hello"
const world string = "World"
```

# Iota

Em uma declaração de constantes, o identificar iota representa uma sequencia de números;

```
package main

import "fmt"

func main() {
  const (
    first = iota
    second
    third
    fourth
  )
  fmt.Println(first, second, third, fourth) // 0, 1, 2, 3

}
```

# Deslocamento de bits

Quando deslocamos digitos binários para a esquerda ou direita;

```
a := 100 // 1100100
b := a << 1

fmt.Println(b) // 200 1100100<<0
```

Ação acima foi deslocando os valores em binário para esquerda 1101100 e adicionando um 0 a direita.

# Fluxo de controle

É como os computadores lêem o código que foi implementado, assim como alguns países, a leitura, acontece da esquerda para direita de cima para baixo, em alguns lugares, a leitura pode ser de cima para baixo da direita para esquerda.
Os computadores lêem os códigos de cima para baixo da esquerda para direita.
Além do controle sequencial, há duas declarações que podem afetar como o computador lê o código:

- Fluxo de controle de repetição, o computador repete a leitura de um mesmo código de uma maneira específica, conhecido como fluxo de controle iterativo.
- Fluxo de controle condicional, ou controle de seleção. O computador encontra uma condição e por meio de uma declação if ou switch, toma um curso ou outro com base na condição.

Totalizando três tipos de fluxo: sequencial, repetição e condicional.

# FOR

Em go, não temos o loop while, somente o FOR para criar loops;

Sintaxe;

```
package main

import "fmt"

func main() {
	for i := 0; i < 10; i += 1 {
		fmt.Println(i)
	}
}
```

Imprimindo os valores de 0 à 9

## continue e break

Existem duas formas de parar uma iteração durante um for, usando break a iteração é parada por completo, parando toda execução, agora usando continue, ele apenas para aquela iteração/ciclo do loop.

# Switch

Usado para declaração de condicionais;

Possiveis sintaxes;

Usando condições de comparação que em caso de verdadeiras execução a linha de case.
```
x := 55
switch true {
  case x < 55:
    fmt.Println("Menor que 55")
  case x > 55:
    fmt.Println("Maior que 55")
  case x == 55:
    fmt.Println("É igual a 55")
}
```

Condições de igualdade, caso a variavel declara seja igual a condição do case.
```
x := "Jonas"
switch x {
  case "maria":
    fmt.Println("É a Maria")
  case "Juka":
    fmt.Println("É o Juka")
  case "Jonas":
    fmt.Println("É o Jonas")
}
```

Multiplas opções nos cases para execução.

```
x := "Jonas"
switch x {
  case "maria", "Juka":
    fmt.Println("Membro da Equipe Alpha")
  case "Juka", "Jonas":
    fmt.Println("Membro da Equipe Bravo")
  default:
    fmt.Println("Não pertence a nenhuma equipe")
}
```

Switch do go, possui fallthrough que após encontrar o bloco apresenta o proximo bloco também, após o bloco encontrado.

```
    case i < 10:
        fmt.Println("i is less than 10")
        fallthrough
    case i < 50:
        fmt.Println("i is less than 50")
    case i < 100:
        fmt.Println("i is less than 100")
    }
```

Neste caso quando entrar no primeiro case ele mostra o conteudo do case mais o conteudo do proximo case.

# Array

É usado em go quando é necessário uma melhor performance comparado ao slice, pois nele é necessário, definir além do tipo um tamanho fixo, no momento da declaração, tirando a possibilidade de adicionar mais elementos além do limite determinado;

```
var x [5]int // [0, 0, 0, 0, 0]

x[0] = 10 //  [10, 0, 0, 0, 0]
```

# Slices

É o mais usado quando lidando com lista por ser mais flexivel com relação ao seu tamanho. Sendo possivel adicionar mais elementos usando o append


```
	array := [5]int{5, 3, 5} // [5, 3, 5, 0, 0]
	slice := []int{5, 3, 5} // [5, 3, 5]
  
	slice = append(slice, 10) // [5, 3, 5, 10]
```

# Range

Usado em lopp for, quando percorrendo uma lista (array, slice);

```
	fruits := []string{"banana", "melon", "strawberry"}

	for index, fruit := range fruits {
		fmt.Printf("A fruta na posição %v é uma %v\n", index, fruit)
	}
```

# Slice of Slices

Criando outro slice, apartir de um slice, removendo informações;
Deve-se indicar em qual indice a nova lista começa e, qual é a ultima posição da nova lista (o ultimo elemento não fará parte da nova lista)

```
	fruits := []string{"banana", "melon", "strawberry"}

	fruit := fruits[1:2] // ["melon"]

```

Também é possivel usar o append, para unir dois slices:

```
	fruits1 := []string{"banana", "melon", "strawberry"}
	fruits2 := []string{"avocado", "pinnaple", "orange"}

	fruits3 := append(fruits1[:1], fruits2[1:2]...)

	fmt.Println(fruits3) // ["banana", "pinnale"]
```

# Slice: make

Slices são feitos de arrays, são dinâmicas, e podem mudar de tamanho, mas sempre que isso acontece, um novo array é criado gerando um custo computacional adicionar, para otimizar isso, é usado o make, onde é determinado, o tipo, seu tamanho e capacidade;

```
	fruits := make([]int, 5, 10)
	fmt.Println(fruits) // [0, 0, 0, 0, 0]
  	fruits = append(fruits, 6, 7, 8, 9, 10, 11, 12)

	fmt.Println(fruits) // [0, 0, 0, 0, 0, 6, 7, 8, 9, 10, 11, 12]
```

No exemplo é determinado que é um slice de numeros inteiros []int, com valor inicial de 5, assim como o array, ja deixa reservado aquele espaço, seguindo da capacidade máxima de elementos que o slice pode armazenar sem alocar mais espaço na memória e é possivel realizar o append de novos elementos, sempre que ele atinge o tamanho limite, ele dobra o tamanho, para continuar alimentando de forma mais performática, mas esse realocamento do array, pode ser um operação custosa em termos de tempo e recurso também.

# Slices multidimensionais

É a composição de slices com slices de um determinado tipo, assim como uma matriz;

```
	matriz := [][]string{
		[]string{"linha1-a", "linha1-b", "linha1-c"},
		[]string{"linha2-a", "linha2-b", "linha2-c"},
		[]string{"linha3-a", "linha3-b", "linha3-c"},
	}

	fmt.Println(matrix[1][1]) // linha2-b
```

# Reutilizando slice

Quando usando das fatias de fatias em slice, deve-se tomar cuidado na seguinte cituação;

```
	primeiroslice := []int{1, 2, 3, 4, 5}
	
	fmt.Println(primeiroslice)
	
	segundoslice := append(primeiroslice[:2], primeiroslice[4:]...)

	fmt.Println(segundoslice)

	fmt.Println(primeiroslice) // [1, 2, 5, 4, 5]
```

Ao tentar reaproveitar o array anterior, por de baixo dos panos, a linguagem faz um embaralhamento da lista, fazendo a terceira posição se tornar o número 5.


# Maps

É um composição de tipos que usa de chave e valor;

```
	users := map[string]int{}
	users["jonas"] = 10
	fmt.Println(users) // map[jonas:10]
```

Chaves jonas possui o valor 10. Sendo possivel atualizar o valor, da mesma forma que foi inserido o valor de jonas. E também para consultar o valor especifico da chave, basta acessar o map com colchetes e a chave que deseja o valor.

Quando acessando um valor inexistente, o Go retorna o valor 0, como fica sublime se é um valor ou se não existe um valor, existe o método comma ok, que consiste em um segundo valor quando acessando o valor de um mapa;

```
user, ok := users["mary"]
```

A variavel ok, caso não existe o usuário que deseja consultar no map de users, deve ser false.

## Map: range

Assim como em uma lista de array ou slice, o uso do range em um Map, é para percorrer as chaves e valores contidos no mapa;

```
	usersMap := map[string]string{
		"jonas": "always here",
		"mad":   "new here",
		"modiz": "dirty",
		"biel":  "is biel",
	}

	for key, value := range usersMap {
		fmt.Println(key, value)
	}
```

Range em Map, returna dois valores, sendo o primeiro o nome da chave, e o segundo o valor da chave.

## Map: Delete

Para deletar uma chave de um map, basta usando o metodo `delete`;

```
delete(usersMap, "modiz")
```

Dessa forma a chave e todo conteudo será eliminado do map usersMap.

# Structs

É um tipo de da dado composto que permite armazenar valores de tipos diferentes;

```
type user struct {
	name 			string
	surname		string
	age				int
}

jonas := user {
	name: "Jonas",
	surname: "Doe",
	age: 33,
}
```

# Struct: nested

É quando temos um structer dentro de outro;

```
type job struct {
	user	user
	task	string
}

type user struct {
	name 			string
	surname		string
	age				int
}

helpdesk := job{
	user: user{
		name: "Jonas",
		surname: "Doe",
		age: 33,
	},
	task: "Remote Acess"
}
```
Para acessar os valores de forma especifica, pode-se usar dot notation;

```
helpdesk.name // jonas
```
Nota-se que não é necessário acessar o user.name, pois ele tem um funcionamento chamado embeded field, onde você pode apontar diretamente para chave interna de um elemento, desde que ele seja único.

Pode-se declarar a de forma menos esplicita a estutura quando criando ela;

```
helpdesk := job{user{"jonas", "doe", 33}, "Remote acess"}
```

# Structers anonimos

São estruturas temporarias criadas apenas no momento da declaração;

```
u := struct {
	name:	string
	age:	int
}{ "jonas", 33 }
```

# Funções

É uma forma de segmentar funcionalidades de abstrair comportamentes de forma mais isolada, sintaxe;

```
func (receiver) nome(argumento tipo) (tipo de retorno) { codigo }
```

Declaração de uma função em Go começa com a palavra reservada `func`, receiver terá um tópico mais especifico sobre, e ele pode ser omitido em sua declaração, seguindo de um nome, com os argumentos dos parametros que serão necessários para função, e a declaração do tipo do dado que a função vai retornar, com o código propriamente digo embolta de chaves {}.

Os argumentos em Go, podem ser variadíco, que quer dizer, que seu argumento com a representação de 3 pontos "...", pode receber uma quantidade indeterminada de argumentos, consequentemente em forma de slice, nota que caso tenha mais de um argumento,  variadíco, deve ser o ultimo argumento declarado na função.

# Defer

É uma declaração definindo que ele determinado conteudo deve ser executado somente no final da função usada.
Quando usando multiplos defer, sempre o primeiro defer vai ser o ultimo, seguindo a ordem como se fosse uma pilha de roupas.

# Métodos

A declaração de um método em uma função é usada para decorar um tipo especificado em seu receiver;

```
type user struct {
	name string
}

func (u user) printName() {
	fmt.Println(u.name)
}

jonas := user{ "jonas" }
jonas.printName()
```

Agora todo tipo user, possui o método printName, quando invocado.

# Interfaces

São coleções nomeadas ou assinaturas de métodos;
[Exemplo](https://gobyexample.com/interfaces);

Uma interface básica com formas geométricas, recebendo os métodos dos tipos:
```
package main 

import "fmt"

type geometry interface {
	area() float64
	perim() float64
}
```

Implementando os métodos nos tipos rect e circle;

```
// Dois tipos para forma geografica, retângulo, recebe as medidas de largura e altura
type rect struct {
	width, height float64
}

// e circulo apenas o raio.
type circle struct {
	radius float64
}

// Método atrelado ao tipo rect onde retorna a area do retângulo.
func (r rect) area() float64 {
	return r.width * r.height
}

// Método atrelado ao retangulo, onde calula o perimetro do retângulo
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}


// Método atrelado ao tipo do circulo, onde retorna a area do circulo
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Método atrelado ao tipo do circulo, onde retorna o perimetro do circulo
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// função measure que recebe um argumento do tipo geometry que é a interface
// A interface geometre possui os métodos area e perim, então nessa função
// é imprimido o resultado do tipo geometry area e perim.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{ width: 3, height: 4 }
	c := circle{ radius: 5 }

	measure(r)
	measure(c)
}
```

Dessa forma fica mais simples criar funcionalidades que sofrem polimorfismo com base em seu tipo usando Go, por de baixo dos panos, é como se tivesse uma condicional apontando que determinado tipo, tem determino comportamento.

# Funções anônimas

Normalmente um função descartavel e temporaria;

```
func(x int) {
	fmt.Println(x, "vezes 15 é:")
	fmt.Println(x * 15)
}(10)
// 150
```

Nota-se após invocar o func, não foi declarado seu nome, e após fechar as chaves {} foi passado o parametro para o argumento da função.

# Funções como expressão

Em go é possivel também passar a função como uma expressão para variaveis;

Funções normais;

```
func teste(x int) {
	fmt.Println(x, "vezes 15 é:")
	fmt.Println(x * 15)
}

func main() {
	x := teste

	x(10)
}
```

Ou também anônimas;

```
y := func(x int) {
	fmt.Println(x, "vezes 15 é:")
	fmt.Println(x * 15)
}
y(10) // 150
```

# Retornando uma função por outra função

É possivel também retornar outra função quando invocando uma função;

```
func main() {
	x := example()
	y := x(3)

	fmt.Println(y)
}

func example() func(int) int {
	return func(i int) int {
		return i * 10
	}
}
```

Também é possivel passar uma segunda chamada na função, para ja executar a função retornada, e coletar o resultado;

```
x := example()(3) // 30
```

# Callbacks

É quando é passada uma função como um parametro para a função;

```
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
```

# Closure

É o uso de um escopo por contexto;

```
func main() {

	a := i()

	fmt.Println(a()) // 1 
	fmt.Println(a()) // 2
	fmt.Println(a()) // 3

	b := i()

	fmt.Println(b()) // 1
	fmt.Println(b()) // 2
	fmt.Println(b()) // 3
}

func i() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
```

Neste exemplo é guardado o contexto de x da função, onde ao invarmos ela sequencialmente sua contagem progride, porém quando criarmos uma nova variavel, é criado um nova closure relativa a nova variavel.

# Ponteiros

Todos valores ficam armazenados em memória, e toda localização na memória possui um endereço, o ponteiro é a referencia a esse endereço;

Para visualizar a localização no endereço da memória de uma variável é usado o E comercial `&`, `&x` dessa forma seria apresentado o endereço da memória da variavel x, agora usando o ponteiro `*` ele mostra o valor que está dentro da variavel, `*x`.

O uso do ponteiro ajuda na performance ao executar alguma funcionalidade, pois em Go tudo é pass-by-value, que significa que quando passamos uma variavel para um função como parametro, significa que ele cria uma cópia daquela variavel para tratar ela, gerando um consumo de performance, ja quando usamos o ponteiro indicando diretamente a localização na memória e para o uso do valor, ele deixa de criar uma copia usando o valor dentro do endereço indicado pelo ponteiro.

# JSON

JavaScript Object Notation, é um formato de dados leve e utilizado para troca de informações entre sistemas.

Para converter uma estrutura de dados, para json usando o go, podemos usar o biblioteca Marshal;

```
import "encoding/json"
type pessoa struct {
	Nome: string
	Sobrenome: string
	idade: int
}

func main() {
	jonas := pessoa{"jonas", "doe", 33}

	j, err := json.Marshal(jonas)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(j))

}
```

Neste exemplo, um ponto importante usando o json.Marshal é a necessidade do uso das chaves da estrutura em letra maiuscula, pois não é aceito para exportação as chaves minusculas.
Também em go deve-se tratar os erros, de forma individual, após rodar o Marshal, devemos verificar se há algum erro, que é realizado com o if.

Agora para realizar o contrário, receber um json e converter para uma estrutura de código;

```
jsonLine := []byte(`{"jonas", "doe", 33}`)

var convertedData pessoa

err := json.Unmarshal(jsonLine, &convertedData)
if err != nil {
	fmt.Println("error:", err)
}

fmt.Println(convertedData.Nome)
```

jsonLine foi criado apenas para simular uma variavel que contem um json, simulando seu formato de byte, criando uma variavel com a tipagem desejada, usamos o Unmarshal, passando o primeiro parametro como o arquivo de conversão seguindo da localização na memória da variavel onde deve ser salva a conversão seguindo de um tratamento de error.

# JSON: Tags

Também quando estamos recebendo informações em json é possivel adicionarmos tags, de referencia para os campos do tipo;

```
type pessoa struct {
	Nome: 			string	`json:"Nome"`
	Sobrenome:	string	`json:"SegundoNome"`
	idade: 			int			`json:"idade"`
}
```
Arquivo recebido em json, tem o campo sobrenome, como SegundoNome, porém queremos que seja SobreNome.(apenas exemplo)

# Sort

Usado para ordenar slices, ordenando strings;

```
import "sort"

func main() {
	names := []string{"Jonas", "juka", "ananias", "ze", "bety"}

	fmt.Println(names) // Jonas juka ananias ze bety

	sort.Strings(names)

	fmt.Println(names) // Jonas ananias bety juka ze
}
```

Segue o mesmo padrão, para ints e floats;

```
	numbers := []int{1000, 10, 3000, 55, 3, 22}

	fmt.Println(numbers)

	sort.Ints(numbers)

	fmt.Println(numbers)
```

# Custom sort

Para usar o metodo que ordena algum slice de um tipo personalizado, como ordenando pelo nome, ou alguma chave especifica, deve-se definir um métodos, que retorne os tipos Len, less e Swap, para o funcionamento do Sort.

```

import (
	"fmt"
	"sort"
)

type User struct {
	name   string
	age    int
	salary float64
}

type orderByName []User

func (o orderByName) Len() int           { return len(o) } // Sort precisa saber o tamanho da lista que está sendo organizada.
func (o orderByName) Less(i, j int) bool { return o[i].name < o[j].name }, verifica se o elemento anterior é maior que o próximo elemento na lista.
func (o orderByName) Swap(i, j int)      { o[i], o[j] = o[j], o[i] } // Troca a posição do primeiro elemento com o segundo.

func main() {
	users := []User{
		User{"Denis", 20, 1000},
		User{"Albert", 25, 1010},
		User{"Bonas", 33, 2500},
	}

	fmt.Println(users)

	sort.Sort(orderByName(users))

	fmt.Println(users)
}
```

# Sort: Slice

Usando o metodo Slice de sorte, sem precisar criar novos métodos;

```
import (
	"fmt"
	"sort"
)

func main() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age:", people)
}
```

# bcrypt

É um modulo, usado para encriptar informações, normalmente usado em senhas;

```
package main

import "fmt"
import "golang.org/x/crypto/bcrypt"

func main() {
	senha := "20julho1980"
	senhaerrada := "20julho1990"

	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sb))

	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senha)))
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senhaerrada)))

}
```

Para gerar a senha encryptada, usando a função GenerateFromPassword deve-se passar a senha em forma de slice de bytes, definindo o custo(nível de criptografia) que será usado parar gerar aquela senha.
Para verificar a autenticidade da senha é usado a função CompareHasAndPassword, recebendo a senha encryptada com a senha em forma de slice the bytes.

# Concorrência e Paralelismo

Concorrencia é um conceito que se refere à capacidade de um programa executar várias tarefas ao mesmo tempo. Independente da forma como essas tarefas são executadas. Sendo várias tarefas em um único processador, por meio do compartilhamento de tempo (time-sharing), ou a execução de várias tarefas em diferentes processadores ou núcles de CPU por meio do paralelismo.

O paralelismo é uma forma específica de concorrência envolvendo a execução simultânea de várias tarefas em diferentes processadores ou núcleos de CPU ao mesmo tempo.

# Thread

É uma sequência de instruções executadas pelo sistema operacional de forma independente, permitindo que programas realizem várias tarefas simultaneamente, dividindo a execução em sequências menores, melhorando o desempenho e a capacidade de resposta de programas, porém isso exige um certo cuidado pois existe um problema de sincronização e condições de corrida, onde duas ou mais threads tentam acessar ou modificar a mesma variável ou recurso ao mesmo tempo.

# Go routines

É usado para definir que uma funcionalidade irá trabalhar de forma assyncrona ao demais itens do código, sendo paralela ao código;

```
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup

func main() {
	wg.Add(2) // define a quantidade de go routines usadas na função.

	go func1()
	go func2()

	wg.Wait() // Faz com que a função main espere pelas rotinas terminaram para encerrar a função.
}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1:", i)
		time.Sleep(20)
	}
	wg.done() // Conclui a rotina
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func2:", i)
		time.Sleep(20)
	}
	wg.done()
}
```

Quando usando go routine é necessário definir quantas rotinas estão para sendo executadas e cada rotina, deve ter um momento de conclusão para então encerrar sua execução. Caso contrário sem definir a quantidade e seu encerramento, a função principal, encerrará sem esperar que as rotinas terminar suas atividades.


# Condição de corrida

É quando trabalhando com paralelismo é acessada uma variavel por diferentes funções em parelalo e ela alteram essa variavel, mas essa variavel acaba se tornando volátil e incistente, pois não há um controle sobre o valor que as variaveis vão estar acessando e alterando;

```
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	contador := 0
	totaldegoroutines := 1000

	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)

}
```

Nessa funcionalidade foram determinada 1000 funções que realizam acesso ao contador, e alteram seu valor de forna incremental, o resultado devido a condição de corrida, não é de 1000, mas sim um valor sem nenhum tipo de controle.

# Mutex

Mutual exclusion, é uma técnica usada na programação concorrente para evitar que várias threads acessem um recurso compartilhado simultaneamente. Ele tem o objetivo de garantir que apenas uma funcionalidade acessa o recurso compartilhado (variavel) em um determinado momento, evitando o race condition;

```
package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}

```

Usando o Lock do mutex, o resultado final do contador será 1000.

# Atomic

Fornece uma maneira segura e eficiente de executar operações atômicas em variáveis compartilhadas entre goroutines. As operações são executadas de forma indivisível, ela ´executada como uma única unidade sem interrupção de outras operações que ocorrem simultaneamente.
As operações atômica incluem leitura, escrita, incremento, decremento, swap e outras;

```
package main

import (
	"fmt"
	"sync/atomic"
)

var counter int32

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&counter, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}

```

No exemplo é usado o método AddInt32, para incrementar o contador. A função recebe um ponteiro para a varáivel que deve ser incrementada e o valor a ser adicionado a essa variável, passando o endereço de memória do counter (`&counter`),  para a função AddInt32 e o valor a ser incrementado.
Operações atômicas são geralmente mais rápidas e mais leves do que o uso de mutex para proteger variáveis compartilhadas, por evitar a necessidade bloquear e desbloquear para acesso da variável. Porém ela não consegue garantir proteção contra race condition em casos mais complexos.

# Channel

Canais um dos diferenciais do Golang, para fazer sincronização de código, eles nos permitem transmitir valores entre goroutines, coordenando, sincronizando e orquestrando, dessa forma é possivel fazer do jeito certo a sincronização de código concorrente.


```
	c := make(chan int)

	c <- 42
	fmt.Println(c)
```

>fatal error: all goroutines are asleep - deadlock!

Quando usando canal, é necessário "trabalhar" de maneira concorrente, nesse caso, a linha em que é inserido o numero 42 ao canal, ele fica sem resposta por não ser uma goroutine, em uma mesma goroute não é possivel realizar duas atividades, ou ela retira ou adiciona.

```
	c := make(chan int)

	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
```

Dessa maneira funciona, pois (main é uma goroutine também), temos duas goroutine, uma adicionando ao canal o 42 na função anônima e outra retirando o valor do canal quando usado o print.
Repetindo, precisamos de uma goroutine para cada ação sobre o canal, nesse caso, uma está adicionando (função anônima) e a outra está retirando (função main com o print)

## Channel: Buffer

Pode-se definir em um canal usando o buffer que não é necessário ao mesmo tempo que alguém está retirando informação ela seja inserida;

```
c := make(chan int, 1) // declarando o numero 1, não é necessário que ao mesmo tempo que é retirado deve-se inserir
c <- 42
fmt.Println(<-c)
```

Porém o uso de buffer não é uma prática recomendada, a forma mais eficiente é sempre usar o receive and sender sincronizado.

# Canais Direcionais

Canais podem ser direcionais, eles podem ser declarados como canais gerais, ou com uma determinada função ja pré-definida, sendo, um receiver que só recebem informação ou sender só enviam informação;

> send chan<-\
> receive <-chan

```
	c := make(chan int)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// general to specific converts
	fmt.Println("-----")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))
```

É possivel converter um canal, geral, para um especifico, mas não um especifico para um geral.

## Channels: Range and Close

Quando iterando um canal é necessário definir um término para ela, caso contrário teremos um erro de deadlock

```
func main() {
	c := make(chan int)

	go feedChannel(20, c)
	printChannel(c)
}

func feedChannel(qt int, c chan<- int) {
	for i := 0; i < qt; i++ {
		c <- i
	}
	close(c)
}

func printChannel(c <-chan int) {
	for value := range c {
		fmt.Println("Channel Value:", value)
	}
}
```

# Channel: Select

Ele é como um switch, onde ele procura por uma condicional, quando encontrado mais que um acertivo, ele escolhe de forma aleatória entre as alternativas.

```
func main() {
	a := make(chan int)
	b := make(chan int)
	qt := 100

	go feedChannel(qt, a)
	go feedChannel(qt, b)

	printChannels(qt, a, b)
}

func feedChannel(qt int, c chan<- int) {
	for i := 0; i< qt; i++ {
		c <- i
	}
	close(c)
}

func printChannels(qt int, a <-chan int, b <-chan int) {
	for i := 0; i < qt; i++ {
		select {
			case v := <-a:
				fmt.Println("Canal A:", v)
			case v:= <-b:
				fmt.Println("Canal B:", v)
		}
	}
}

```

Quem pertencer ao canal A, terá o print com a frase Canal A e o B com B.\

Dentro de um select é possivel também enviar ou receber canais;

```
func main() {
	c := make(chan int)
	q := make(chan int)
	go receiveQuit(c, q)
	sendChannel(c, q)
}

func receiveQuit(channel chan int, quit chan int) {
	for i := 0; i < 50; i++ {
		fmt.Println("Recebido", <- channel)
	}
	quit <- 0
}

func sendChannel(channel chan int, quit chan int) {
	counter := 1
	for {
		select {
			case channel <- counter:
				counter++
			case <- quit:
				return
		}
	}
}
```

Sem usar o close, ele printa até o 50, e quando chega na ultima faixa do canal, ele para.

## Channel: Comma ok

Para verificar se o valor do canal é 0 ou não existe mesmo, pois como ele retorna um zero value;

```
canal := make(chan int)

go func() {
	canal <- 0
	close(canal)
}()

v, ok := <-canal

fmt.Println(v, ok) // 0, true

v, ok = <-canal

fmt.Println(v, ok) // 0, false
```

Primeira vez, ele retira o valor 0 do canal, e mostra na tela, como valor 0 e true, o useja existe o valor 0, e depois ele tira novamente o valor que não existe, sendo o value zero 0 porém é vazio false.

## Convergência

É o conceito de unir informações de dois canais em um.

```
func main() {
	even := make(chan int)
	odd := make(chan int)
	converge := make(chan int)

	go send(even, odd)
	go receive(even, odd, converge)

	for v := range converge{
		fmt.Println("Convergence Number:", v)
	}
}

func send(even, odd chan int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

func receive(even, odd, converge chan int) {
	var wge sync.WaitGroup
	wg.Add(1)

	go func() {
		for v:= range even {
			converge <- v
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for v:= range odd {
			converge <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(converge)
}

```

## Divergência

É o conceito contrário a convergência, onde tiramos informação de um canal, e quebramos em mais canais, com a finalidade de realizar uma concorrencia entre o trabalho necessário com a informação do canal.

```
func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)

	funcoes := 5

	go send(100, canal1)
	go other(funcoes, canal1, canal2)

	for v := range canal2 {
		fmt.Println(v)
	}
}

func send(n int, canal chan int) {
	for i := 0; i < n; i++ {
		canal <- i
	}
	close(canal)
}

func other(funcoes int, canal1, canal2 chan int) {
	var wg sync.WaithGroup

	for i := 0; i < funcoes; i++ {
		wg.Add(1)
		go func() {
			for v: range canal1 {
				canal2 <- workLoad(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(canal2)
}

func workLoad(n int) {
	time.Sleep(time.Millisecon * 1000)
	return n
}

```

Nesse exemplo, é determinado para 5 go routine trabalharem de forma concorrente, definida na variavel funcoes.

# [Context](https://pkg.go.dev/context)

É usado para realizar comunicação entregue go routines, para encerramento deles, caso uma operação não seja necessário.

- Destaques:
    - ctx := context.Background
    - ctx, cancel = context.WithCancel(context.Background)
    - goroutine: select case ←ctx.Done(): return; default: continua trabalhando.
    - check ctx.Err();
    - Tambem tem WithDeadline/Timeout
- Exemplos (Todd):
    - Analisando:
        - Background: https://play.golang.org/p/cByXyrxXUf 
        - WithCancel: https://play.golang.org/p/XOknf0aSpx
        - Função Cancel: https://play.golang.org/p/UzQxxhn_fm 
    - Exemplos práticos:
        - func WithCancel: https://play.golang.org/p/Lmbyn7bO7e
        - func WithCancel: https://play.golang.org/p/wvGmvMzIMW 
        - func WithDeadline: https://play.golang.org/p/Q6mVdQqYTt 
        - func WithTimeout: https://play.golang.org/p/OuES9sP_yX 
        - func WithValue: https://play.golang.org/p/8JDCGk1K4P

# Tratamento de erros

No desenvolvimetno de go foi definido que não usar exceções seria uma forma mais organizada de tratar os erros.\
Por tanto é interessante criar o hábito de lidar com os erros imediantamente.

## Verificando erros

Normalmente uma função em go, retorna o resultado ou seu erro, por isso o tratamento é realizando recebendo dois valores;

```
v, err := os.Open("no-file.txt")
if err != nil {
	fmt.Println("err happend", err)
}
```

Nessa funcionalidade, estamos abrindo um arquivo, e caso ele não exista, ou se enquadre dentro de alguma regra de erro da função, ele deve retornar o erro, e caso tudo ocorra conforme deve, ele retorn nil, que representa vazio/nada, então se ele for diferente de vazio, ocorreu um erro, então tratamos esse erro.

>- Recomendação: use log.
>- Código: 
>    - 1. fmt.Println
>    - 2. log.Println
>    - 3. log.SetOutput
>    - 4. log.Fatalln
>    - 5. log.Panicln
>    - 6. panic
>- panic: http://godoc.org/builtin#panic

# Recouver

```
func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
```

É demonstrado o funcionamento do panick, que quando acionado, realiza todos defers da função, ele não para a função completamente. E é usado o recouver para lidar com o panic.

- https://go.dev/blog/defer-panic-and-recover
- https://pkg.go.dev/builtin#recover

# Erros customizados

- Para que nossas funções retornem erros customizados, podemos utilizar:
    - return errors.New()
    - return fmt.Errorf() ← tem um errors.New() embutido, olha na fonte!
    - https://golang.org/pkg/builtin/#error
- “Error values in Go aren’t special, they are just values like any other, and so you have the entire language at your disposal.” - Rob Pike
- Código: 
    - 1. errors.New
		```
		package main

		import (
			"errors"
			"log"
		)

		func main() {
			_, err := sqrt(-10)
			if err != nil {
				log.Fatalln(err)
			}
		}

		func sqrt(f float64) (float64, error) {
			if f < 0 {
				return 0, errors.New("norgate math: square root of negative number")
			}
			return 42, nil
		}
		```
    - 2. var errors.New
		```
		package main

		import (
			"errors"
			"fmt"
			"log"
		)

		var ErrNorgateMath = errors.New("norgate math: square root of negative number")

		func main() {
			fmt.Printf("%T\n", ErrNorgateMath)
			_, err := sqrt(-10)
			if err != nil {
				log.Fatalln(err)
			}
		}

		func sqrt(f float64) (float64, error) {
			if f < 0 {
				return 0, ErrNorgateMath
			}
			return 42, nil
		}

		// see use of errors.New in standard library:
		// http://golang.org/src/pkg/bufio/bufio.go
		// http://golang.org/src/pkg/io/io.go
		```
    - 3. fmt.Errorf
		```
		package main

		import (
			"fmt"
			"log"
		)

		func main() {
			_, err := sqrt(-10.23)
			if err != nil {
				log.Fatalln(err)
			}
		}

		func sqrt(f float64) (float64, error) {
			if f < 0 {
				return 0, fmt.Errorf("norgate math again: square root of negative number: %v", f)
			}
			return 42, nil
		}
		```
    - 4. var fmt.Errorf
		```
		package main

		import (
			"fmt"
			"log"
		)

		func main() {
			_, err := sqrt(-10.23)
			if err != nil {
				log.Fatalln(err)
			}
		}

		func sqrt(f float64) (float64, error) {
			if f < 0 {
				ErrNorgateMath := fmt.Errorf("norgate math again: square root of negative number: %v", f)
				return 0, ErrNorgateMath
			}
			return 42, nil
		}
		```
    - 5. type + method = error interface
		```
		package main

		import (
			"fmt"
			"log"
		)

		type norgateMathError struct {
			lat  string
			long string
			err  error
		}

		func (n norgateMathError) Error() string {
			return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
		}

		func main() {
			_, err := sqrt(-10.23)
			if err != nil {
				log.Println(err)
			}
		}

		func sqrt(f float64) (float64, error) {
			if f < 0 {
				nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
				return 0, norgateMathError{"50.2289 N", "99.4656 W", nme}
			}
			return 42, nil
		}

		// see use of structs with error type in standard library:
		//
		// http://golang.org/pkg/net/#OpError
		// http://golang.org/src/pkg/net/dial.go
		// http://golang.org/src/pkg/net/net.go
		//
		// http://golang.org/src/pkg/encoding/json/decode.go
		```

# Documentações

## go doc

Usado para visualizar a documentação do go, no terminal de comando, basta digitar go doc [package] essa parte da func, seria o nome do pacote que deseja ler a documentação, ela é pouco usada, pois normalmente a visualização no terminal, é desajeitada.
## godoc

Ele segue a mesma linha que o go doc, mas com ele é possivel a visualização pelo navegador usando o comando: godoc -http=:PORT, podemos definir a porta que quisermos, desde que ela não esteja sendo usada por outra porta; `godoc -http=:8080`

# godoc.org

É um site para "deploy" de uma documentação que foi criada no github. Onde basta copiar o caminho do pacote no github.

# Escrevendo documentação

A documentação no go é acoplada ao código, evoluindo juntamente com ele, todo comentário realizado junto ao código gera a documentação.\
- Para documentar um tipo, uma variável, uma constante, ou um pacote, escreva um comentário imediatamente antes de sua declaração, sem linhas em branco
- Comece a frase com o nome do elemento. No caso de pacotes, a primeira linha aparece no "package list."
- Caso esteja escrevendo bastante documentação, utilize um arquivo doc.go. Exemplo: package fmt.
- A melhor parte dessa abordagem minimalista é que é super fácil de usar. Como resultado, muita coisa em Go, incluindo toda a standard library, já segue estas convenções.
- Outro exemplo: errors package.

> main.go
```
// Package escrevendo ilustra como escrever documentação.
package escrevendo

// x é um número inútil.
const X = 10

// Doc é um monte de coisa nenhuma.
func Doc() {
	fmt.Println("Essa função não faz nada.")
}

// doc2 começa com letra minúscula
func doc2() {
	fmt.Println("Essa função não faz nada.")
}

// Doc3 é a última!
func Doc3() {
	fmt.Println("Essa função não faz nada. X é:", x)
}
```

>doc.go
```
/*
	Package escrevendo ilustra como escrever documentação.
	É apenas exemplo.
*/
package escrevendo
```

Ao rodar godoc -http=:8080, será gerado a documentação para http.\

Nota, que as funções são escritas com a letra inicial maiuscula, para ser exportavel e visivel para documentação, a função com letra minusculo, só pode ser usada internamente no package pois ela não é exportável.


# Links uteis

- [Documentação Go](https://go.dev/doc/)
- [Go online Playground](https://go.dev/play/)
- [Grupo de Estudos](https://github.com/crgimenes/estudos)
- []()

*All tools successfully installed. You are ready to Go. :)*
