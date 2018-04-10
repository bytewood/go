package main;

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

const (
	A_MESSAGE = "Message A"
	B_MESSAGE = "Message B"
	BIG = 1 << 100
	SMALL = BIG >> 99

)

// package level declared variables; basic go types; factored into a block as with import statement
var (
	b bool = false
	s string = ""
	i int = 0
	i8 int8
	i16 int16
	i32 int32
	i64 int64
	u uint
	u8 uint8
	u16 uint16
	u32 uint32
	u64 uint64
	up uintptr

	by byte //alias for uint8
	r rune //alias for int32; represents a Unicode code point
	f32 float32
	f64 float64
	c64 complex64
	c128 complex128

)
func main()  {
	rand.Seed(0)
	fmt.Println("Hello to Playground")
	fmt.Println("The time is", time.Now())
	fmt.Println("A Random Number :", rand.Intn(100))
	fmt.Println("PI :", math.Pi)
	fmt.Println("Sum", sum(33,78))

	// short assignment := in place of var inside func onlu
	a, b := swap("1","2")
	fmt.Println("Multiple return types", a, b)
	fmt.Println(split(100))

	// function level declared variables
	var m, n string = "Woo", "Boo"
	fmt.Println("Function level declared variables", m, n)
	fmt.Printf("Type cast from int to %T\n", cast())
	fmt.Printf("Type of -32 %T\n", -32)
	// no downcasting
	//fmt.Printf("Downcast of -32 %T\n", int32(math.MaxInt32))

	c64 = 3 + 4i

	fmt.Println("complex number", c64)
	fmt.Println("Big", float64(BIG))
	fmt.Println("Small", SMALL)
}

//
func sum(x int, y int) int {
	return x + y
}
// multiple return values
func swap(a, b string) (string, string) {
	return b, a
}

// named return values
func split (sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum -x
	return
}

// explicit casting; type inference
func cast() uint {
	oi := 42
	of := float64(oi)
	return uint(of)
}