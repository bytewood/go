package main

import (
	"fmt"
	"math"
	"runtime"
)

func sqrt(x float64) (z float64, iter int) {
	z = x/2
	d := dfunc(z, x)
	for d > 0.000001 {
		fmt.Println(d, math.Sqrt(x), z)
		z -= (z*z - float64(x)) / (2 * z)
		if iter++; iter > 100 {
			break
		}
		d = dfunc(z, x)
	}
	return
}

func dfunc(z float64, x float64) float64 {
	return math.Abs(z*z - x)
}

func main() {
	v := float64(2)
	sqrt(v)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOROOT())
}
