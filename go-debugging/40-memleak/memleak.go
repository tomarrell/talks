package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

type shape struct {
	edges    int
	vertices int
}

func main() {
	go addShapes()

	fmt.Println("Starting web server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func addShapes() {
	items := []shape{}

	for {
		items = append(items, shape{3, 3})
		time.Sleep(1 * time.Microsecond)
	}
}
