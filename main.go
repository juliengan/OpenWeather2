package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	owm "github.com/briandowns/openweathermap"
)

type coord struct {
	long float32 `json:"long"`
	lat  float32 `json:"lat"`
}
type weather struct {
	id          int `json:"id"`
	main        int `json:"main"`
	description int `json:"description"`
	icon        int `json:"icon"`
}

type mainn struct {
	temp       float32 `json:"temp"`
	feels_like float32 `json:"feels_like"`
	temp_min   float32 `json:"temp_min"`
	temp_max   float32 `json:"temp_max"`
	pressure   int     `json:"pressure"`
	humidity   int     `json:"humidity"`
}

type wind struct {
	speed float32 `json:"speed"`
	deg   int     `json:"deg"`
}

type sys struct {
	typee   int     `json:"typee"`
	id      int     `json:"id"`
	message float32 `json:"message"`
	country string  `json:"country"`
	sunrise int     `json:"sunrise"`
	sunset  int     `json:"sunset"`
}

type Weather struct {
	coord      coord   `json:"coord"`
	weather    weather `json:"Title"`
	base       string  `json:"base"`
	main       mainn   `json:"main"`
	visibility int     `json:"visibility"`
	wind       wind    `json:"wind"`
	clouds     int     `json:"clouds"`
	dt         int     `json:"dt"`
	sys        sys     `json:"sys"`
	timezone   int     `json:"timezone"`
	id         int     `json:"id"`
	name       string  `json:"name"`
	cod        int     `json:"cod"`
}

type allweathers []Weather

// let's declare a global Weather array
// that we can then populate in our main function
// to simulate a database
// &{{102.7542 5.9028} {0 0 0 MY 1655160925 1655205783} stations [{802 Clouds scattered clouds 03n}] {77.94 77.94 77.94 79.2 1007 1007 980 80} 10000 {6.4 187} {44} {0 0} {0 0} 1655158855 1736405 Jertih 200 28800 imperial EN 62bd02468799bb9568074245d9b8631e 0xc000010030}

var Weathers = allweathers{}

// GET
// /?lat=5.902785&lon=102.754175
func getOneWeather(w http.ResponseWriter, r *http.Request) {
	weatherLat := mux.Vars(r)["coord"]["lat"]
	weatherLong := mux.Vars(r)["coord"]["long"]

	for _, singleWeather := range Weathers {
		if int(singleWeather.coord.lat) == int(weatherLat) && int(singleWeather.coord.long) == int(weatherLong) {
			json.NewEncoder(w).Encode(singleWeather)
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
	json.NewEncoder(w).Encode(Weathers)
}

func main() {
	// Retrieve OWM info
	var apiKey = os.Getenv("API_KEY")
	fmt.Println(apiKey)
	w, err := owm.NewCurrent("F", "EN", apiKey)

	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByCoordinates(
		&owm.Coordinates{
			Longitude: 102.754175,
			Latitude:  5.902785,
		},
	)

	owm.ValidAPIKey(apiKey)
	fmt.Println(w)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/?lat=5.902785&lon=102.754175", getOneWeather).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))

}
