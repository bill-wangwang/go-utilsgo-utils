package utils

import "time"

func DateToTimestamp(str string) int64 {
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		return 0
	}
	return t.Unix()
}

func DateTimeToTimestamp(str string) int64 {
	t, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		return 0
	}
	return t.Unix()
}

func TimestampToDate(t int64) string {
	return time.Unix(t, 0).Format("2006-01-02")
}
func TimestampToDateTime(t int64) string {
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}
