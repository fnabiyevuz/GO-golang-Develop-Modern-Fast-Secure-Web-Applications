package main

import (
	"fmt"
	"log"
)

func main() {
	number := 23
	log.Println("Let's talk about", number)

	if number < 0 {
		fmt.Println(number, "is negative")
	} else if number < 10 {
		fmt.Println(number, "has 1 digit")
	} else {
		fmt.Println(number, "has multiple digits")
	}
}
