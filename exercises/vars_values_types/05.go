package vars_values_types

import "fmt"

// Utilizando a solução do exercício anterior:
// 1. Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "yy".
//  O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
// 2. Na função main:
//     1. Isto já deve estar feito:
//         1. Demonstre o valor da variável "x"
//         2. Demonstre o tipo da variável "x"
//         3. Atribua 42 à variável "x" utilizando o operador "="
//         4. Demonstre o valor da variável "x"
//     2. Agora faça tambem:
//         1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e,
//  utilizando o operador "=", atribua o valor de "x" a "yy"
//         2. Demonstre o valor de "yy"
//         3. Demonstre o tipo de "yy"

var yy int

func E05() {

	yy = int(x)

	fmt.Println("=== EX05 ===")
	fmt.Printf("yy = %v, %T\n", yy, yy)
}
