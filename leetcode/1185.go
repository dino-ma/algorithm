package main

import "fmt"

const _MinDay int = 1
const _MaxDay int = 31
const _MinMonth int = 1
const _MaxMonth int = 12
const _MinYear int = 1971
const _MaxYear int = 2100

var monthDays = [12]int{
	31,
	28,
	31,
	30,
	31,
	30,
	31,
	31,
	30,
	31,
	30,
	31,
}
var weeks = [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

// dayOfTheWeek 1971 to 2100 return week
func dayOfTheWeek(day int, month int, year int) string {
	if day < _MinDay || day > _MaxDay {
		return ""
	}
	if month < _MinMonth || month > _MaxMonth {
		return ""
	}
	if year < _MinYear || year > _MaxYear {
		return ""
	}
	// 初始化需要返回的天数
	var needDays int
	// 计算年份的天数
	needDays = (year-_MinYear)*365 + (year-1969)/4
	// 为什么年份减1969 (1971年之前的上一个闰年)
	// 因为1968年为闰年 ，可以让4或400整除并且不被100整除，所以才会对能被 (year - 1969)/4的闰年有贡献
	// needDays 通过年份运算可以计算出天数，
	// 计算月份中贡献的天数
	for _, d := range monthDays[:month-1] {
		needDays += d
	}
	// 判断是否为闰年
	// 来自百度 闰年分为普通闰年和世纪闰年，其判断方法为：公历年份是4的倍数，且不是100的倍数，为普通闰年。公历年份是整百数，且必须是400的倍数才是世纪闰年。归结起来就是通常说的：四年一闰；百年不闰，四百年再闰。闰年是为了弥补因人为历法规定造成的年度天数与地球实际公转周期的时间差而设立的。
	if month >= 3 && (year%400 == 0 || year%4 == 0 && year%100 != 0) {
		needDays++
	}
	// 输入月份中的天数贡献
	// 运算 返回
	return weeks[(needDays+day+3)%7]
}

func main() {
	fmt.Println(dayOfTheWeek(31, 8, 2019))
	fmt.Println(dayOfTheWeek(18, 7, 1999))
	fmt.Println(dayOfTheWeek(15, 8, 1993))
	fmt.Println(dayOfTheWeek(1, 1, 1971))
}
