package manet

import (
	"fmt"
	"math/rand"
	"node"
	"packet"
	"strconv"
)

const (
	NumNodi    = node.NumNodi
	Spazio     = 50
	RadioRange = 5
)

type Manet struct {
	Nodes [NumNodi]node.Node
}

func NewManet() *Manet {
	var m Manet
	fmt.Println("INIT MANET")
	for i := uint8(0); i < NumNodi; i++ {
		m.Nodes[i].Id = i
		m.Nodes[i].Ch = make(chan packet.Packet)
		m.Nodes[i].Position = uint8(rand.Intn(int(Spazio)))
		m.Nodes[i].RadioRange = RadioRange
	}

	for _, n := range m.Nodes {
		go n.Start(m.Nodes)
	}

	go m.Nodes[0].Generate(m.Nodes)

	return &m
}

func (m Manet) Print() {
	fmt.Println("PRINT MANET")
	for i := 0; i < Spazio; i++ {
		present := false
		for idx := 0; idx < NumNodi; idx++ {
			if i == int(m.Nodes[idx].Position) {
				if present {
					fmt.Print("-")
				}
				present = true
				fmt.Print(strconv.Itoa(int(m.Nodes[idx].Id)))
			}
		}
		if !present {
			fmt.Print(" ")
		}
		fmt.Print("|")
	}
	fmt.Println()
	return
}
