package main

import (
	"log"

	"github.com/derekkenney/weather-report/app"
)

// main is the entry point for the weather-report application
func main() {
	log.Println("Main() called")
	err := app.StartApp()

	if err != nil {
		log.Printf("An error occurred starting the HTTP server ...%v", err)
		return
	}

}
