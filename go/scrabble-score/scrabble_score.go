package scrabble

import "strings"

func lookUpScore(letter string) int {
	letter = strings.ToLower(letter)
	score_map := map[string]int{
		"aeioulnrst": 1,
		"dg":         2,
		"bcmp":       3,
		"fhvwy":      4,
		"k":          5,
		"jx":         8,
		"qz":         10,
	}
	for str, score := range score_map {
		if strings.Contains(str, letter) {
			return score
		}
	}
	return 0
}
func Score(word string) int {
	total := 0
	for _, letter := range word {
		total += lookUpScore(string(letter))
	}
	return total
}
