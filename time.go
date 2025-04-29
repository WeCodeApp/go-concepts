package main

import (
	"fmt"
	"time"
)

func main() {
	// Mon Jan 2 15:04:05 MST 2006

	// Current local time
	fmt.Println("Current local time:", time.Now())

	// Specific time
	specificTime := time.Date(2025, time.April, 29, 10, 48, 30, 0, time.UTC)
	fmt.Println("Specific time:", specificTime)

	// Parse time
	parseTime, _ := time.Parse("2006-01-02", "2025-05-01")
	parseTime1, _ := time.Parse("06-01-02", "25-05-01")
	parseTime2, _ := time.Parse("06-1-2", "25-5-1")
	parseTime3, _ := time.Parse("06-1-2 15-04", "25-5-1 18-03")

	fmt.Println("=======")
	fmt.Println(parseTime)
	fmt.Println(parseTime1)
	fmt.Println(parseTime2)
	fmt.Println(parseTime3)
	fmt.Println("=======")

	// Formatting time
	t := time.Now()
	fmt.Println("Formatted time:", t.Format("Monday 06-01-02 03:04:05 PM"))

	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println("Current time + 24 hours:", oneDayLater)
	fmt.Println("Current time + 24 hours:", oneDayLater.Weekday())

	fmt.Println("Rounded time:", t.Round(time.Hour))

	loc, _ := time.LoadLocation("Asia/Manila")
	t = time.Date(2025, time.April, 25, 10, 48, 30, 0, time.UTC)
	tLocal := t.In(loc)
	roundedTime := t.Round(time.Hour)
	roundedTimeLocal := roundedTime.In(loc)

	fmt.Println("======")
	fmt.Println("Manila Original Time (UTC)", t)
	fmt.Println("Manila Original Time (Local)", tLocal)
	fmt.Println("Manila Rounded Time (UTC)", roundedTime)
	fmt.Println("Manila Rounded Time (Local)", roundedTimeLocal)
	fmt.Println("Truncate Time:", t.Truncate(time.Hour))

	loc2, _ := time.LoadLocation("America/New_York")

	tInNY := time.Now().In(loc2)
	fmt.Println("New York Time:", tInNY)

	t1 := time.Date(2025, time.April, 28, 12, 0, 0, 0, loc)
	t2 := time.Date(2025, time.April, 28, 9, 0, 0, 0, loc)
	duration := t2.Sub(t1)
	fmt.Println("Duration between two times:", duration)

	fmt.Println("t2 is after t1?", t2.After(t1))
}
