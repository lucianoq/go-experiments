package main

import (
	"fmt"
	"time"
)

func fibonacci(i int) (f int) {
	switch {
	case i == 0:
		f = 0
	case i == 1:
		f = 1
	default:
		f = fibonacci(i-1) + fibonacci(i-2)
	}
	return
}

func main() {
	t0 := time.Now()
	for i := 0; i < 45; i++ {
		fmt.Println(fibonacci(i))
	}
	fmt.Println("Durata: ", time.Now().Sub(t0))
}
