package node

import (
	"fmt"
	"math/rand"
	"packet"
	"strconv"
	"time"
)

type Node struct {
	Id         uint8
	Ch         chan packet.Packet
	Position   uint8
	RadioRange uint8
	cache      [20]uint16
	nCache     uint8
}

const NumNodi = 30

func (this Node) Start(nodes [NumNodi]Node) {
	for {
		var p packet.Packet
		p = <-this.Ch

		var toMe, toRoute, isReply bool = false, false, false

		if isIn(p.Id, this.cache, this.nCache) {
			continue
		} else {
			this.cache[this.nCache] = p.Id
			this.nCache++

			if p.To == this.Id {
				toMe = true
				toRoute = false
			} else {
				toMe = false
				toRoute = true
			}
			if p.Reply {
				isReply = true
			}
		}

		if toRoute {
			if isReply {
				//SEND BOTTOM UP
				dest := p.Hops[p.NHops-1]
				p.NHops--
				go send(p, nodes[dest])
			} else {
				//SEND BROADCAST RADIORANGE
				p.Hops[p.NHops] = this.Id
				p.NHops++

				this.SendBroadcast(p, nodes)
			}
			continue
		}

		if toMe {
			if isReply {
				//OK RETURNED HOME
				fmt.Println("[" + strconv.Itoa(int(this.Id)) +
					"] Pack" +
					strconv.Itoa(int(p.Id)) +
					" Reply returned home")
				fmt.Print("PATH FROM: " + strconv.Itoa(int(p.To)) +
					" TO: " + strconv.Itoa(int(p.From)) +
					" is: ")
				p.PrintPath()
			} else {
				//SEND REPLY
				fmt.Println("[" + strconv.Itoa(int(this.Id)) +
					"] Pack" + strconv.Itoa(int(p.Id)) +
					" Arrivato a destinazione e invio Reply")

				var pr packet.Packet
				pr.From = p.To
				pr.To = p.From
				pr.Id = uint16(rand.Uint32() % 65535)
				pr.Reply = true
				pr.Path = p.Hops
				pr.Hops = p.Hops
				pr.NHops = p.NHops
				pr.NPath = p.NHops
				pr.Path[p.NPath] = this.Id
				pr.NPath++

				dest := pr.Hops[pr.NHops-1]
				pr.NHops--

				this.cache[this.nCache] = pr.Id
				this.nCache++
				go send(pr, nodes[dest])
			}
			continue
		}
	}
}

func isIn(i uint16, l [20]uint16, size uint8) bool {
	for e := uint8(0); e < size; e++ {
		if i == l[e] {
			return true
		}
	}
	return false
}

func (this Node) Distance(a Node) (dist uint8) {
	var _dist int = int(this.Position) - int(a.Position)
	if _dist < 0 {
		_dist = -_dist
	}
	dist = uint8(_dist)
	return dist
}

func (this Node) Generate(nodes [NumNodi]Node) {
	for i := 0; i < 10; i++ {
		t := time.Duration(250)
		time.Sleep(time.Millisecond * t)

		nonce := uint16(rand.Uint32() % 65535)

		var p packet.Packet
		p.Reply = false
		p.NHops = 0
		p.NPath = 0
		p.Id = nonce
		p.From = this.Id

		p.To = uint8(rand.Intn(int(NumNodi)))
		for p.To == p.From {
			p.To = uint8(rand.Intn(int(NumNodi)))
		}

		p.Hops[p.NHops] = this.Id
		p.NHops++

		this.cache[this.nCache] = p.Id
		this.nCache++

		fmt.Println("[" + strconv.Itoa(int(this.Id)) + "] Pack" +
			strconv.Itoa(int(p.Id)) + " CREATO.  FROM: " +
			strconv.Itoa(int(p.From)) + " TO: " +
			strconv.Itoa(int(p.To)))

		this.SendBroadcast(p, nodes)
	}
	return
}

func (this Node) SendBroadcast(p packet.Packet, nodes [NumNodi]Node) {
	inoltrato := false
	for i := uint8(0); i < NumNodi; i++ { //tra tutti i nodi
		if this.Id != nodes[i].Id { //tranne me stesso
			if this.Distance(nodes[i]) <= this.RadioRange {
				inoltrato = true
				go send(p, nodes[i])
			}
		}
	}

	if !inoltrato {
		fmt.Println("[" + strconv.Itoa(int(this.Id)) + "] Pack" +
			strconv.Itoa(int(p.Id)) +
			" Non inoltrabile da qui")
	}
}

func send(p packet.Packet, n Node) {
	n.Ch <- p
}
