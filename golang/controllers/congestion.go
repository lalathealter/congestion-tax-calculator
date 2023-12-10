package controllers

import (
	"congestion-calculator/calculator"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CongestionTaxInput struct {
	Intervals []string     `binding:"required"`
	Vehicle   VehicleModel `binding:"required"`
}

var ErrEmptyTimeIntervals = errors.New("No time intervals were supplied for the calculation")

const timeStringFormat = "2006-01-02 15:04:05"

func (cti CongestionTaxInput) ParseIntervals() ([]time.Time, error) {
	if len(cti.Intervals) == 0 {
		return nil, ErrEmptyTimeIntervals
	}

	timeIntervals := make([]time.Time, len(cti.Intervals))
	for i, timeStr := range cti.Intervals {
		parsed, err := time.Parse(timeStringFormat, timeStr)
		if err != nil {
			return nil, err
		}
		timeIntervals[i] = parsed
	}
	return timeIntervals, nil
}

type VehicleModel struct {
	Data struct{}
	Type calculator.VehicleType
}

func HandleCongestionCalculation(c *gin.Context) {

	location := c.Param("location")
	rules, err := calculator.GetRuleSetIn(location)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var congTaxInput CongestionTaxInput
	if err := c.BindJSON(&congTaxInput); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	timeIntervals, err := congTaxInput.ParseIntervals()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	vehicle, err := calculator.ParseVehicleType(congTaxInput.Vehicle.Type)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result := rules.GetTax(vehicle, timeIntervals)
	c.JSON(http.StatusCreated, result)
}
