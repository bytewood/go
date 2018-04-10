package main

import "fmt"

var io = 0

func main() {

	io = 1

	defer fmt.Println(getio())

	io = 2

	fmt.Print("i = ")
}

func getio() int {
	return io;
}
