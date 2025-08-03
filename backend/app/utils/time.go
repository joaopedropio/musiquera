package utils

import "time"

func IsTimeEqual(first, second time.Time) bool {
	a := first.UTC()
	b := second.UTC()
	return a.Day() == b.Day() &&
		a.Month() == b.Month() &&
		a.Year() == b.Year() &&
		a.Hour() == b.Hour() &&
		a.Minute() == b.Minute() &&
		a.Second() == b.Second()
}
