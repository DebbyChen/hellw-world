package main

import "fmt"

func work(c chan string) {
	c <- "Get a new chance"
}

func main() {
	c := make(chan string, 1)
	work(c)
	s := <-c
	fmt.Println("Get mesg:", s)
}
