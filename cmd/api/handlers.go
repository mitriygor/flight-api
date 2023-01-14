package main

import (
	"net/http"
	"strings"
)

type JSONPayload struct {
	Path   []string `json:"path"`
	Output string   `json:"output"`
}

type Info struct {
	Destination []string   `json:"destination"`
	Itinerary   [][]string `json:"itinerary"`
	Count       int        `json:"count"`
}

func (app *Config) destination(w http.ResponseWriter, transfers [][]string) {
	var payload jsonResponse
	payload.Error = false
	payload.Message = "Destinations"
	payload.Data = GetDestination(transfers)

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) itinerary(w http.ResponseWriter, transfers [][]string) {
	var payload jsonResponse
	payload.Error = false
	payload.Message = "Itinerary"
	payload.Data = GetItinerary(transfers)

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) count(w http.ResponseWriter, transfers [][]string) {
	var payload jsonResponse
	payload.Error = false
	payload.Message = "Count"
	payload.Data = GetCount(transfers)

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) info(w http.ResponseWriter, transfers [][]string) {
	var payload jsonResponse
	payload.Error = false
	payload.Message = "Info"
	payload.Data = Info{GetDestination(transfers),
		GetItinerary(transfers),
		GetCount(transfers)}

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) Home(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Flight API",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) Calc(w http.ResponseWriter, r *http.Request) {
	var payload jsonResponse
	var requestPayload JSONPayload
	var transfers [][]string

	q := r.URL.Query()

	if _, ok := q["path"]; !ok {
		payload.Error = false
		payload.Message = "Calc API"
		app.writeJSON(w, http.StatusOK, payload)
		return
	}

	requestPayload.Path = strings.Split(q["path"][0], ",")

	if _, ok := q["output"]; ok {
		requestPayload.Output = q["output"][0]
	}

	transfers = GetTransfers(requestPayload.Path)

	switch requestPayload.Output {
	case "count":
		app.count(w, transfers)
	case "destination":
		app.destination(w, transfers)
	case "itinerary":
		app.itinerary(w, transfers)
	default:
		app.info(w, transfers)
	}
}
