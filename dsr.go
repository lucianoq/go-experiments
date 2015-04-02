package main

import (
	"fmt"
	"manet"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("")

	m := *manet.NewManet()
	fmt.Println("Creata la MANET")

	m.Print()

	time.Sleep(time.Second * 4)

	return
}
