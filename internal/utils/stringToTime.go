package utils

import (
	"time"
)

func (m *Utils) StringToTime(sd, ed string) (time.Time, time.Time, error) {

	// Date time layout
	// startDate.Format("2006-01-02 15:04:05")
	// endDate.Format("2006-01-02 15:04:05")
	dateLayout := "2006-01-02"
	startDate, err := time.Parse(dateLayout, sd)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	// zero time.Time value is time.Time{s}
	endDate, err := time.Parse(dateLayout, ed)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return startDate, endDate, nil
}
