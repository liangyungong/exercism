package techpalace

import (
	s "strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + s.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	banner := s.Repeat("*", numStarsPerLine)

	return s.Join([]string{banner, welcomeMsg, banner}, "\n")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	msg_no_star := s.Replace(oldMsg, "*", "", -1)
	return s.TrimSpace(msg_no_star)
}
