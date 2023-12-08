package calculator

import (
	"errors"
	"strings"
	"time"
)

type TollRuleSet interface {
	getTollFee(time.Time, Vehicle) int
	GetTax(Vehicle, []time.Time) int
}

type TollRuleSetCollection map[string]TollRuleSet

var GlobalTollRSColl = TollRuleSetCollection{
	"gothenburg": GothenburgRuleSetInst,
}

var ErrNoTollRuleSetFound = errors.New("No Toll Rule Set was found for this location")

func GetRuleSetIn(location string) (TollRuleSet, error) {
	rs, ok := GlobalTollRSColl[strings.ToLower(location)]
	if !ok {
		return nil, ErrNoTollRuleSetFound
	}
	return rs, nil
}

type TollFreeVehicles map[VehicleType]bool

func (tfv TollFreeVehicles) isTollFreeVehicle(v Vehicle) bool {
	if v == nil {
		return false
	}
	vehtype := v.getVehicleType()
	isTollFree, _ := tfv[vehtype]

	return isTollFree
}

type TollFeeMap [][3]int // interval's hour, minute and tax amount

func (fees TollFeeMap) FindAmount(t time.Time) int {
	var amount int

	i := 1
	for cycles := 0; cycles <= len(fees); cycles++ {
		inData := fees[i]
		inhour, inminute, tax := inData[0], inData[1], inData[2]
		interval := time.Date(
			t.Year(), t.Month(), t.Day(),
			inhour, inminute,
			t.Second(), t.Nanosecond(), t.Location(),
		)

		if interval.After(t) {
			break
		}
		amount = tax
		i = abound(i, len(fees))
	}
	return amount
}

func abound(n, length int) int {
	n++
	if n > length {
		return 0
	}
	return n
}

type TollFreeDates map[[2]int]bool // month, day
const AllDaysOfMonth = -1

func (tfd TollFreeDates) isFreeDate(date time.Time) bool {
	month := date.Month()
	day := date.Day()
	checkPair := [2]int{int(month), day}
	isFreeDate, _ := tfd[checkPair]
	if !isFreeDate {
		checkPair[1] = AllDaysOfMonth
		isFreeMonth, _ := tfd[checkPair]
		isFreeDate = isFreeMonth
	}
	return isFreeDate
}

type TollFreeWeekDays [7]bool // starting from Sunday (in accordance with golang's time package)

func (tfw TollFreeWeekDays) isFreeWeekDay(date time.Time) bool {
	return tfw[date.Weekday()]
}
