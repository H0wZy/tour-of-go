package vars_values_types

import "fmt"

func ExercisesList() {
	fmt.Println("<- Variáveis, valores e tipos ->")
	E01()
	E02()
	E03()
	E04()
	E05()
}

// E01 Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
//  1. 42
//  2. "James Bond"
//  3. true
//
// Agora demonstre os valores nestas variáveis, com:
//  1. Uma única declaração print
//  2. Múltiplas declarações print
func E01() {
	fmt.Println("\n=== E01 ===")
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(y, "have", x, "years old and is a", z, "spy")
	fmt.Printf("x = %v, %T\n", x, x)
	fmt.Printf("y = %v, %T\n", y, y)
	fmt.Printf("z = %v, %T\n", z, z)

}

var x int
var y string
var z bool

// E02 Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos
// para estas variáveis:
//  1. Identificador "x" deverá ter tipo int
//  2. Identificador "y" deverá ter tipo string
//  3. Identificador "z" deverá ter tipo bool
//
// Na função main:
//  1. Demonstre os valores de cada identificador
//  2. O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?
func E02() {
	fmt.Println("\n=== E02 ===")

	fmt.Printf("x = %v, %T\n", x, x)
	fmt.Printf("y = %v, %T\n", y, y)
	fmt.Printf("z = %v, %T\n", z, z)

	fmt.Println("Os valores atribuídos pelo compilador se chamam zero values.")
}

// E03 Utilizando a solução do exercício anterior:
// 1. Em package-level scope, atribua os seguintes valores às variáveis:
//  1. para "x" atribua 42
//  2. para "y" atribua "James Bond"
//  3. para "z" atribua true
//
// 2. Na função main:
//  1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando
//
// o operador curto de declaração.
//  2. Demonstre a variável "s".
func E03() {
	fmt.Println("\n=== E03 ===")
	// Reutilizando a solução do exercício anterior
	x = 42
	y = "James Bond"
	z = true

	s := fmt.Sprintf("x = %v, y = %v, z = %v", x, y, z)
	fmt.Println(s)
}

// E04 Crie um tipo. O tipo subjacente deve ser int.
// Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
// Na função main:
//  1. Demonstre o valor da variável "x"
//  2. Demonstre o tipo da variável "x"
//  3. Atribua 42 à variável "x" utilizando o operador "="
//  4. Demonstre o valor da variável "x"
//
// Para os aventureiros: https://go.dev/ref/spec#Types
func E04() {
	type myType int
	var x myType
	fmt.Println("\n=== E04 ===")
	fmt.Printf("x = %v, %T\n", x, x)
	x = 42
	fmt.Printf("x = %v, %T\n", x, x)
}

var yy int

// E05 Utilizando a solução do exercício anterior:
//  1. Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "yy".
//     O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
//  2. Na função main:
//  1. Isto já deve estar feito:
//  1. Demonstre o valor da variável "x"
//  2. Demonstre o tipo da variável "x"
//  3. Atribua 42 à variável "x" utilizando o operador "="
//  4. Demonstre o valor da variável "x"
//  2. Agora faça também:
//  1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e,
//     utilizando o operador "=", atribua o valor de "x" a "yy"
//  2. Demonstre o valor de "yy"
//  3. Demonstre o tipo de "yy"
func E05() {

	yy = int(x)

	fmt.Println("\n=== EX05 ===")
	fmt.Printf("yy = %v, %T\n", yy, yy)
}
