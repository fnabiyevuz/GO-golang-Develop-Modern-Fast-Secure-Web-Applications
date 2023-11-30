package main

import "fmt"

var x = 23

func main() {

	fmt.Println(x)
	fmt.Printf("x is of Type %T\n", x)

	y := 42.0
	fmt.Println(y)
	fmt.Printf("y is of Type %T\n", y)

	z := "Doh"
	fmt.Println(z)
	fmt.Printf("z is of Type %T\n", z)
}
