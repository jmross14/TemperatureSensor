// Package server runs a REST Server that is used to respond to HTTP Requests.
package server

// getRoutes adds all routes to the router
func (app *App) getRoutes() {
	//Route to get a temperature reading
	app.router.HandleFunc("/getReading", app.getReading).Methods("GET")
}