package main

import (
	"fmt"

	"github.com/H0wZy/tour-of-go/exercises/programming_fundamentals"
	"github.com/H0wZy/tour-of-go/exercises/vars_values_types"
)

func space() {
	fmt.Println()
}

func main() {
	space()
	vars_values_types.ExercisesList()
	space()
	programming_fundamentals.ExercisesList()
}
