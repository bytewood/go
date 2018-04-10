package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println(time.Now().Location())
	fmt.Println(time.Now().Local())
	fmt.Println(time.Now())
	fmt.Println(time.Now().YearDay())
	fmt.Println(time.Now().Weekday())
	fmt.Println(time.Now().Day())
	fmt.Println(time.Saturday)
}
