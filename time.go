package gofast

import (
	"time"
)

// VietnamTimeLoc returns location +07:00
func VietnamTimeLoc() *time.Location {
	loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err == nil {
		return loc
	}
	t0, err := time.Parse(time.RFC3339, "2020-03-27T11:00:31+07:00")
	if err == nil {
		return t0.Location()
	}
	return time.Local
}

// VietnamTimeNow returns now in location +07:00
func VietnamTimeNow() time.Time {
	return time.Now().In(VietnamTimeLoc())
}

// VietnamTimeNowIso returns now in format 2006-01-02T15:04:05+07:00
func VietnamTimeNowIso() string {
	return VietnamTimeNow().Format(time.RFC3339)
}

func VietnamDateNowIso() string {
	return VietnamTimeNow().Format("2006-01-02")
}
