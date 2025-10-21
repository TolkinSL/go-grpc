package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	arr := []int{1, 5, 7}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	fmt.Println("Break and Continue")
	for i := 1; i <= 10; i++ {
		if i == 8 {
			fmt.Println("Break for")
			break
		}

		if i%2 != 0 {
			continue
		}

		fmt.Println(i)

	}
	fmt.Println("End for")

	// для версии 1.22 и выше
	for i := range 5 {
		//i++
		fmt.Println("Range", i)
	}

	for i := 1; i <= 5; i++ {
		i++
		fmt.Println("i", i)
	}
}
