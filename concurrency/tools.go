package main

import (
	"fmt"
)

func main() {
	c := make(chan string)

	c <- "hello"

	greeting := <-c

	fmt.Println(greeting)
}

func helloWorld() {
	fmt.Println("Hello world")
}
