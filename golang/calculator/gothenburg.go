package calculator

import (
	"time"
)

var GothenburgRuleSetInst = GothenburgRuleSet{
	GothenburgTollFeeByTime,
	GothenburgTollFreeVehicles,
	GothenburgTollFreeDates,
	GothenburgTollFreeWeekDays,
}

type GothenburgRuleSet struct {
	FeeIntervals TollFeeMap
	FreeVehicles TollFreeVehicles
	FreeDates    TollFreeDates
	FreeWeekDays TollFreeWeekDays
}

var GothenburgTollFreeVehicles = TollFreeVehicles{
	Emergency:  true,
	Bus:        true,
	Diplomat:   true,
	Motorcycle: true,
	Military:   true,
	Foreign:    true,
}

var GothenburgTollFreeWeekDays = TollFreeWeekDays{
	true, false, false, false, false, false, true,
} // Saturdays and Sundays

var GothenburgTollFreeDates = TollFreeDates{
	{3, 28}: true, {3, 29}: true,
	{4, 1}: true, {4, 30}: true,
	{5, 1}: true, {5, 8}: true,
	{5, 9}: true, {6, 5}: true,
	{6, 6}: true, {6, 21}: true,
	{7, AllDaysOfMonth}: true, {11, 1}: true,
	{12, 24}: true, {12, 25}: true,
	{12, 26}: true, {12, 31}: true,
	{1, 1}: true,
}

var GothenburgTollFeeByTime = TollFeeMap{
	{6, 0, 8},
	{6, 30, 13},
	{7, 0, 18},
	{8, 0, 13},
	{8, 30, 8},
	{15, 0, 13},
	{15, 30, 18},
	{17, 0, 13},
	{18, 0, 8},
	{18, 30, 8},
}

func (grs GothenburgRuleSet) isTollFreeDate(date time.Time) bool {
	if grs.FreeWeekDays.isFreeWeekDay(date) {
		return true
	}

	return grs.FreeDates.isFreeDate(date)

}

func (grs GothenburgRuleSet) getTollFee(t time.Time, v Vehicle) int {
	if grs.isTollFreeDate(t) ||
		grs.FreeVehicles.isTollFreeVehicle(v) {
		return 0
	}

	return grs.FeeIntervals.FindAmount(t)
}

func (grs GothenburgRuleSet) GetTax(vehicle Vehicle, dates []time.Time) int {
	intervalStart := dates[0]
	totalFee := 0
	for _, date := range dates {
		nextFee := grs.getTollFee(date, vehicle)
		tempFee := grs.getTollFee(date, vehicle)

		minutes := date.Sub(intervalStart).Minutes()

		if minutes <= 60 {
			if totalFee > 0 {
				totalFee = totalFee - tempFee
			}
			if nextFee >= tempFee {
				tempFee = nextFee
			}
		} else {
			totalFee = totalFee + nextFee
		}
	}

	if totalFee > 60 {
		totalFee = 60
	}
	return totalFee
}
