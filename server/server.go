// Package server runs a REST Server that is used to respond to HTTP Requests.
package server

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jmross14/TemperatureSensor/temperaturesensor"
)

// App contains an http server for a temperature sensor.
type App struct {
	router *mux.Router
	temperatureSensor *temperaturesensor.TemperatureSensor
}

// Initialize initializes member variables and routes.
func (app *App) Initialize(){
	app.router = mux.NewRouter()
	app.temperatureSensor = temperaturesensor.StartActor()

	app.getRoutes()

}

// Run starts the http server.
func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.router))
}

// getReading Sends a message to the temperature sensor actor to get the current
// temperature in Celcius.
func (app *App) getReading(w http.ResponseWriter, r *http.Request) {
	reading := app.temperatureSensor.GetTempReading()
	respondWithJSON(w, http.StatusOK, reading)
}

// respondWithJSON creates the JSON response that will be sent back in the response.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}