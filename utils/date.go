package utils

import (
	"fmt"
	"math"
	"time"
)

/*
func GetDate(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04")
}
func GetDateMH(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("01-02 03:04")
}*/

func GetDateFormat(timestamp int64, format string) string {
	if timestamp <= 0 {
		return ""
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format(format)
}

func GetDate(timestamp int64) string {
	if timestamp <= 0 {
		return ""
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02")
}

func GetDateMH(timestamp int64) string {
	if timestamp <= 0 {
		return ""
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04")
}

func GetDateMHS(timestamp int64) string {
	if timestamp <= 0 {
		return ""
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:01")
}

func GetDateDiff(date1 string, date2 string) string {
	layout := "2006-01-02 15:04:05"
	var diff string
	t1, _ := time.Parse(layout, date1)
	t2, _ := time.Parse(layout, date2)
	t := t1.Sub(t2)

	if t.Seconds() < 0 {
		diff = "0天0小时0分0秒"
	} else {
		days := math.Floor(t.Hours() / 24)
		hours := math.Floor(t.Hours()) - (days * 24)
		minutes := math.Floor(t.Minutes()) - (math.Floor(t.Hours()) * 60)
		seconds := t.Seconds() - (math.Floor(t.Minutes()) * 60)

		diff = fmt.Sprintf("%v天%v小时%v分%v秒", days, hours, minutes, seconds)
	}

	return diff
}

func GetDateDiffColor(date1 string, date2 string) string {
	layout := "2006-01-02 15:04:05"
	var level string
	t1, _ := time.Parse(layout, date1)
	t2, _ := time.Parse(layout, date2)
	t := t1.Sub(t2)

	if t.Seconds() < 0 {
		level = ""
	} else {
		days := math.Floor(t.Hours() / 24)
		hours := math.Floor(t.Hours()) - (days * 24)
		//minutes := math.Floor(t.Minutes()) - (math.Floor(t.Hours()) * 60)
		//seconds := t.Seconds() - (math.Floor(t.Minutes()) * 60)

		if days >= 1 {
			return "item-red"
		} else if hours >= 1 {
			return "item-yellow"
		} else {
			return ""
		}
	}

	return level
}

func GetTimeParse(times string) int64 {
	if "" == times {
		return 0
	}
	loc, _ := time.LoadLocation("Local")
	parse, _ := time.ParseInLocation("2006-01-02 15:04", times, loc)
	return parse.Unix()
}

func GetDateParse(dates string) int64 {
	if "" == dates {
		return 0
	}
	loc, _ := time.LoadLocation("Local")
	parse, _ := time.ParseInLocation("2006-01-02", dates, loc)
	return parse.Unix()
}
