package main

import (
	"fmt"
	"time"
)

func fibonacci(i int, ch chan int) {
	fmt.Println("Thread ", i, " started")
	t0 := time.Now()
	switch {
	case i == 0:
		ch <- 0
	case i == 1:
		ch <- 1
	default:
		ch <- fib(i-1) + fib(i-2)
	}
	fmt.Println("Thread ", i, " ended")
	fmt.Println("Thread ", i, " Durata: ", time.Now().Sub(t0))
}

func fib(i int) (f int) {
	switch {
	case i == 0:
		f = 0
	case i == 1:
		f = 1
	default:
		f = fib(i-1) + fib(i-2)
	}
	return
}

func main() {
	t0 := time.Now()

	ch := make(chan int)

	go fibonacci(38, ch)

	go fibonacci(39, ch)
	go fibonacci(40, ch)

	x := <-ch
	fmt.Println(x)
	x = <-ch
	fmt.Println(x)
	x = <-ch
	fmt.Println(x)

	fmt.Println("TOTALE Durata: ", time.Now().Sub(t0))
}
