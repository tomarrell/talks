package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func main() {
	// Required to report blocking events to the profiler
	runtime.SetBlockProfileRate(1)

	ch := make(chan struct{})
	go writeRead(ch)

	fmt.Println("Starting web server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func writeRead(ch chan struct{}) {
	time.Sleep(2 * time.Second)

	ch <- struct{}{}
	fmt.Println(<-ch)
}
