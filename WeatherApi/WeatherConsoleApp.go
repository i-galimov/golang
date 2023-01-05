package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var city string
var apiKey string = "&limit=5&appid=f7047650c6b759c44bb67724cbb2ccad"
var GeoResponse string = "http://api.openweathermap.org/geo/1.0/direct?q="
var jsonStr string

var CurrentWeather = "https://api.open-meteo.com/v1/forecast?latitude="
var LonWeather = "&longitude="
var endResp = "&timezone=auto&daily=temperature_2m_max"

var Lat float64
var Lon float64
var jsonWeaherStr string
var latStr string
var lonStr string

type Geo struct {
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
	State   string  `json:"state,omitempty"`
}

type WeatherJson struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	DailyUnits           struct {
		Time             string `json:"time"`
		Temperature2MMax string `json:"temperature_2m_max"`
	} `json:"daily_units"`
	Daily struct {
		Time             []string  `json:"time"`
		Temperature2MMax []float64 `json:"temperature_2m_max"`
	} `json:"daily"`
}

func main() {
	fmt.Println("Введите название города на английском, чтобы узнать погоду:")
	for {
		fmt.Scan(&city)
		MakeRequest(city)
		data := []byte(jsonStr)
		u := make([]Geo, 0)
		err := json.Unmarshal(data, &u)
		if err != nil {
			fmt.Println("We can not get json Geo")
			fmt.Println(err.Error())
		}
		for _, info := range u {
			if strings.Compare(city, info.Name) == 0 {
				fmt.Println("Город:", city)
				fmt.Println("Широта:", info.Lat, "\nДолгота:", info.Lon)
				Lat = info.Lat
				Lon = info.Lon
				latStr = fmt.Sprintf("%f", Lat)
				lonStr = fmt.Sprintf("%f", Lon)
				WeatherRequest()
				break
			}
		}
	}
}

func MakeRequest(city string) {
	resp, err := http.Get(GeoResponse + city + apiKey)
	if err != nil {
		fmt.Println("We can not get json Geo")
		fmt.Println(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("We can not get json Geo")
		fmt.Println(err.Error())
	}
	jsonStr = string(body)
}

func WeatherRequest() {
	var ResGet = CurrentWeather + latStr + LonWeather + lonStr + endResp
	resp, err := http.Get(ResGet)
	if err != nil {
		fmt.Println("We can not get Current Weather")
		fmt.Println(err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("We can not read Current Weather")
		fmt.Println(err.Error())
	}
	jsonWeaherStr = string(body)

	dataWeather := []byte(jsonWeaherStr)
	var u2 WeatherJson
	err2 := json.Unmarshal(dataWeather, &u2)
	if err2 != nil {
		fmt.Println("We can not get json Weather")
		fmt.Println(err2.Error())
	}
	fmt.Println("Сегодня до", u2.Daily.Temperature2MMax[0], "градусов по Цельсию.")
	fmt.Println("Завтра до", u2.Daily.Temperature2MMax[1], "градусов по Цельсию.")
	fmt.Println("Послезавтра до", u2.Daily.Temperature2MMax[2], "градусов по Цельсию.")
	defer resp.Body.Close()
}
