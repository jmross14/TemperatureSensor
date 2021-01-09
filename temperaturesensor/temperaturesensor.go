// Package temperaturesensor is used to define a temperature sensor.
package temperaturesensor

import (
	"time"
	"github.com/jmross14/go-smarthome-common"
	"github.com/MichaelS11/go-dht"
)

//action is a type define for a function. Used in messaging.
type action func()

// TemperatureSensor represents a temperature sensor actor
type TemperatureSensor struct{
	// actionChan is the channel messages are sent to the actor loop on.
	actionChan chan action
	// Reading contains a reading from the sensor
	reading gosmarthomecommon.TemperatureReading
	// dht reads from the sensor
	dht *dht.DHT
}

// StartActor starts the actor and starts the action chan
func StartActor() *TemperatureSensor {
	ch := make(chan action, 16)
	
	dht.HostInit()
	dht, _ := dht.NewDHT("GPIO4", dht.Celsius, "")

	humidity, temperature, _ := dht.ReadRetry(15)
	sensor := TemperatureSensor{
		ch,
		gosmarthomecommon.TemperatureReading{Humidity: humidity, Temperature: temperature},
		dht,
	}
	go sensor.actorLoop(sensor.actionChan)
	return &sensor
}

// GetTempReading gets the current temperature stored in the temperature sensor struct
func (temperatureSensor *TemperatureSensor) GetTempReading() gosmarthomecommon.TemperatureReading {
	ch := make(chan gosmarthomecommon.TemperatureReading, 1)
	temperatureSensor.actionChan <- func() {
		ch <- temperatureSensor.reading
	}

	return <-ch
}

// actorLoop runs in the background to process incoming messages and to fill
// the temperature reading of the temperature sensor struct
func (temperatureSensor *TemperatureSensor) actorLoop(ch <-chan action) {
	//Used to take the temperature every 10 seconds due to dht.Read() failing sometimes.
	//Do not want to block for the 10-12 seconds for dht.ReadRetry().
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case action := <-ch:
			action()
		case <-ticker.C:
			temperatureSensor.takeTemperatureReading()
		}		
	}
} 

// takeTemperature sets reading in the temperature sensor struct from the physical
// sensor
func (temperatureSensor *TemperatureSensor) takeTemperatureReading() {
	
	humidity, temperature, err := temperatureSensor.dht.Read()
	if err == nil {
		temperatureSensor.reading = gosmarthomecommon.TemperatureReading{Humidity: humidity, Temperature: temperature}
	} 
}


