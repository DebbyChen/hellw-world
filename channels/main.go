package main

import "fmt"

func work(c chan string) {
	c <- "Hello!"
}

func main() {
	c1 := make(chan string)
	go work(c1)

	ss := <-c1
	fmt.Println(ss)
}
