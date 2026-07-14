package clock

import (
	"fmt"
)

const (
	daysInHour   = 24
	daysInMinute = daysInHour * 60
)

// Define the Clock type here.
type Clock struct {
	hour, minute int
}

func New(h, m int) Clock {
	totalMinutes := h*60 + m
	if totalMinutes < 0 {
		totalMinutes = daysInMinute + (totalMinutes % daysInMinute)
	}
	return Clock{(totalMinutes / 60) % 24, totalMinutes % 60}
}

func (c Clock) Add(m int) Clock {
	return New(c.hour, c.minute+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hour, c.minute-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
