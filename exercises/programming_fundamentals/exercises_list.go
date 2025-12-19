package programming_fundamentals

import "fmt"

func ExercisesList() {
	fmt.Println("<- Fundamentos da programação ->")
	E01()
	E02()
	E03()
}

// E01 Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
func E01() {
	d := 10
	b := 10
	h := 200
	fmt.Println("\n=== E01 ===")
	fmt.Printf("decimal: %d\nbinário: %b\nhexadecimal: %#x\n", d, b, h)
}

// E02 Escreva expressões utilizando os operadores:
// == != <= < >= >
// e atribua seus valores a variáveis.
// Demonstre estes valores.
func E02() {
	fmt.Println("\n=== E02 ===")
	n1 := 10
	n2 := 30
	a := n1 == n2
	b := n1 != n2
	c := n1 <= n2
	d := n1 < n2
	e := n1 >= n2
	f := n1 > n2

	fmt.Println(a, b, c, d, e, f)
}

// E03 Crie constantes tipadas e não-tipadas.
// Demonstre seus valores.
func E03() {
	const a string = "10"
	const b = 10
	fmt.Println("\n=== E03 ===")
	fmt.Printf("a = %v, %T\nb = %v, %T\n", a, a, b, b)
}

// TODO E04
