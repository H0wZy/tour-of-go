package vars_values_types

import "fmt"

// Crie um tipo. O tipo subjacente deve ser int.
// Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
// Na função main:
//     1. Demonstre o valor da variável "x"
//     2. Demonstre o tipo da variável "x"
//     3. Atribua 42 à variável "x" utilizando o operador "="
//     4. Demonstre o valor da variável "x"
// Para os aventureiros: https://go.dev/ref/spec#Types
func E04() {
	type myType int
	var x myType
	fmt.Println("=== E04 ===")
	fmt.Printf("x = %v, %T\n", x, x)
	x = 42
	fmt.Printf("x = %v, %T\n", x, x)
}
