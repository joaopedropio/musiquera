package domain

import "fmt"

type Date interface {
	Day() int
	Month() int
	Year() int
	String() string
}

type date struct {
	day   int
	month int
	year  int
}

func (d *date) Day() int {
	return d.day
}

func (d *date) Month() int {
	return d.month
}

func (d *date) Year() int {
	return d.year
}

func (d *date) String() string {
	return fmt.Sprintf("%d-%02d-%02d", d.year, d.month, d.day)
}

func NewDate(year, month, day int) Date {
	return &date{
		day,
		month,
		year,
	}
}
