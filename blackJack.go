package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func player(name string, p chan string, d chan string) {
	name = "[" + name + "]"
	fmt.Println(name + " ready")

	for {
		cardS := <-p
		card, _ := strconv.Atoi(cardS)
		somma := card

		<-p
		for somma < 17 {
			time.Sleep(time.Second * 1)
			fmt.Println(name + " Card")
			d <- "Card"
			cardS := <-p
			card, _ := strconv.Atoi(cardS)
			somma += card
		}

		if somma > 21 {
			fmt.Println(name + " Busted")
			d <- "Busted"
		} else {
			fmt.Println(name + " Halt")
			d <- "Halt"
		}
	}

	return
}

func main() {
	var players [3]chan string
	var numPlayers int = len(players)

	for i := 0; i < numPlayers; i++ {
		players[i] = make(chan string)
	}
	d := make(chan string)

	dealer(players, d)

	return
}

func generateCard() int {
	card := rand.Intn(13) + 1
	if card == 13 {
		card = 10
	}
	if card == 12 {
		card = 10
	}
	if card == 11 {
		card = 10
	}
	if card == 1 {
		card = 11
	}
	return card
}

func dealer(players [3]chan string, d chan string) {
	for i := 0; i < len(players); i++ {
		go player("Player"+strconv.Itoa(i), players[i], d)
	}

	i := 0
	for {
		i = i % len(players)
		card := generateCard()
		players[i] <- strconv.Itoa(card)
		handle("Player"+strconv.Itoa(i),
			players[i], d)
		i++
	}

	return
}

func handle(name string, p chan string, d chan string) {
	fmt.Println("[Dealer] Handle " + name)

	fmt.Println("[Dealer] " + name + " turn")
	p <- "YourTurn"

	str := <-d
	for str == "Card" {
		card := generateCard()
		fmt.Println("[Dealer] Sending Card: " + strconv.Itoa(card))
		p <- strconv.Itoa(card)
		str = <-d
	}

	if str == "Busted" {
		fmt.Println("[Dealer] I win!!!")
	}

	if str == "Halt" {
		fmt.Println("[Dealer] " + name + " halts")
	}
	return
}
