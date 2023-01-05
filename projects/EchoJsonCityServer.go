package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	_ "strconv"
	_ "strings"
)

type bodyRequest struct {
	CityName string `json:"city"`
}

func main() {
	http.HandleFunc("/", saycity)            // Устанавливаем роутер
	err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func saycity(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != "POST" {
		fmt.Fprintf(w, "Привет, чтобы все заработало - нужно отправить запрос методом Post c переменной city!")
		return
	}
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "500")
		return
	}

	var cityReq bodyRequest

	err = json.Unmarshal(bytes, &cityReq)
	if err != nil {
		fmt.Fprintf(w, "400")
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, cityReq.CityName)
}
