package main

import "fmt"

var role1 = "Billionaire"

func main() {
	// role1 := "Playboy"
	var role2 = "Inventor"
	fmt.Println("Bruce Wayne is", role1)
	fmt.Println("Bruce Wayne is", role2)

	tellMeWhoYouAre("Batman")
}

func tellMeWhoYouAre(role1 string) string {
	fmt.Println("Bruce Wayne is", role1)
	return role1
}
