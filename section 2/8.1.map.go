package main

import "fmt"

func main() {
	fmt.Println()

	aSlice := []string{"Fry", "Leela", "Bender", "Amy", "Hubert", "Zoidberg", "Hermes", "Scruffy", "Nibbler"}
	fmt.Println(aSlice)

	aString := "One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin."
	fmt.Println(aString)

	aMap := map[string]string{"dog": "Scooby Doo", "cat": "Garfield", "mouse": "Jerry", "platybus": "Perry", "bird": "Tweety", "rabbit": "Bugs", "elephant": "Dumbo"}
	fmt.Println(aMap)

}
