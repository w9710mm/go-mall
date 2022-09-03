package util

import time "time"

func GetDateDayStart(t time.Time) time.Time {

	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return t
}
func GetTimeStart(t time.Time) time.Time {

	t = time.Date(1970, 0, 1, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	return t
}
