package main

import (
	"fmt"
	"time"
)

func main2() {
	currentYear, currentMonth, _ := time.Now().Date()
	currentLocation := time.Now().Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, 0)

	fmt.Println(firstOfMonth.Unix(), lastOfMonth.Unix())
}
