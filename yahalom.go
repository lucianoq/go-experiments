package main

import (
	"caesar"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func initiator(a chan string, b chan string, j chan string) {

	keyA := 53453453

	// BUILD MESSAGE 1
	nonceA := strconv.Itoa(rand.Int())
	initiator := "alice"
	responder := "bob"

	msg1 := initiator + " " + nonceA
	fmt.Println("[Alice] Message 1 to send to Bob: " + msg1)

	// SEND
	r := time.Duration(rand.Intn(5000))
	time.Sleep(time.Millisecond * r)
	b <- msg1

	msg3 := <-a
	fmt.Println("[Alice] Message 3 received: " + msg3)

	array := strings.Split(msg3, " ")
	if len(array) != 6 {
		panic("[Alice] Security problem. Incomprehensible message")
	}
	_msg3_part1_a := array[0] + " " + array[1] + " " +
		array[2] + " " + array[3]
	_msg3_part2_b := array[4] + " " + array[5]

	msg3_part1 := caesar.Decode(_msg3_part1_a, keyA)
	array = strings.Split(msg3_part1, " ")
	responder_2 := array[0]
	keyAB, _ := strconv.Atoi(array[1])
	nonceA_2 := array[2]
	nonceB := array[3]

	if responder != responder_2 {
		panic("[Alice] Security problem. Different nonceA")
	}

	if nonceA != nonceA_2 {
		panic("[Alice] Security problem. Different nonceA")
	}

	// _initiator_b and _keyAB_b just got and incomprehensible
	_nonceB_ab := caesar.Encode(nonceB, keyAB)

	// BUILD MESSAGE 4
	msg4 := _msg3_part2_b + " " + _nonceB_ab
	fmt.Println("[Alice] Message 4 to send to Bob: " + msg4)

	// SEND
	r = time.Duration(rand.Intn(5000))
	time.Sleep(time.Millisecond * r)
	b <- msg4

	return
}

func responder(a chan string, b chan string, j chan string) {
	keyB := 12755767

	// RECEIVE MESSAGE 1
	r := time.Duration(rand.Intn(5000))
	time.Sleep(time.Millisecond * r)
	msg1 := <-b
	fmt.Println("[Bob] Message 1 received: " + msg1)

	//BUILD MESSAGE 2
	responder := "bob"
	nonceB := strconv.Itoa(rand.Int())
	array := strings.Split(msg1, " ")
	if len(array) != 2 {
		panic("[Bob] Security problem. Incomprehensible message")
	}
	initiator := array[0]
	nonceA := array[1]

	toEncrypt := initiator + " " + nonceA + " " + nonceB
	_msg_b := caesar.Encode(toEncrypt, keyB)
	msg2 := responder + " " + _msg_b
	fmt.Println("[Bob] Message 2 to send to Jeeves: " + msg2)

	//SEND
	r = time.Duration(rand.Intn(5000))
	time.Sleep(time.Millisecond * r)
	j <- msg2

	//RECEIVE MESSAGE 4
	r = time.Duration(rand.Intn(5000))
	time.Sleep(time.Millisecond * r)
	msg4 := <-b

	fmt.Println("[Bob] Message 4 received: " + msg4)

	array = strings.Split(msg4, " ")
	if len(array) != 3 {
		panic("[Bob] Security problem. Incomprehensible message")
	}
	_msg4_part1_b := array[0] + " " + array[1]
	_msg4_part2_ab := array[2]

	msg4_part1 := caesar.Decode(_msg4_part1_b, keyB)

	array = strings.Split(msg4_part1, " ")
	initiator_2 := array[0]
	keyAB, _ := strconv.Atoi(array[1])

	nonceB_2 := caesar.Decode(_msg4_part2_ab, keyAB)

	if nonceB != nonceB_2 {
		panic("[Bob] Security problem. Different NonceB")
	}

	if initiator != initiator_2 {
		panic("[Bob] Security problem. Different initiator")
	}

	return
}

func server(a chan string, b chan string, j chan string) {

	keyA := 53453453
	keyB := 12755767

	//RECEIVE MESSAGE 2
	r := time.Duration(rand.Intn(5000))
	time.Sleep(time.Millisecond * r)
	msg2 := <-j
	fmt.Println("[Jeeves] Message 2 received: " + msg2)

	array := strings.Split(msg2, " ")
	if len(array) != 4 {
		panic("[Jeeves] Security problem." +
			"Incomprehensible message")
	}
	responder := array[0]
	toDecrypt := array[1] + " " + array[2] + " " + array[3]
	decrypted := caesar.Decode(toDecrypt, keyB)
	array = strings.Split(decrypted, " ")
	initiator := array[0]
	nonceA := array[1]
	nonceB := array[2]

	//Create a new session key
	keyAB := rand.Int()

	// BUILD MESSAGE 3
	toEncrypt_a := responder + " " + strconv.Itoa(keyAB) + " " +
		nonceA + " " + nonceB
	toEncrypt_b := initiator + " " + strconv.Itoa(keyAB)

	_msg3_part1_a := caesar.Encode(toEncrypt_a, keyA)
	_msg3_part2_b := caesar.Encode(toEncrypt_b, keyB)

	msg3 := _msg3_part1_a + " " + _msg3_part2_b
	fmt.Println("[Jeeves] Message 3 to send to Alice: " + msg3)

	//SEND
	r = time.Duration(rand.Intn(5000))
	time.Sleep(time.Millisecond * r)
	a <- msg3

	return
}

func main() {

	rand.Seed(time.Now().UnixNano())

	a := make(chan string)
	b := make(chan string)
	j := make(chan string)

	go initiator(a, b, j)
	go responder(a, b, j)
	go server(a, b, j)

	time.Sleep(time.Second * 20)

	return
}
