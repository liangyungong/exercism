package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse(
		"1/02/2006 15:04:05",
		date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	form := "January 2, 2006 15:04:05"

	t, _ := time.Parse(form, date)
	return t.UTC().Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	form := "Monday, January 2, 2006 15:04:05"

	t, _ := time.Parse(form, date)
	hour := t.UTC().Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	form := "1/2/2006 15:04:05"

	t, _ := time.Parse(form, date)

	output_form := "Monday, January 2, 2006, at 15:04"
	return fmt.Sprintf("You have an appointment on %s.",
		t.Format(output_form))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	year := time.Now().UTC().Year()
	date_s := fmt.Sprintf("%d-09-15T00:00:00+00:00", year)

	t, _ := time.Parse(
		time.RFC3339,
		date_s)

	return t
}
