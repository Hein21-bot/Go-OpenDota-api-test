package utils

import "time"

func CalculateMatchStatus(startTime int64, duration int) string {
	now := time.Now().Unix()

	if duration > 0 {
		return "FINISHED"
	}

	if now >= startTime {
		return "LIVE"
	}

	return "UPCOMING"
}

func GetMatchStatus(startTime int64, duration int) string {
	now := time.Now().Unix()

	if duration > 0 {
		return "finished"
	}

	if startTime <= now {
		return "live"
	}

	return "upcoming"
}

func FormatUnix(ts int64) string {
	return time.Unix(ts, 0).Format("2006-01-02 15:04")
}

func TodayRange() (int64, int64) {
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return start.Unix(), start.Add(24 * time.Hour).Unix()
}

func TomorrowRange() (int64, int64) {
	start, _ := TodayRange()
	return start + 86400, start + 2*86400
}

func WeekRange() (int64, int64) {
	now := time.Now()
	return now.Unix(), now.AddDate(0, 0, 7).Unix()
}
