package L_1154

import (
	"fmt"
	"strconv"
	"time"
)

// 给你一个字符串 date ，按 YYYY-MM-DD 格式表示一个 现行公元纪年法 日期。请你计算并返回该日期是当年的第几天。

// 通常情况下，我们认为 1 月 1 日是每年的第 1 天，1 月 2 日是每年的第 2 天，依此类推。每个月的天数与现行公元纪年法（格里高利历）一致。

// Parse后直接输出
func dayOfYear_(date string) int {
	t, _ := time.Parse("2006-01-02", date)
	return t.YearDay()
}

// 转数字后用Date类，比上面那个直接Parse效率高
func dayOfYear__(date string) int {
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).YearDay()
}

// 转数字后，手动计算，效率更高些
func dayOfYear(date string) int {
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])
	yearDay := 0
	switch month {
	case 2:
		yearDay += 31 // 31
	case 3:
		yearDay += 59 // 30
	case 4:
		yearDay += 90 // 31
	case 5:
		yearDay += 120 // 30
	case 6:
		yearDay += 151 // 31
	case 7:
		yearDay += 181 // 30
	case 8:
		yearDay += 212 // 31
	case 9:
		yearDay += 243 // 31
	case 10:
		yearDay += 273 // 30
	case 11:
		yearDay += 304 // 31
	case 12:
		yearDay += 334 // 30
	}
	if month > 2 && (year%400 == 0 || (year%4 == 0 && year%100 != 0)) {
		yearDay += 1
	}
	return yearDay + day
}

func Test1154() {
	fmt.Println(dayOfYear("2019-01-09")) // 9
	fmt.Println(dayOfYear("2019-02-10")) // 41
	fmt.Println(dayOfYear("2003-03-01")) // 60
	fmt.Println(dayOfYear("2004-03-01")) // 61
	fmt.Println(dayOfYear("2008-10-10")) // 284
	fmt.Println(dayOfYear("1992-09-14")) // 258

}
