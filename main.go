package main

func main() {
	cards := newDeckFromFile("SebaToDisk");
	cards.shuffle()
	cards.print()
}
