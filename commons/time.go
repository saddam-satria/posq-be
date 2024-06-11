package commons

import "time"

func UTCTime() time.Time {
	localTime := time.Now()

	utcTime := localTime.UTC()
	return utcTime
}

func LocalTime() time.Time {
	localTime := time.Now()

	return localTime
}