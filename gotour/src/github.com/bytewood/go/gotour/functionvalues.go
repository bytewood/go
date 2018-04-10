package main

import (
	"math"
	"fmt"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	funVal()
	closures()

	f := fibonacci()
	for i :=0; i < 10; i++ {
		fmt.Print(f(), " ")
	}
}

func funVal() {
	hypotenuse := func (x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypotenuse(5,12))

	fmt.Println(compute(hypotenuse))
	fmt.Println(compute(math.Pow))
}

func closures() {
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	a := 0
	b := 1
	return func () int{
		r := a + b
		a = b
		b = r
		return r
	}
}
