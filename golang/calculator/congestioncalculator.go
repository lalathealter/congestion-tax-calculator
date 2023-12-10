package calculator

import (
	"errors"
	"strings"
	"time"
)

type TollRuleSet interface {
	getTollFee(time.Time) int
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

func (fees TollFeeMap) findAmount(t time.Time) int {
	var amount int

	for _, inData := range fees {
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
	}
	return amount
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

type TollRuleSetWithMaxTaxI interface {
	GetMaxTax() int
}

func ClampTax(tax int, trs TollRuleSetWithMaxTaxI) int {
	max := trs.GetMaxTax()
	if tax > max {
		tax = max
	}
	return tax
}

type TollRuleSetWithTimeSpanningI interface {
	GetGroupingTimeSpan() time.Duration
	ConcludeDatesIntoOne([]*time.Time) *time.Time
}

func GroupByTimeSpan(dates []time.Time, trs TollRuleSetWithTimeSpanningI) []time.Time {
	dur := trs.GetGroupingTimeSpan()
	coll := make([]*time.Time, 1, len(dates))
	coll[0] = &dates[0]
	currStart := dates[0]
	lastUnion := 0

	for i := 1; i < len(dates); i++ {
		el := dates[i]
		diff := el.Sub(currStart)

		if diff >= dur {
			currStart = el
			res := trs.ConcludeDatesIntoOne(coll)
			dates[lastUnion] = *res
			lastUnion++
			coll = nil
		}
		coll = append(coll, &el)
	}

	if coll != nil {
		dates[lastUnion] = *trs.ConcludeDatesIntoOne(coll)
	}

	return dates[:lastUnion+1]
}
