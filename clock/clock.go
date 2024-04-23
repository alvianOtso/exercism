package clock

import (
	"fmt"
)

const (
	dayMinutes  = 1440
	hourMinutes = 60
)

// Clock type
type Clock struct {
	minutes int
}

func normalize(minutes int) Clock {
	if minutes < 0 {
		return Clock{dayMinutes + (minutes % dayMinutes)}
	} else if minutes >= dayMinutes {
		return Clock{minutes % dayMinutes}
	}
	return Clock{minutes}
}

// New returns a Clock that represents the hour and minutes.
func New(hour, minutes int) Clock {
	return normalize(hour*hourMinutes + minutes)
}

// Add returns a Clock that represents the Clock plus minutes.
func (c Clock) Add(minutes int) Clock {
	return normalize(c.minutes + minutes)
}

// Subtract returns a Clock that represents the Clock minus minutes.
func (c Clock) Subtract(minutes int) Clock {
	return normalize(c.minutes - minutes)
}

// String returns the Clock formatted into hours:minutes, e.g. 01:02.
func (c *Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/hourMinutes, c.minutes%hourMinutes)
}