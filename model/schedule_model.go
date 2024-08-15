package model

import (
	"fmt"
	"time"
)

type ReadableTime struct {
	time.Time
}

type PrayerScheduleResponse struct {
	Status   int            `json:"status,omitempty"`
	Timezone string         `json:"timezone,omitempty"`
	Data     []DataSchedule `json:"data,omitempty"`
}

type DataSchedule struct {
	Date             string           `json:"date,omitempty"`
	Schedule         Schedule         `json:"schedule,omitempty"`
	ReadableSchedule ReadableSchedule `json:"readableSchedule,omitempty"`
}

type ReadableSchedule struct {
	Fajr    ReadableTime `json:"fajr,omitempty"`
	Sunrise ReadableTime `json:"sunrise,omitempty"`
	Zuhr    ReadableTime `json:"zuhr,omitempty"`
	Asr     ReadableTime `json:"asr,omitempty"`
	Maghrib ReadableTime `json:"maghrib,omitempty"`
	Isha    ReadableTime `json:"isha,omitempty"`
}

type Schedule struct {
	Fajr    time.Time `json:"fajr,omitempty"`
	Sunrise time.Time `json:"sunrise,omitempty"`
	Zuhr    time.Time `json:"zuhr,omitempty"`
	Asr     time.Time `json:"asr,omitempty"`
	Maghrib time.Time `json:"maghrib,omitempty"`
	Isha    time.Time `json:"isha,omitempty"`
}

func (t ReadableTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", t.Format("15:04 MST"))
	return []byte(stamp), nil
}
