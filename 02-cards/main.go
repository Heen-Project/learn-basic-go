package main

// go basic data types: bool, string, int, float64
// Array: static length array -- Slice:  dynamic length array
func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
