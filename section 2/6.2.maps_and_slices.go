package main

import "fmt"

func main() {
	var meeseeks []string
	fmt.Printf("%T\n", meeseeks)
	fmt.Println(meeseeks, len(meeseeks), cap(meeseeks))
	meeseeks = []string{"Meeseek#1"}
	fmt.Println(meeseeks, len(meeseeks), cap(meeseeks))

	someIntergers := []int{42, 23, 1000, -1}
	fmt.Printf("%T\n", someIntergers)
	fmt.Println(someIntergers, len(someIntergers), cap(someIntergers))
	someIntergers = append(someIntergers, 987, 654, 321, 123, 456, 789)
	fmt.Println(someIntergers, len(someIntergers), cap(someIntergers))
	someIntergers = append(someIntergers, 555)
	fmt.Println(someIntergers, len(someIntergers), cap(someIntergers))

	someFloats := make([]float64, 3, 6)
	fmt.Printf("%T\n", someFloats)
	fmt.Println(someFloats, len(someFloats), cap(someFloats))
	someFloats[0] = 42.42
	someFloats = append(someFloats, 23.23, 123.456)
	fmt.Println(someFloats, len(someFloats), cap(someFloats))
	someFloats = append(someFloats, 555.555, 888.888)
	fmt.Println(someFloats, len(someFloats), cap(someFloats))
}
