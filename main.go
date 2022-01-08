package main

func main() {
	cards := newDeck()
	// cards.print()
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	// data := cards.toByte()
	// saveToFile("SebaToDisk", data, 0644)
	cards.saveToFile("SebaToDisk")
}
