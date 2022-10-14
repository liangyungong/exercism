package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	Hour   int
	Minute int
}

func positive_minutes(m int) int {
	for m < 0 {
		m += 60 * 24
	}
	return m
}

func minutes_to_h_m(minute int) (int, int) {
	h := (minute / 60) % 24
	m := minute % 60
	return h, m
}

func New(h, m int) Clock {
	total_min := h*60 + m

	total_min = positive_minutes(total_min)
	h, m = minutes_to_h_m(total_min)
	return Clock{
		Hour:   h,
		Minute: m,
	}

}

func (c Clock) Add(m int) Clock {
	total_min := c.Minute + c.Hour*60 + m
	total_min = positive_minutes(total_min)
	h, m := minutes_to_h_m(total_min)
	return Clock{
		Hour:   h,
		Minute: m,
	}
}

func (c Clock) Subtract(m int) Clock {
	total_min := c.Hour*60 + c.Minute - m
	total_min = positive_minutes(total_min)
	h, m := minutes_to_h_m(total_min)
	return Clock{
		Hour:   h,
		Minute: m,
	}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}
