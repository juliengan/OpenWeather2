package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	owm "github.com/briandowns/openweathermap"
	"github.com/gorilla/mux"
)

type coord struct {
	Long float64
	Lat  float64
}
type weather struct {
	Id          int
	Main        string
	Description string
	Icon        string
}

type mainn struct {
	Temp       float64
	Feels_like float64
	Temp_min   float64
	Temp_max   float64
	Pressure   int
	Humidity   int
}

type wind struct {
	Speed float64
	Deg   int
}

type sys struct {
	Typee   int
	Id      int
	Message float32
	Country string
	Sunrise int
	Sunset  int
}

type Weathers struct {
	Coord      coord
	Weather    weather
	Base       string
	Main       mainn
	Visibility int
	Wind       wind
	Clouds     int
	Dt         int
	Sys        sys
	Timezone   int
	Id         int
	Name       string
	Cod        int
}

func (m weather) MarshalJSON2() ([]byte, error) {
	j, err := json.Marshal(struct {
		Id          int
		Main        string
		Description string
		Icon        string
	}{
		Id:          m.Id,
		Main:        m.Main,
		Description: m.Description,
		Icon:        m.Icon,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (m wind) MarshalJSON3() ([]byte, error) {
	j, err := json.Marshal(struct {
		Speed float64
		Deg   int
	}{
		Speed: m.Speed,
		Deg:   m.Deg,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (m sys) MarshalJSON4() ([]byte, error) {
	j, err := json.Marshal(struct {
		Typee   int
		Id      int
		Message float32
		Country string
		Sunrise int
		Sunset  int
	}{
		Typee:   m.Typee,
		Id:      m.Id,
		Message: m.Message,
		Country: m.Country,
		Sunrise: m.Sunrise,
		Sunset:  m.Sunset,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (m mainn) MarshalJSON5() ([]byte, error) {
	j, err := json.Marshal(struct {
		Temp       float64
		Feels_like float64
		Temp_min   float64
		Temp_max   float64
		Pressure   int
		Humidity   int
	}{
		Temp:       m.Temp,
		Feels_like: m.Feels_like,
		Temp_min:   m.Temp_min,
		Temp_max:   m.Temp_max,
		Pressure:   m.Pressure,
		Humidity:   m.Humidity,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (m coord) MarshalJSON6() ([]byte, error) {
	j, err := json.Marshal(struct {
		Long float64
		Lat  float64
	}{
		Long: m.Long,
		Lat:  m.Lat,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (m Weathers) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Coord      coord
		Weather    weather
		Base       string
		Main       mainn
		Visibility int
		Wind       wind
		Clouds     int
		Dt         int
		Sys        sys
		Timezone   int
		Id         int
		Name       string
		Cod        int
	}{
		Coord:      m.Coord,
		Weather:    m.Weather,
		Base:       m.Base,
		Main:       m.Main,
		Visibility: m.Visibility,
		Wind:       m.Wind,
		Clouds:     m.Clouds,
		Dt:         m.Dt,
		Sys:        m.Sys,
		Timezone:   m.Timezone,
		Id:         m.Id,
		Name:       m.Name,
		Cod:        m.Cod,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

// Weather data for long and lat desired here
var weatherr = Weathers{
	Coord:      coord{102.7542, 5.9028}, // coordinates
	Sys:        sys{0, 0, 0, "MY", 1655160925, 1655205783},
	Base:       "stations",
	Weather:    weather{802, "Clouds scattered", "clouds", "03n"},
	Main:       mainn{77.94, 79.2, 1007, 1007, 980, 80},
	Visibility: 10000,
	Wind:       wind{6.4, 187},
	Clouds:     44,
	Dt:         1655158855,
	Id:         1736405,
	Name:       "Jertih",
	Cod:        200,
	Timezone:   28800,
}

// GET
// /?lat=5.902785&lon=102.754175
func getWeather(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getWeatherPage")

	var response Weathers
	weathers := weatherr
	response = weathers

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Something went wrong")
		return
	}

	w.Write(jsonResponse)

	// Retrieve OWM info : TP1
	var apiKey = os.Getenv("API_KEY")
	own, err := owm.NewCurrent("F", "EN", apiKey)

	if err != nil {
		log.Fatalln(err)
	}

	own.CurrentByCoordinates(
		&owm.Coordinates{
			Longitude: 102.754175,
			Latitude:  5.902785,
		},
	)
	fmt.Println(own)

	//var responsetwo Weathers
	//weatherstwo := own
	//responsetwo = weatherstwo
	//w.Write(responsetwo)
}

func main() {
	// TP2 : creation of API
	router := mux.NewRouter()
	router.HandleFunc("/weather", getWeather).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))

}
