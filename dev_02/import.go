package main

import (
	"fmt"
	mynet "net/http"
)

func main() {
	fmt.Println("Hello GO library")

	resp, err := mynet.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error net:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("HTTP Response Status:", resp.Status)
	fmt.Println(resp.TLS)
}
