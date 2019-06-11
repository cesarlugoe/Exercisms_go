package clock

import "fmt"

const (
	minutesInHour = 60
	minutesInDay  = 1440
)

// Clock comment
type Clock int

// New function comment
func New(h, m int) Clock {
	c := Clock(h*minutesInHour + m)
	return c.max24Hour()
}

//Add minutes to the Clocl
func (c Clock) Add(m int) Clock {
	c = (c + Clock(m))
	return c.max24Hour()

}

//Substract minutes to the clock
func (c Clock) Substract(m int) Clock {
	c = (c - Clock(m))
	return c.max24Hour()
}

// Checks if clock value overextends the 24 hour window
func (c Clock) max24Hour() Clock {
	c = c % minutesInDay
	if int(c) < 0 {
		return c + minutesInDay
	}
	return c
}

// String comment
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/minutesInHour, c%minutesInHour)
}
