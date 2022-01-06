package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var number int
	switch card {
	case "ace":
		number = 11
	case "two":
		number = 2
	case "three":
		number = 3
	case "four":
		number = 4
	case "five":
		number = 5
	case "six":
		number = 6
	case "seven":
		number = 7
	case "eight":
		number = 8
	case "nine":
		number = 9
	case "ten", "jack", "queen", "king":
		number = 10
	default:
		number = 0
	}
	return number
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return ParseCard(card1)+ParseCard(card2) == 21 //|| card1 == "ace" && card2 =="ace"
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	if isBlackjack {
		if dealerScore < 10 {
			return "W"
		} else if dealerScore >= 10 {
			return "S"
		}
	}
	return "P"
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {

	if handScore > 16 {
		return "S"
	} else if handScore > 11 && dealerScore < 7 {
		return "S"
	} else {
		return "H"
	}
}
