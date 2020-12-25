package main

import "fmt"

func main() {
	cardPublicKey := 3418282
	doorPublicKey := 8719412

	part1(cardPublicKey, doorPublicKey)
}

func part1(cardPubKey int, doorPubKey int) {

	cardLoopSize := 0
	number := 1
	for number != cardPubKey {
		cardLoopSize++
		number = 7 * number % 20201227
	}
	fmt.Println(cardLoopSize)
	doorLoopSize := 0
	number = 1
	for number != doorPubKey {
		doorLoopSize++
		number = 7 * number % 20201227
	}
	fmt.Println(doorLoopSize)
	encKeyCard := 1
	for cardLoop := 0; cardLoop < cardLoopSize; cardLoop++ {
		encKeyCard = doorPubKey * encKeyCard % 20201227
	}
	fmt.Println("card enc key:", encKeyCard)
	encKeyDoor := 1
	for doorLoop := 0; doorLoop < doorLoopSize; doorLoop++ {
		encKeyDoor = cardPubKey * encKeyDoor % 20201227
	}
	fmt.Println("door enc key:", encKeyDoor)
}
