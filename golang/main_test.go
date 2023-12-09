package main

import (
	"bytes"
	"congestion-calculator/controllers"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGothenburgGetTax(t *testing.T) {
	server := SetupServer()
	w := httptest.NewRecorder()

	uri := "/congestion-calculator/gothenburg"

	for _, samp := range testSamples {
		jsonBody, err := json.Marshal(samp.Data)
		if err != nil {
			assert.Fail(t, "failed constructing a http request")
		}

		req := httptest.NewRequest("POST", uri, bytes.NewBuffer(jsonBody))
		server.ServeHTTP(w, req)

		got := w.Body.String()
		assert.Equal(t, samp.Want, got)
	}
}

type TaxTestSamp struct {
	Want string
	Data controllers.CongestionTaxInput
}

var testSamples = []TaxTestSamp{
	{"31", controllers.CongestionTaxInput{
		Intervals: []string{"2013-02-08 14:35:00",
			"2013-02-08 15:29:00",
			"2013-02-08 15:47:00",
			"2013-02-08 16:01:00",
		}, Vehicle: controllers.VehicleModel{
			Type: "", Data: struct{}{},
		}},
	},
}
