package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	uniq_letters := ""

	word = strings.ToLower(word)
	word = strings.Replace(word, " ", "", -1)
	word = strings.Replace(word, "-", "", -1)
	for _, letter := range word {
		s_letter := string(letter)
		if strings.Contains(uniq_letters, s_letter) {
			return false
		}
		uniq_letters += s_letter
		// fmt.Println("uniq_letters:", uniq_letters)
	}

	return true
}
