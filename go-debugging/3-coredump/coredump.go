package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("HELLO WORLD")
	})

	aList := []int{1, 2, 3, 4}
	num := aList[10]
	fmt.Println(num)

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
