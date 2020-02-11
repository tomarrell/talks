package main

import "fmt"

// What will this code print?

func main() {
	a := getRand()
	fmt.Println(a)
}

func getRand() int {
	v := 3
	defer oddFunc(&v)

	return v
}

func oddFunc(v *int) {
	*v = 1923801928307774927
}
