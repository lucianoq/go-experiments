package main

import (
	"fmt"
	"time"
)

func fibonacci(i int, ch chan int) {
	switch {
	case i == 0:
		ch <- 0
	case i == 1:
		ch <- 1
	default:
		ch <- fib(i-1) + fib(i-2)
	}
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

	go fibonacci(0, ch)
	go fibonacci(1, ch)
	go fibonacci(2, ch)
	go fibonacci(3, ch)
	go fibonacci(4, ch)
	go fibonacci(5, ch)
	go fibonacci(6, ch)
	go fibonacci(7, ch)

	for i := 8; i < 45; i++ {
		x := <-ch
		fmt.Println(x)
		go fibonacci(i, ch)
		i++
	}

	x := <-ch
	fmt.Println(x)
	x = <-ch
	fmt.Println(x)
	x = <-ch
	fmt.Println(x)
	x = <-ch
	fmt.Println(x)
	x = <-ch
	fmt.Println(x)
	x = <-ch
	fmt.Println(x)
	x = <-ch
	fmt.Println(x)
	x = <-ch
	fmt.Println(x)

	fmt.Println("Durata: ", time.Now().Sub(t0))
}
