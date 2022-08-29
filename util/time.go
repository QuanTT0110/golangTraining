package util

import "time"

func CreateExpiryByHours(hours int64) int64 {
	return time.Now().Add(time.Hour * time.Duration(hours)).Unix()
}

func CreateExpiryByMinutes(min int64) int64 {
	return time.Now().Add(time.Minute * time.Duration(min)).Unix()
}

func ConvertUnixToTime(tm int64) time.Time {
	return time.Unix(tm, 0)
}
