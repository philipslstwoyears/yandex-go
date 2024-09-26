package main

import (
	"time"
)

func NextWorkday(start time.Time) time.Time {
	if start.Weekday() == time.Friday {
		return start.Add(3 * 24 * time.Hour)
	} else if start.Weekday() == time.Saturday {
		return start.Add(2 * 24 * time.Hour)
	}
	return start.AddDate(0, 0, 1)
}
