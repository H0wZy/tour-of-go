package vars_values_types

import "fmt"

// Utilizando a solução do exercício anterior:
// 1. Em package-level scope, atribua os seguintes valores às variáveis:
//     1. para "x" atribua 42
//     2. para "y" atribua "James Bond"
//     3. para "z" atribua true
// 2. Na função main:
//     1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando
// o operador curto de declaração.
//     2. Demonstre a variável "s".

func E03() {
	fmt.Println("=== E03 ===")
	// Reutilizando a solução do exercício anterior
	x = 42
	y = "James Bond"
	z = true

	s := fmt.Sprintf("x = %v, y = %v, z = %v", x, y, z)
	fmt.Println(s)
}
