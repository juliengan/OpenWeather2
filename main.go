package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	owm "github.com/briandowns/openweathermap"
)

var apiKey = os.Getenv("API_KEY")
var lat = os.Getenv("LATI")
var long = os.Getenv("LONGI")

func main() {
	fmt.Println(apiKey)
	w, err := owm.NewCurrent("F", "EN", apiKey) // (internal - OpenWeatherMap reference for Farenheit) with English output
	lon, err := strconv.ParseFloat(long, 64)
	latitude, err := strconv.ParseFloat(lat, 64)

	if err != nil {
		log.Fatalln(err)
	}
	w.CurrentByCoordinates(
		&owm.Coordinates{
			Longitude: lon,
			Latitude:  latitude,
		},
	)
	if err != nil {
		log.Fatalln(err)
	}
	owm.ValidAPIKey(apiKey)

	fmt.Println(w)

}
