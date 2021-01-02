package main

import (
	"time"
	"fmt"
	"github.com/MichaelS11/go-dht"
	"github.com/jmross14/sensorPlayGround/temperaturesensor"
)

func main() {
	// fmt.Println("Hello World!");
	// err := dht.HostInit()
	// if err != nil {
	// 	fmt.Println("HostInit error: ", err)
	// 	return
	// }

	// dht, err := dht.NewDHT("GPIO4", dht.Fahrenheit, "")
	// if err != nil {
	// 	fmt.Println("NewDHT error: ", err)
	// 	return
	// }
	// ticker := time.NewTicker(5 * time.Second)

	// for {
	// 	select {
	// 	case <- ticker.C:
	// 		humidity, temperature, err := dht.Read()
	// 		if err != nil {
	// 			fmt.Println("Read error: ", err)
	// 		} else {
	// 			fmt.Printf("humidity: %v\n", humidity)
	// 			fmt.Printf("temperature: %v\n", temperature)
	// 		}
	// 	}
	// }
	
	Hello()
}