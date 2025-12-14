package exercises_list

import "fmt"

// Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
//     1. 42
//     2. "James Bond"
//     3. true
// Agora demonstre os valores nestas variáveis, com:
//     1. Uma única declaração print
//     2. Múltiplas declarações print

func E01() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(y, "have", x, "years old and is a", z, "spy")
	fmt.Printf("x = %v, %T\n", x, x)
	fmt.Printf("y = %v, %T\n", y, y)
	fmt.Printf("z = %v, %T\n", z, z)

}

func E02() {

}
