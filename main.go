package main

func main() {
	s := []string{"First", "Second"}

	oldPosition := 0
	newPosition := 1

	// swap values
	s[oldPosition], s[newPosition] = s[newPosition], s[oldPosition]

	// typing "newP" and hitting tab will result in this:
	// s[o], s[newPnewPosition]

	// however, this bug does not happen on the right hand side of the assignment
	// nor does it happen when going back later and re-editing the left hand side
}
