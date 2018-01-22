// Copyright 2016 ~ 2017 AlexStocks(https://github.com/AlexStocks).
// All rights reserved.  Use of this source code is
// governed by Apache License 2.0.

// Package gxtime encapsulates some golang.time functions
package gxtime

import (
	"strconv"
	"time"
)

func TimeSecondDuration(sec int) time.Duration {
	return time.Duration(sec) * time.Second
}

func TimeMillisecondDuration(m int) time.Duration {
	return time.Duration(m) * time.Millisecond
}

func TimeMicrosecondDuration(m int) time.Duration {
	return time.Duration(m) * time.Microsecond
}

func TimeNanosecondDuration(n int) time.Duration {
	return time.Duration(n) * time.Nanosecond
}

// desc: convert year-month-day-hour-minute-seccond to int in second
// @month: 1 ~ 12
// @hour:  0 ~ 23
// @minute: 0 ~ 59
func YMD(year int, month int, day int, hour int, minute int, sec int) int {
	return int(time.Date(year, time.Month(month), day, hour, minute, sec, 0, time.Local).Unix())
}

// @YMD in UTC timezone
func YMDUTC(year int, month int, day int, hour int, minute int, sec int) int {
	return int(time.Date(year, time.Month(month), day, hour, minute, sec, 0, time.UTC).Unix())
}

func YMDPrint(sec int, nsec int) string {
	return time.Unix(int64(sec), int64(nsec)).Format("2006-01-02 15:04:05.99999")
}

func Future(sec int, f func()) {
	time.AfterFunc(TimeSecondDuration(sec), f)
}

func Unix2Time(unix int64) time.Time {
	return time.Unix(unix, 0)
}

func UnixNano2Time(nano int64) time.Time {
	return time.Unix(nano/1e9, nano%1e9)
}

func UnixString2Time(unix string) time.Time {
	i, err := strconv.ParseInt(unix, 10, 64)
	if err != nil {
		panic(err)
	}

	return time.Unix(i, 0)
}

// 注意把time转换成unix的时候有精度损失，只返回了秒值，没有用到纳秒值
func Time2Unix(t time.Time) int64 {
	return t.Unix()
}

func Time2UnixNano(t time.Time) int64 {
	return t.UnixNano()
}
