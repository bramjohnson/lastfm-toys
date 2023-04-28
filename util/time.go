package util

import (
	"fmt"
	"time"
)

type MonthRange struct {
	StartOfMonth time.Time
	EndOfMonth   time.Time
}

func (mr *MonthRange) shiftDatesBy(duration time.Duration) MonthRange {
	return MonthRange{StartOfMonth: mr.StartOfMonth.Add(duration), EndOfMonth: mr.EndOfMonth.Add(duration)}
}

func (mr *MonthRange) NextMonth() MonthRange {
	return MonthRange{StartOfMonth: mr.StartOfMonth.AddDate(0, 1, 0), EndOfMonth: mr.EndOfMonth.AddDate(0, 1, 0)}
}

func (mr *MonthRange) LastMonth() MonthRange {
	return MonthRange{StartOfMonth: mr.StartOfMonth.AddDate(0, -1, 0), EndOfMonth: mr.EndOfMonth.AddDate(0, -1, 0)}
}

func GetMonthRangeDate(date time.Time) MonthRange {
	currentYear, currentMonth, _ := date.Date()
	currentLocation := date.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, 0)

	fmt.Println(firstOfMonth, lastOfMonth)

	return MonthRange{StartOfMonth: firstOfMonth, EndOfMonth: lastOfMonth}
}

func GetMonthRange() MonthRange {
	return GetMonthRangeDate(time.Now())
}
