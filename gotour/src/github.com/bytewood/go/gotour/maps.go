package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func main() {
	maps()
	mapLiterals()
	mapLiteralsTerse()
	mutateAMap()

	testWordCount()
}

type Coordinate struct {
	Lat, Long float64
}

var m map[string]Coordinate

func maps() {
	m = make(map[string]Coordinate)
	m["Bytewood Labs"] = Coordinate{41.231, -78.765,}

	fmt.Println(m["Bytewood Labs"])

}

func mapLiterals() {
	var m = map[string]Coordinate{
		"a b" : Coordinate{35.342, -76.987,},
		"c d" : Coordinate{23.657, -98.654,},
	}
	fmt.Println(m)
}

func mapLiteralsTerse() {
	var m = map[string]Coordinate{
		"a b" : {35.342, -76.987,},
		"c d" : {23.657, -98.654,},
	}
	fmt.Println(m)
}

func mutateAMap() {
	m := map[string]int{}

	m["abc"] = 1
	m["def"] = 2
	m["ghi"] = 3

	delete(m, "def")

	fmt.Println(m)

	elem, ok := m["xyz"]

	fmt.Println("Value : ", elem, "Present?", ok)
}

func wordCount(s string)map[string]int {
	count := map[string]int{}
	for _, v := range strings.Fields(s) {
		count[v] += 1
	}
 	return count
}

func testWordCount() {
	wc.Test(wordCount)
}