package gigasecond

import "time"

// AddGigasecond adds one giga second to the input
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
