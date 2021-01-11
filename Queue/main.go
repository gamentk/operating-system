package main

import (
	"fmt"
	"time"
)

func show(a int) {
	for {
		fmt.Printf("I'm %d\n", a)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go show(5)
	go show(10)
	for {
		fmt.Printf("I am main.\n")
		time.Sleep(500 * time.Millisecond)
	}
}
