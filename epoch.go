package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	unixTime := now.Unix()
	fmt.Println("Current Unix Time:", unixTime)

	t := time.Unix(unixTime, 0)
	fmt.Println("Date Time:", t)

	fmt.Println("Time", t.Format("02 January 2006"))
}
