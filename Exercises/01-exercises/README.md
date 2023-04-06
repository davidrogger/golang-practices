# Nível 1

## Exercício #1

Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".

1. 42
2. "James Bond"
3. true

Agora demonstre os valores nestas variáveis, com:

- Uma única declaração print
- Múltiplas declarações print

## Exercício #2

Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos para estas variáveis:
1. Identificador "x" deverá ter tipo int
2. Identificador "y" deverá ter tipo string
3. Identificador "z" deverá ter tipo bool

Na função main:
- Demonstre os valores de cada identificador
- O compilador atribuiu valores para essas variáveis. Como esses valores se chamam? R: *valor zero*/*zero value*

# Exercício #3

Utilizando a solução do exercício anterior:
1. Em package-level scope, atribua os seguintes valores às variáveis:
2. para "x" atribua 42
3. para "y" atribua "James Bond"
4. para "z" atribua true

- Na função main:
  - Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
  - Demonstre a variável "s".

# Exercício #4

1. Crie um tipo. O tipo subjacente deve ser int.
2. Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.

- Na função main:
  - Demonstre o valor da variável "x"
  - Demonstre o tipo da variável "x"
  - Atribua 42 à variável "x" utilizando o operador "="
  - Demonstre o valor da variável "x"

> Para os aventureiros: https://golang.org/ref/spec#Types

# Exercício #5

1. Utilizando a solução do exercício anterior:
2. Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
- Na função main:
  - Isto já deve estar feito:
  - Demonstre o valor da variável "x"
  - Demonstre o tipo da variável "x"
  - Atribua 42 à variável "x" utilizando o operador "="
  - Demonstre o valor da variável "x"
- Agora faça tambem:
  1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
  2. Demonstre o valor de "y"
  3. Demonstre o tipo de "y"
  