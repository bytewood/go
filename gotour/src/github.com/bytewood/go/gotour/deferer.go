package main

import "fmt"

var io = 0

func main() {

	io = 1

	defer fmt.Println(getIO())

	io = 2

	fmt.Print("i = ")
}

func getIO() int{
	return io
}
