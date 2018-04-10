package main

import (
	"fmt"
	"math"
	"runtime"
	"math/rand"
	"time"
)


func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	v := float64(2)
	sqrt(v)

	forLoop()
	ifelse()

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOROOT())
}

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

func forLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum +=i
	}
	fmt.Println(sum);

	sum = 1
	for sum < 5 {
		sum += sum
	}
	fmt.Println(sum)

	for {
		fmt.Println("Forever")
		if 1==1 {
			break
		}
	}
}

func ifelse() {
	if i:=rand.Intn(10); i < 5 {
		fmt.Println("less than 5 :", i)
	} else {
		fmt.Println("greateer than or equal to :", i)
	}
	//fmt.Println(i)
}
