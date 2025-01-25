package lib

import "time"

func ParseDates(startDate, endDate string) (time.Time, time.Time, error) {
	parsedStartDate, err := time.Parse("2001-01-01", startDate)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	parsedEndDate, err := time.Parse("2001-01-01", endDate)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return parsedStartDate, parsedEndDate, nil
}
