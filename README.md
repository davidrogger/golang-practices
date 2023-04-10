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


# Links uteis

- [Documentação Go](https://go.dev/doc/)
- [Go online Playground](https://go.dev/play/)
- [Grupo de Estudos](https://github.com/crgimenes/estudos)
- []()

*All tools successfully installed. You are ready to Go. :)*
