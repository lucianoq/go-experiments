package packet

import (
	"fmt"
	"strconv"
)

type Packet struct {
	Id    uint16
	From  uint8
	To    uint8
	Reply bool
	Hops  [20]uint8
	Path  [20]uint8
	NHops uint8
	NPath uint8
}

func (p Packet) PrintPath() {
	fmt.Println()
	for e := 0; e < int(p.NPath); e++ {
		fmt.Print(strconv.Itoa(int(p.Path[e])) + " ")
	}
	fmt.Println()
}

func (p Packet) PrintHops() {
	fmt.Println()
	for e := 0; e < int(p.NHops); e++ {
		fmt.Print(strconv.Itoa(int(p.Hops[e])) + " ")
	}
	fmt.Println()
}
