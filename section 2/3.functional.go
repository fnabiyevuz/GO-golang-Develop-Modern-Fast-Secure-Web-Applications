package main

import "fmt"

func main() {

	x := 17
	y := 32.0
	z := "Doe"

	done, trueOrNot := outputThat(x, y, z)
	fmt.Println(done, trueOrNot)
}

// func (r receiver) identifier(arg argumentType) (returnType(s)) { Code }
func outputThat(a int, b float64, c string) (string, bool) {

	fmt.Printf("Integer = %d, FloatingPoint = %f, String = %s\n", a, b, c)

	return "Okay.", true
}
