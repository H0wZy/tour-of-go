package basics

import (
	"fmt"
	"math"
	"math/rand"
)

func RandomNumbers() {
	wanted := 13
	random := rand.Intn(20)
	fmt.Println("My favorite number is", wanted, "but i received", random)

	if favorite := wanted; random == favorite {
		fmt.Println("You guessed my favorite number!", favorite)
	}
}

func Problems() {
	random := rand.Float64() * 100
	problems := math.Sqrt(random)
	fmt.Printf("Now you have %g problems.\n", problems)
}

func CanExport() {
	if true {
		fmt.Println(math.Pi)

	} else {
		// fmt.Println(math.pi) // unexported: starts with lowercase letter")
	}

}

func Add(a, b int) int {
	return a + b
}

func Swap(x, y string) (string, string) {
	return y, x
}

func LangNames() {
	var c, python, java bool
	fmt.Println(c, python, java)
}

func ChangeType() {
	v := 0.323 + 0.5i // change type by reassigning
	fmt.Printf("v is of type %T\n", v)
}

func Constants() {
	const Pi = 3.14159265358979323
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy Pi Day!", Pi)

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func NeedInt(x int) int           { return x*10 + 1 }
func NeedFloat(x float64) float64 { return x * 0.1 }
