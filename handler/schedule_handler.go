package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hablullah/go-prayer"
	"github.com/jawara-insinyur/prayer-schedule-api/model"
	"github.com/ringsaturn/tzf"
)

type PrayScheduleHandler struct {
	Finder tzf.F
}

func (p *PrayScheduleHandler) PrayScheduleHandler(w http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query()

	lat, err := strconv.ParseFloat(query.Get("lat"), 64)
	if err != nil {
		return NewHttpError(err, 400, "lat can't be empty and must be float")
	}

	lon, err := strconv.ParseFloat(query.Get("lon"), 64)
	if err != nil {
		return NewHttpError(err, 400, "lon can't be empty and must be float")
	}

	year, err := strconv.Atoi(query.Get("year"))
	if err != nil {
		return NewHttpError(err, 400, "year can't be empty and must be integer")
	}

	timezone := query.Get("timezone")

	latlon := model.LatLon{Lat: lat, Lon: lon}

	data, nil := p.calculatePrayerSchedule(latlon, year, timezone)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("Error encode json: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
	return nil
}

func (p *PrayScheduleHandler) calculatePrayerSchedule(latlon model.LatLon, year int, timezone string) (*model.PrayerScheduleResponse, error) {
	var tz *time.Location
	var err error

	if len(timezone) != 0 {
		tz, err = time.LoadLocation(timezone)
		if err != nil {
			return nil, NewHttpError(err, 400, "Invalid timezone")
		}
	} else {
		tz, err = time.LoadLocation(p.Finder.GetTimezoneName(latlon.Lon, latlon.Lat))
		if err != nil {
			return nil, fmt.Errorf("Error generating timezone: %v", err)
		}
	}

	schedules, err := prayer.Calculate(prayer.Config{
		Latitude:           latlon.Lat,
		Longitude:          latlon.Lon,
		Timezone:           tz,
		TwilightConvention: prayer.Kemenag(),
		AsrConvention:      prayer.Shafii,
		PreciseToSeconds:   true,
		Corrections: prayer.ScheduleCorrections{
			Fajr:    2 * time.Minute,
			Sunrise: -4 * time.Minute,
			Zuhr:    3 * time.Minute,
			Asr:     2 * time.Minute,
			Maghrib: 2 * time.Minute,
			Isha:    2 * time.Minute,
		},
	}, year)
	if err != nil {
		return nil, fmt.Errorf("Error calculating pray schedule: %v", err)
	}

	var data []model.DataSchedule

	for _, v := range schedules {
		data = append(
			data,
			model.DataSchedule{
				Date: v.Date,
				Schedule: model.Schedule{
					Fajr:    v.Fajr,
					Sunrise: v.Sunrise,
					Zuhr:    v.Zuhr,
					Asr:     v.Asr,
					Maghrib: v.Maghrib,
					Isha:    v.Isha,
				},
				ReadableSchedule: model.ReadableSchedule{
					Fajr:    model.ReadableTime{Time: v.Fajr},
					Sunrise: model.ReadableTime{Time: v.Sunrise},
					Zuhr:    model.ReadableTime{Time: v.Zuhr},
					Asr:     model.ReadableTime{Time: v.Asr},
					Maghrib: model.ReadableTime{Time: v.Asr},
					Isha:    model.ReadableTime{Time: v.Isha},
				},
			},
		)
	}

	res := &model.PrayerScheduleResponse{
		Status:   200,
		Timezone: tz.String(),
		Data:     data,
	}

	return res, nil
}
