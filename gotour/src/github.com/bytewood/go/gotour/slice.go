package main

import (
	"fmt"
	"golang.org/x/tour/pic"
)

func main() {
	pic.Show(generatePicture)
}

func generatePicture(dx int, dy int) [][]uint8 {
	ry := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		ry[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			ry[y][x] = uint8((x + y) / 2)
		}
	}

	return ry
}

func appendToSlice() {
	a := [10]int{100, 101, 102, 103, 104}

	s1 := a[3:7]
	s2 := a[1:5]

	printSlice(s1)
	printSlice(s2)

	s1 = append(s1, 7, 8, 9, 10, 11, 12)

	//after append s2 no longer points to the same array as s1
	a[4] = 114

	printSlice(s1)
	printSlice(s2)
}

func forRangeOfSlice() {
	s := make([]int, 10)
	s[5] = 10

	for itemIndex, itemValue := range s {
		s[itemIndex] = itemValue + itemIndex
	}

	fmt.Printf("%v\n", s)

	sum := 0
	for _, itemValue := range s {
		sum += itemValue
	}
	fmt.Println(sum)

	for itemIndex := range s {
		sum -= itemIndex
	}

	fmt.Println(sum)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
