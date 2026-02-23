package datagroup

import "fmt"

func Slices() {
	sabores := []string{"Calabresa", "Queijo", "Doce"}
	fatias := sabores[:]
	fmt.Printf("%s\n", fatias)
}
