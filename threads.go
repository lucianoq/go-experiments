package main

import (
	"fmt"
	"time"
)

func thread(i time.Duration, ch chan int) {
	fmt.Println("Thread ", i, " started")
	t0 := time.Now()

	time.Sleep(time.Second * i)
	ch <- 9
	fmt.Println("Thread ", i, " ended. Durata: ", time.Now().Sub(t0))
	return
}

func main() {
	t0 := time.Now()

	ch := make(chan int)

	go thread(3, ch)
	go thread(6, ch)
	go thread(1, ch)

	x := <-ch
	fmt.Println(x)
	x = <-ch
	fmt.Println(x)
	x = <-ch
	fmt.Println(x)

	fmt.Println("TOTALE Durata: ", time.Now().Sub(t0))
	return
}
