package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	s := "hello"
	p := &s
	fmt.Println(*p, s)

	v := Vertex{1, 2}
	vp := &v
	vp.X = 3
	fmt.Println("v :",v)
	fmt.Println(v.X)

	fmt.Println(Vertex{X: 4})
	fmt.Println(Vertex{})
	fmt.Printf("%T\n", vp)


	var a [10]int
	a[0] = 0
	a[1] = 1
	fmt.Println("array a : ", a)
	fmt.Println(a[1])

	b := [10]int{}
	fmt.Println("array b :", b)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	fmt.Println("slice : ")
	var slice []int
	if (slice == nil) {
		fmt.Println("nil")
		fmt.Println(len(slice))
		fmt.Println(cap(slice))
	}
	slice = primes[1:4]
	fmt.Println(slice)

	var sliceb = primes[3:6]
	slice[2] = 6

	fmt.Println(primes)
	fmt.Println(slice)
	fmt.Println(sliceb)

	sliceLiteral := []int{1,2,3,4,5,6,7}
	fmt.Println("sliceLiteral : ", sliceLiteral)

	sliceStructLiteral := []struct{
		b bool
		s string
	}{
		{false, "a a"},
		{true, "b"},
		{true, "c"},
	}

	fmt.Println("SliceStructLiteral : ", sliceStructLiteral)
	fmt.Println(sliceStructLiteral[0].s)
	fmt.Println(len(primes))
	fmt.Println(len(slice))
	fmt.Println(cap(slice))


}