# Go lang mini-manual

[Go](https://go.dev/) é uma linguagem de programação de alto-nível (poderia-se dizer de "médio nível" pros padrões atuais) desenvolvida pela Google para eles tentarem se livrar do cpp no desenvolvimento de sua infraestrtura web, tem sintaxe parecida com C (possuindo poquíssmas palavras e conceitos). Pode ser interpretada, mas seu diferencial é sua compilação à código de máquina nativa, tornando-a potencialmente rápida tanto para executar como para escrever. Ainda, possui _garbage collector_, o que evita erros de _memory leak_ em troca de um pouco da velocidade. Ferramentas como o Docker e infraestruturas de nuvem são escritas em Go. Têm um ecossistema relativamente robusto de bibliotecas e imports da comunidade (não é nenhum Python ou Node/Bun, mas tem bastante coisa legal no [https://pkg.go.dev/](https://pkg.go.dev/))

> A pasta go-basic-examples tem alguns exemplos de código

## Iniciando

- `go mod init` ->  initializa um módulo go de nome domínio/módulo
- `go mod tidy` -> busca dependências e outros módulos
- `go mod run .` -> inicia o modulo main da pasta atual 
- Usuários constrõem os testes
- Tem _pointers_

## Sintaxe Básica

```go
package name 
// main é o main entry point, enquanto qualquer outro nome é o nome composto do módulo
// apenas o módulo main aceita go run .

// coment

import(
    "Go  native dependencies"
    
    "my modules"
)

// =======================================
// Key example modules
// =======================================
//fmt -> print e suas variáveis fromatadas
// dependem de fmt
//errors -> criação e gerencimaneto de erros
//log -> dunções de log e tipos de retorno
//test -> para construir testes
//regexp
//time

func main(type param) return-type{

    // apenas dentro de funções
    variable := // initializes and assigns variable_
    //or
    var variable string = fmt.Sprintf("Hi, %v. Welcome!", name)

}
```

## Tipos de dados
### Básicos

```go
bool

string

int  int8  int16  int32  int64 // Inteiros
uint uint8 uint16 uint32 uint64 uintptr // Apenas naturais

byte // pseudônimo para uint8

rune // pseudônimo para int32
     // representa um ponto de código Unicode

float32 float64

complex64 complex128
```
**Importante**
0/nil = `0`, `false`, `""`

## Conceitos
### Estrturas de dados

```go
// Estrturas de dados similares entre py e go
// Não são iguais mas são similares
//Map no go-> Dict no Python
map[key_type]value_type
variable := map[key] 
// variable = value_stored_in_key_on_map


//Slice no go -> Parece a list do Python
[]value_type{value1,value2,valu3,...}
```

### Exported names
```go
pkg.Exported // vai importar e achar
pkg.exported // Vai dar erro- nomes expor-
             // táveis precisam começar com
             // Maiúscula. Conveção é o uso
			 // do CamelCase
```

### Declarando funções

```go
func add(x, y int) int {
	return x + y
}
// é o mesmo que
func add(x int, y int) int {
	return x + y
}
// é claro se os parâmetros 
// tiverem o mesmo tipo
```

### Retorno Limpo

No go você pode retornar o valor das variáveis sem declará-las no `return`.
Por exemplo:

```go
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```

> **Atenção**: Isso não é nada recomendado pois pode dificultar a legibilidade

### _Loops_

- Go só possui o `for` loop
- Esse for loop possui 3 expressões separadas por ponto e vírgula;
    - 1ª expressão - roda antes da primeira iteração (ocpional)
    - 2ª expressão - roda antes de cada iteração
    - 3ª expressão - roda após cada iteração (opcional)

```go
for exp1; exp2; exp3{
   fmt.Println("hello world") 
}
```

Ao se retirar as expressões opcionais temos um "`while` loop":

```go
// Isto é um while loop
for exp2 {
   fmt.Println("hello world") 
}
```

Sem expressão alguma, temos um loop infinito

```go
for {
   fmt.Println("hello world") domain/module
}
```

### Condicional (`if`)

```go
if cond {
    do.Things()
    return true
} return false
```
O if também pode ter declarações separadas por ;. Essa declaração é escutada antes da condição. A variável criada aqui (se alguma), pertence ao **escopo do if**

```go
if variável := "coisas"; cond {
    do.Things()
    return variável // cond is true
} return false // cond is false
```

Contudo as varíaveis declaradas no `if` também passam pro `else`

```go
if variável := "coisas"; cond {
    do.Things()
    return variável // cond is true
} else cond2 { 
    variable := variável
    return variable // cond2 is true
} // cond is false
```

### Switch

Mesma lógica que `case` no BASh / `match` no Python
> Switch sem uma condição é mesma coisa que switch true

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("macOS.")
		case "linux":
			fmt.Println("Linux.")
		default:
			// freebsd, openbsd,
			// plan9, windows...
			fmt.Printf("%s.\n", os)
	}
}
```

### Defer
Espera a saída da próxima função para daí poder executá-la
> Defer tem lógica de _stack_: a última função deferida é a primeira a ser chamada ao final da função após defer

```go
package main

import "fmt"
// Nesse caso world espera hello
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```

## PONTEIROS/_POINTERS_

Go tem ponteiros, mas sem aritimética

```go
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // aponta para i, logo guarda o endereço de i
	fmt.Println(*p) // lê o valor de i através do ponteiro, 
					// ou seja "manda buscar" o valor de i
	*p = 21         // define/modifica i através do ponteiro
	fmt.Println(i)  // mostra o novo valor de i

	p = &j         // aponta para j
	*p = *p / 37   // divide j pelo ponteiro
	fmt.Println(j) // mostra o novo valor de j
}
```