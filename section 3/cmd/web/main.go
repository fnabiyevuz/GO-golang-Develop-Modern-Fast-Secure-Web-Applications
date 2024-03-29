package main

import (
	"fmt"
	"net/http"

	handlers "github.com/jagottsicher/myGoWebApplication/pkg/handler"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
