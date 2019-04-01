package utils

import (
	"fmt"
	"time"
)

func GetNow() time.Time {
	return time.Now()
}

func GetNowString() string {
	return time.Now().String()
}

func GetNowTimes() int64 {
	return time.Now().Unix()
}

func GetNowFormat() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetNowFormatYYYYMMDDHHMMSS(value string) time.Time {
	times, err := time.Parse("2006-01-02 15:04:05", value)
	fmt.Println(times, err)
	return times
}
