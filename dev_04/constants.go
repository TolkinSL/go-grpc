package main

import "fmt"

const PI = 3.14
const TELE = 2

func main() {
	const days int = 7

	const (
		t1 = 1
		t2 = 2
		t3
		t4
	)
	fmt.Println("t3", t3)

	const (
		m1 = iota
		m2
		m3
	)
	fmt.Println("m iota", m1, m2, m3)

}