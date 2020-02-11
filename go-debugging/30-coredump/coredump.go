package main

import (
	"log"
	"net/http"
	"syscall"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HELLO WORLD\n"))

		// Uh oh...
		syscall.Kill(syscall.Getpid(), syscall.SIGQUIT)
	})

	// panic("test")

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
