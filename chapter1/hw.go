package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println('M')

	// prints the data type
	fmt.Println(reflect.TypeOf("hello"))
	fmt.Println(reflect.TypeOf(42))

	// zero values
	var name string
	var age int
	var isStudent bool
	var score float64

	fmt.Println(name, age, isStudent, score)

	// variable declaration and initialization
	var originalCount int = 10
	var eatenCount int = 4

	fmt.Println("I Started with", originalCount, "apples")
	fmt.Println("Some Jerk ate", eatenCount, "apples")
	fmt.Println("There are", originalCount-eatenCount, "apples left")

	// only variable, functions and types that starts with captial letters are
	// considered exported
}
