package main

import "fmt"

func main() {
	// welcome.Hello()
	// basics.RandomNumbers()
	// basics.Problems()
	// fmt.Println("The sum is:", basics.Add(35, 13))
	// a, b := basics.Swap("hello", "world")
	// fmt.Println("Swap:", a, b)
	// var i int
	// fmt.Println(i)
	// basics.LangNames()
	// basics.ChangeType()
	// basics.Constants()

	// fmt.Println(basics.NeedInt(basics.Small))
	// fmt.Println(basics.NeedFloat(basics.Small))
	// fmt.Println(basics.NeedFloat(basics.Big))

	// var name string
	var name string
	fmt.Println("What is your name?")
	_, err := fmt.Scanln(&name)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Hello, %s!\n", name) // You might want to handle the error here
}
