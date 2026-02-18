package main

import (
	"fmt"
)

func main() {
	Append()
}

// FUNÇÃO APPEND DEMONSTRANDO COMO USAR A FUNÇÃO APPEND PARA ADICIONAR ELEMENTOS A UM SLICE
func Append() {
	slice := []int{1, 2, 3}
	slize := []int{4, 5, 6}

	slice = append(slice, slize...)

	fmt.Println("Slice after appending:", slice)
}
