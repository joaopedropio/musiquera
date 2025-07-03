package domain

import "fmt"

type Date interface {
	Day() uint
	Month() uint
	Year() uint
	String() string
}

type date struct {
	day   uint
	month uint
	year  uint
}

func (d *date) Day() uint {
	return d.day
}

func (d *date) Month() uint {
	return d.month
}

func (d *date) Year() uint {
	return d.year
}

func (d *date) String() string {
	return fmt.Sprintf("%d-%d-%d", d.year, d.month, d.day)
}

func NewDate(year, month, day uint) Date {
	return &date{
		day,
		month,
		year,
	}
}
