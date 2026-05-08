package main

import (
	"fmt"
	"time"
)

func main() {
	go SayHello()
	time.Sleep(2 * time.Second)
}

func SayHello() {
	fmt.Println("Hello Go")
	time.Sleep(1 * time.Second)
}