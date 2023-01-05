package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

// Переменная для работы с базой данных
var DB *sql.DB

// Константы для работы со сторонним погодным API
const currentWeather = "https://api.open-meteo.com/v1/forecast?latitude="
const geoResponse = "http://api.openweathermap.org/geo/1.0/direct?q="
const lonWeather = "&longitude="
const endResp = "&timezone=auto&daily=temperature_2m_max"
const apiKey = "&limit=5&appid=f7047650c6b759c44bb67724cbb2ccad"

// Структура для чтения данных json, переданных на наш сервер, через переменную city
type bodyRequest struct {
	CityName string `json:"city"`
}

// Структура для чтения гео данных json, переданных через API стороннего сервиса
type Geo struct {
	Country string  `json:"country"`
	State   string  `json:"state,omitempty"`
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

// Переменные для работы с погодными данными нашего сервиса
var jsonWeaherStr string
var jsonAnswer string
var jsonStr string
var latStr string
var lonStr string
var city string
var Day1 float64
var Day2 float64
var Day3 float64
var Lat float64
var Lon float64

// Структура для чтения погодных данных json, переданных через API стороннего сервиса
type WeatherJson struct {
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	DailyUnits           struct {
		Longitude        float64 `json:"longitude"`
		Elevation        float64 `json:"elevation"`
		Latitude         float64 `json:"latitude"`
		Timezone         string  `json:"timezone"`
		Time             string  `json:"time"`
		Temperature2MMax string  `json:"temperature_2m_max"`
	} `json:"daily_units"`
	Daily struct {
		Time             []string  `json:"time"`
		Temperature2MMax []float64 `json:"temperature_2m_max"`
	} `json:"daily"`
}

func main() {
	// Данные для авторизации в БД безопаснее хранить в ENV в переменной KEY
	connStr := "user=postgres password=sql999 dbname=weather sslmode=disable"
	var err error
	// PostgreSQL подключение к БД
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer DB.Close()
	// Соединение с сервером
	http.HandleFunc("/", SayCity)           // Устанавливаем роутер
	err = http.ListenAndServe(":8080", nil) // Устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Функция-обработчик запросов для нашего сервера
func SayCity(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != "POST" {
		fmt.Fprintf(w, "Привет, перед тобой сервис Погоды. Чтобы все заработало - нужно отправить запрос методом Post c переменной city!")
		return
	}
	for {
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "500")
			return
		}
		var cityReq bodyRequest
		err = json.Unmarshal(bytes, &cityReq)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		city = cityReq.CityName
		MakeRequest(city)
		data := []byte(jsonStr)
		u := make([]Geo, 0)
		err2 := json.Unmarshal(data, &u)
		if err2 != nil {
			fmt.Fprintf(w, "We can not get json Geo")
		}
		for _, info := range u {
			if strings.Compare(city, info.Name) == 0 {
				Lat = info.Lat
				Lon = info.Lon
				latStr = fmt.Sprintf("%f", Lat)
				lonStr = fmt.Sprintf("%f", Lon)
				WeatherRequest()
				Day1Str := fmt.Sprintf("%f", Day1)
				Day2Str := fmt.Sprintf("%f", Day2)
				Day3Str := fmt.Sprintf("%f", Day3)
				jsonAnswer = "city: " + city + " lat: " + latStr + " lon: " + lonStr + " today: " + Day1Str + " tomorrow: " + Day2Str + " after_tomorrow: " + Day3Str
				jsonServer, _ := json.Marshal(jsonAnswer)
				w.Write(jsonServer)
				// Добавляем название города, широту и долготу в БД Postgresql
				_, err3 := DB.Exec("INSERT into weather (city, lat, lon) values ($1, $2, $3)", city, Lat, Lon)
				if err3 != nil {
					fmt.Fprint(w, err.Error())
					return
				} else {
					log.Printf("добавлен город: %s, %f, %f\n", city, Lat, Lon)
				}
				break
			}
		}
	}

}

// Функция получает географические координаты места по названию
func MakeRequest(city string) {
	resp, err := http.Get(geoResponse + city + apiKey)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	jsonStr = string(body)
	defer resp.Body.Close()

}

// Функция получает данные погоды по координатам
func WeatherRequest() {
	var ResGet = currentWeather + latStr + lonWeather + lonStr + endResp
	resp, err := http.Get(ResGet)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	jsonWeaherStr = string(body)
	defer resp.Body.Close()
	dataWeather := []byte(jsonWeaherStr)
	var u2 WeatherJson
	err2 := json.Unmarshal(dataWeather, &u2)
	if err2 != nil {
		fmt.Println(err.Error())
		return
	}
	Day1 = u2.Daily.Temperature2MMax[0]
	Day2 = u2.Daily.Temperature2MMax[1]
	Day3 = u2.Daily.Temperature2MMax[2]
}
