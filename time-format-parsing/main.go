package main

import (
	"fmt"
	"time"
)

func main() {
	// Penggunaan time.Time merupakan representasi object data-time
	// menjadikan informasi waktu sekarng sebagai object time.Time dengan time.Now()
	// membuat object baru bertipe time.Time dengan time.Date()

	time1 := time.Now()
	fmt.Printf("time1 %v\n", time1)
	// time1 2025-02-18 10:18:38.31716952 +0700 WIB m=+0.000073886

	time2 := time.Date(2021, 12, 24, 10, 20, 0, 0, time.UTC)
	// time.Date(tahun, bulan, tanggal, jam, menit, detik, nanodetik, timezone)
	fmt.Printf("time2 %v\n", time2)
	// time2 2021-12-24 10:20:00 +0000 UTC

	// time.Time merupakan struct yang memiliki beberapa method yang bisa dipakai
	now := time.Now()
	fmt.Println("year:", now.Year(), "month:", now.Month())

	// Parsing dari string ke time.Time menggunakan time.Parse

	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = "2015-09-02 08:04:00"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())
	// 2015-09-02 08:04:00 +0000 UTC

	layoutFormat = "02/01/2006 MST"
	value = "02/09/2015 WIB"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t\t->", date.String())
	// 2015-09-02 00:00:00 +0700 WIB

	// Predefined layout format
	// time.RFC822 dan format lainnya yang bisa di lihat di docs
	date2, _ := time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
	fmt.Println(date2.String())
	// 2015-09-02 08:00:00 +0700 WIB

	// Format dari time.Time ke string menggunakan Format()
	dateS1 := date2.Format("Monday 02, January 2006 15:04 MST")
	fmt.Println("dateS1", dateS1)
	// Wednesday 02, September 2015 08:00 WIB

	dateS2 := date2.Format(time.RFC3339)
	fmt.Println("dateS2", dateS2)
	// 2015-09-02T08:00:00+07:00

	// Handle error parsing time.Time
	date3, err := time.Parse("06 Jan 15", "02 Sep 15 08:00 WIB")
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}

	fmt.Println(date3)
}
