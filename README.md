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

# Links uteis

- [Documentação Go](https://go.dev/doc/)
- [Go online Playground](https://go.dev/play/)
- [Grupo de Estudos](https://github.com/crgimenes/estudos)
- []()

*All tools successfully installed. You are ready to Go. :)*
