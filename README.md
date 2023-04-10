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

# Links uteis

- [Documentação Go](https://go.dev/doc/)
- [Go online Playground](https://go.dev/play/)
- [Grupo de Estudos](https://github.com/crgimenes/estudos)
- []()

*All tools successfully installed. You are ready to Go. :)*
