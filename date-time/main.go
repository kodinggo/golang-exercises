package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	var t time.Time
	fmt.Println(t)

	// timeNow := time.Now()

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	t = time.Now().In(loc)
	fmt.Println(t)

	// Parsing dari string ke time.Time
	timeStr := "2024-07-19 09:40:22" // string
	t, err = time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		log.Fatalf("error parse string time: %v", err)
	}

	fmt.Println(t.Day(), t.Year()) // mendapatkan single unit of datetime e.g. year, day etc

	timeStr = "19/07/2024 09:40:22"
	t, err = time.Parse("02/01/2006 15:04:05", timeStr)
	if err != nil {
		log.Fatalf("error parse string time: %v", err)
	}

	fmt.Println(t.Day(), t.Year())

	// Konversi dari time.Time ke string
	fmt.Println(t.Format("2006 Jan 02"))

	// timeProcess := time.Since(timeNow)
	// fmt.Println(timeProcess)

	timeNow := time.Now()      // WIB
	timeNow.Add(1 * time.Hour) // menambahkan 1 jam dari sekarang

	// Mendapatkan start and end date
	startDate := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, timeNow.Location())
	endDate := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 23, 59, 59, 0, timeNow.Location())

	fmt.Println(startDate.UTC(), endDate.UTC())

	// mendapatkan sisa tanggal di bulan ini
	// misalnya sekarang tgl 19 -> 30

	startOfMonth := time.Date(timeNow.Year(), timeNow.Month(), 1, 0, 0, 0, 0, timeNow.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-1 * time.Second)
	fmt.Println("the end day of this month", endOfMonth.Day())
	theRestDayOfTheMonth := endOfMonth.Day() - timeNow.Day()
	fmt.Println("the rest day of this month", theRestDayOfTheMonth)
}
