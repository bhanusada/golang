package calendar

import (
	"errors"
)

//Date Type
type Date struct {
	year  int
	month int
	day   int
}

//Event type
type Event struct {
	title string
	Date
}

//SetTitle of event
func (e *Event) SetTitle(title string) {
	e.title = title
}

//GetTitle of event
func (e *Event) GetTitle() string {
	return e.title
}

//SetYear of Date
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

//GetYear to return year
func (d *Date) GetYear() int {
	return d.year
}

//SetMonth of Date
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("Invalid month")
	}
	d.month = month
	return nil
}

//GetMonth to return month
func (d *Date) GetMonth() int {
	return d.month
}

//SetDay of Date
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid day")
	}
	d.day = day
	return nil
}

//GetDay to return day
func (d *Date) GetDay() int {
	return d.day
}
