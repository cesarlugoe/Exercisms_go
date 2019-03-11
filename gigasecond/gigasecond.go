package gigasecond

import (
	"time"
)

// AddGigasecond comment for it being able to be exported without error to the test
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * 1000000000)
	return t
}
