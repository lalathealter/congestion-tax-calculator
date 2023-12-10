package main

import (
	"bytes"
	"congestion-calculator/calculator"
	"congestion-calculator/controllers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGothenburgGetTax(t *testing.T) {
	server := SetupServer()

	uri := "/congestion-calculator/gothenburg"

	for _, samp := range testSamples {
		jsonBody, err := json.Marshal(samp.Data)
		if err != nil {
			assert.Fail(t, "failed constructing a http request")
		}

		req := httptest.NewRequest("POST", uri, bytes.NewBuffer(jsonBody))
		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)

		got := w.Body.String()
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, samp.Want, got, "SAMPLE USED:\n", samp.Data)
	}
}

type TaxTestSamp struct {
	Want string
	Data controllers.CongestionTaxInput
}

var testSamples = []TaxTestSamp{
	{"31", controllers.CongestionTaxInput{
		Intervals: []string{
			"2013-02-08 14:35:00",
			"2013-02-08 15:29:00",
			"2013-02-08 15:47:00",
			"2013-02-08 16:01:00",
		}, Vehicle: controllers.VehicleModel{
			Type: calculator.Basic, Data: struct{}{},
		}}},
	{"29", controllers.CongestionTaxInput{
		Intervals: []string{
			"2013-01-14 21:00:00",
			"2013-01-15 21:00:00",
			"2013-02-07 06:23:27",
			"2013-02-07 15:27:00",
			"2013-02-08 06:27:00",
		}, Vehicle: controllers.VehicleModel{
			Type: calculator.Tractor, Data: struct{}{},
		}}},
	{"0", controllers.CongestionTaxInput{
		Intervals: []string{
			"2013-01-14 21:00:00",
			"2013-01-15 21:00:00",
		}, Vehicle: controllers.VehicleModel{
			Type: calculator.Tractor, Data: struct{}{},
		}}},
	{"49", controllers.CongestionTaxInput{
		Intervals: []string{
			"2013-02-08 15:47:00",
			"2013-02-08 16:01:00",
			"2013-02-08 16:48:00",
			"2013-02-08 17:49:00",
			"2013-02-08 18:29:00",
		}, Vehicle: controllers.VehicleModel{
			Type: calculator.Basic, Data: struct{}{},
		}}},
	{"31", controllers.CongestionTaxInput{
		Intervals: []string{
			"2013-02-08 16:48:00",
			"2013-02-08 17:49:00",
			"2013-02-08 18:29:00",
			"2013-02-08 18:35:00",
			"2013-03-29 14:25:00",
		}, Vehicle: controllers.VehicleModel{
			Type: calculator.Basic, Data: struct{}{},
		}}},
	{"18", controllers.CongestionTaxInput{
		Intervals: []string{
			"2013-02-08 15:29:00",
			"2013-02-08 15:47:00",
			"2013-02-08 16:01:00",
		}, Vehicle: controllers.VehicleModel{
			Type: calculator.Basic, Data: struct{}{},
		}}},
	{"54", controllers.CongestionTaxInput{
		Intervals: []string{
			"2013-01-18 16:01:00",
			"2013-01-15 15:47:00",
			"2013-01-14 15:39:00",
		}, Vehicle: controllers.VehicleModel{
			Type: calculator.Basic, Data: struct{}{},
		}}},
	{"18", controllers.CongestionTaxInput{
		Intervals: []string{
			"2013-01-08 15:39:00",
			"2013-01-13 15:47:00",
			"2013-01-27 16:01:00",
		}, Vehicle: controllers.VehicleModel{
			Type: calculator.Basic, Data: struct{}{},
		}}},
	{"49", controllers.CongestionTaxInput{
		Intervals: []string{
			"2013-02-08 15:29:00",
			"2013-02-08 15:47:00",
			"2013-02-08 16:01:00",
			"2013-02-08 16:48:00",
			"2013-02-08 17:49:00",
		}, Vehicle: controllers.VehicleModel{
			Type: calculator.Basic, Data: struct{}{},
		}}},
	{"0", controllers.CongestionTaxInput{
		Intervals: []string{
			"2013-02-08 14:35:00",
			"2013-02-08 15:29:00",
			"2013-02-08 15:47:00",
			"2013-02-08 16:01:00",
		}, Vehicle: controllers.VehicleModel{
			Type: calculator.Diplomat, Data: struct{}{},
		}}},
	{"0", controllers.CongestionTaxInput{
		Intervals: []string{
			"2013-02-07 06:23:27",
			"2013-02-07 15:27:00",
			"2013-02-08 06:27:00",
			"2013-02-08 06:20:27",
			"2013-02-08 14:35:00",
		}, Vehicle: controllers.VehicleModel{
			Type: calculator.Military, Data: struct{}{},
		}}},
}
