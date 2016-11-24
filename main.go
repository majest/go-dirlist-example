package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting http file sever")
	http.Handle("/", http.FileServer(http.Dir("/data")))
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		fmt.Println(err)
	}
}
