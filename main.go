package main

import (
	"log"

	"github.com/derekkenney/weather-report/app"
)

// main is the entry point for the weather-report application
func main() {
	log.Println("Main() called")
	app.StartApp()
}
