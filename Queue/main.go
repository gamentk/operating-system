package main

import "fmt"

func main() {
	a := [10] int{}
	a[0] = 10
	a[1] = 20
	a[2] = 30

	for i := range a {
		fmt.Printf("%d \n", a[i])
	}
}