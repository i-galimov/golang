package main

import (
	"fmt"
	"log"
	"net/http"
	_ "strconv"
)

func main() {
	http.HandleFunc("/", sayhello)           // Устанавливаем роутер
	err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет!")
}
