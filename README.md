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

# Links uteis

- [Documentação Go](https://go.dev/doc/)
- [Go online Playground](https://go.dev/play/)
- [Grupo de Estudos](https://github.com/crgimenes/estudos)
- []()

*All tools successfully installed. You are ready to Go. :)*
