package xtime

import "time"

func NowDay() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 0, 0, 0, now.Location())
}
func PrevDay() time.Time {
	now := time.Now()
	prevDay := now.AddDate(0, 0, -1)
	return time.Date(
		prevDay.Year(),
		prevDay.Month(),
		prevDay.Day(),
		23,
		0,
		0,
		0,
		prevDay.Location(),
	)
}
func PrevWeek() time.Time {
	now := time.Now()
	prevDay := now.AddDate(0, 0, -7)
	return time.Date(
		prevDay.Year(),
		prevDay.Month(),
		prevDay.Day(),
		23,
		0,
		0,
		0,
		prevDay.Location(),
	)
}
func NowEndOfDay() time.Time {
	now := time.Now()
	return time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		23,
		59,
		59,
		999999999,
		now.Location(),
	)
}
func EndTimeEndOfDay(times time.Time) time.Time {
	return time.Date(
		times.Year(),
		times.Month(),
		times.Day(),
		23,
		59,
		59,
		999999999,
		times.Location(),
	)
}
