package main

import "github.com/jmross14/go-smarthome-temperaturesensor/server"

func main() {	
	app := server.App{}
	//Initialize the server before starting the http server
	app.Initialize()

	//Start the http server
	app.Run(":8080")
}