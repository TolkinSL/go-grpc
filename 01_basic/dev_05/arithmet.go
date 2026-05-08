package main

import (
	"fmt"
	"math"
)

func main() {
	//Считается выражение потом только выведение типа
	var f1 = 10 / 3
	fmt.Println("10 / 3 =", f1)

	var f2 = 10 / 3.0
	fmt.Println("10 / 3.0 =", f2)

	// Overflow
	var over int8 = 127
	over += 1
	fmt.Println(over)

	// Underflow
	var overF float64 = 2e-310
	overF = overF / math.MaxFloat64 //слишком маленькое для типа значит = 0
	fmt.Println(overF)
}