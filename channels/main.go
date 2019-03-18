package main

import "fmt"

func main() {
	c := make(chan string)

	go func(c chan string) {
		c <- "Hi there!"
	}(c)

	fmt.Println(<-c)
}
