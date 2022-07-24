package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	value := 0
	switch card {
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten":
		value = 10
	case "jack":
		value = 10
	case "king":
		value = 10
	case "queen":
		value = 10
	}
	return value
}

func CountArrayElement(a []string, element string) int {
	count := 0
	for _, v := range a {
		if v == element {
			count ++
		}
	}

	return count
}
// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cards := [] string {card1, card2, dealerCard}
	count_of_ace := CountArrayElement(cards, "ace")
	if  card1 == card2 && count_of_ace >= 2 {
		return "P"
	}

	if count_of_ace == 2 {
		return "S"
	}

	card1_int := ParseCard(card1)
	card2_int := ParseCard(card2)

	if (card1_int + card2_int == 21) {
		if (dealerCard != "ace") && (ParseCard(dealerCard) != 10) {
			return "W"
		} else {
			return "S"
		}
	}

	sum := card1_int + card2_int
	if sum >= 17 && sum <= 20 {
		return "S"
	}

	if sum >=12 && sum <= 16 {
		if ParseCard(dealerCard) >=7 {
			return "H"
		}

		return "S"
	}

	if sum <= 11 {
		return "H"
	}

	return ""
}
