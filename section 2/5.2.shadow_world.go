package main

import (
	"fmt"
	"time"
)

type FamilyMember struct {
	FamilyName string
	FirstName  string
	Age        int
	Birthdate  time.Time
	Species    string
}

func main() {
	francine := FamilyMember{
		FamilyName: "Smith",
		FirstName:  "Francine",
		Species:    "Human",
	}

	roger := FamilyMember{
		FamilyName: "Smith",
		FirstName:  "Roger",
		Age:        1600,
		Species:    "Alien",
	}

	fmt.Println(francine)
	fmt.Println(francine.FirstName)

	fmt.Println(roger)
	fmt.Println("Roger is", roger.Age, "earth years old.")

}
