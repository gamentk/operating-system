package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	data = make([]int, 0)
)

func producer() {
	for {
		if len(data) < 100 {
			mydata := rand.Intn(100)
			data = append(data, mydata)
			fmt.Printf("Producer --> %d", mydata)
		} else {
			fmt.Printf("Buffer full")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func consumer() {
	for {
		if len(data) != 0 {
			mydata := data[0]
			data = data[1:]
			fmt.Printf("Consumer --> %d\n", mydata)
		} else {
			fmt.Printf("No data\n")
		}
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	go producer()
	go consumer()
	select {}
}
