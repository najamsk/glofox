package main

import (
	"fmt"
	"time"

	"github.com/najamsk/glofox/src/api/models"
	uuid "github.com/satori/go.uuid"
)

func main() {

	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006"
	)

	date := "1999-12-31"
	t, _ := time.Parse(layoutISO, date)
	fmt.Println(t)                  // 1999-12-31 00:00:00 +0000 UTC
	fmt.Println(t.Format(layoutUS)) // December 31, 1999

	fmt.Println("now = ", time.Now())
	today := time.Now().Format("2006-01-02")
	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

	fmt.Println("start", today)
	fmt.Println("tomorrow", tomorrow)

	u2 := uuid.NewV4()
	startDate := "2021-02-14"
	tStart, _ := time.Parse(layoutISO, startDate)
	endDate := "2021-02-22"
	tEnd, _ := time.Parse(layoutISO, endDate)

	fmt.Printf("UUIDv4: %s\n", u2)
	class := &models.Class{Name: "najam awan", ID: u2, Capacity: 20, StartDate: tStart, EndDate: tEnd}
	fmt.Printf("class = %+v \n", class)

}
