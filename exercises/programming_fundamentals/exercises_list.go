package programming_fundamentals

import "fmt"

func ExercisesList() {
	fmt.Println("<- Fundamentos da programação ->")
	E01()
	E02()
	E03()
	E04()
	E05()
	E06()
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

// Crie um programa que:
// Atribua um valor int a uma variável
// Demonstre este valor em decimal, binário e hexadecimal
// Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
// Demonstre esta outra variável em decimal, binário e hexadecimal
func E04() {
	x := 10

	fmt.Println("\n=== E04 ===")
	fmt.Printf("%d\t%b\t%#x", x, x, x)
	y := x << 1
	fmt.Printf("\n%d\t%b\t%#x\n", y, y, y)

}

// Crie uma variável de tipo string utilizando uma raw string literal.
// Demonstre-a.
func E05() {
	s := `Esta é uma 
				raw string 
						literal.`
	fmt.Println("\n=== E05 ===")
	fmt.Println(s)
}

// Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
// Demonstre estes valores.
func E06() {
	const (
		_ = 2025 + iota
		a
		b
		c
		d
	)
	fmt.Println("\n=== E06 ===")
	fmt.Println(a, b, c, d)
}
