package main

import "fmt"

func work(c chan bool) {
	c <- true
}

func main() {
	c := make(chan bool)
	go work(c)
	<-c
	fmt.Println("Done")
}
