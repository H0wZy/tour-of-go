package main

import (
	"fmt"
)

func main() {
	IntAppend()
	StringAppend()
}

// FUNÇÃO APPEND DEMONSTRANDO COMO USAR A FUNÇÃO APPEND PARA ADICIONAR ELEMENTOS A UM SLICE
func IntAppend() {
	slice := []int{1, 2, 3}
	slize := []int{4, 5, 6}

	slice = append(slice, slize...)

	fmt.Println("Slice after appending:", slice)
}

func StringAppend() {
	slice := []string{"Hello", "World"}
	slize := []string{"Go", "Programming", "Language"}

	slice = append(slice, slize...)

	fmt.Println("Slice after appending:", slice)
}
