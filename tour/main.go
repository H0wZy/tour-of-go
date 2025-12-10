package main

import (
	"fmt"

	welcome "github.com/H0wZy/tour-of-go/tour/1-welcome"
	basics "github.com/H0wZy/tour-of-go/tour/2-basics"
)

func main() {
	welcome.Hello()
	basics.RandomNumbers()
	basics.Problems()
	fmt.Println("The sum is:", basics.Add(35, 13))
	a, b := basics.Swap("hello", "world")
	fmt.Println("Swap:", a, b)
	var i int
	fmt.Println(i)
	basics.LangNames()
	basics.ChangeType()
	basics.Constants()

	fmt.Println(basics.NeedInt(basics.Small))
	fmt.Println(basics.NeedFloat(basics.Small))
	fmt.Println(basics.NeedFloat(basics.Big))
}
