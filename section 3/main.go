package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	owner, saying := getData()
	sum := addValues(2, 2)
	fmt.Fprintf(w, fmt.Sprintf("this is about %s, \n like to say %s, \n %d", owner, saying, sum))
}

func getData() (string, string) {
	o := "FN"
	s := "NI"
	return o, s
}

func addValues(x, y int) int {
	return x + y
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(":8080", nil)
}
