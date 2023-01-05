# Проекты на Golang

### 
> *Путеводитель по проектам.*
* [**Погодный сервис**](https://github.com/i-galimov/golang/blob/main/WeatherApi/serverWeatherApi.go)
* [**Текстовая игра с генерацией уровней**](https://github.com/i-galimov/golang/blob/main/projects/detective_game.go)
* [**Генератор паролей**](https://github.com/i-galimov/golang/blob/main/projects/code_generator_goroutines.go)
* [**Бот в Телеграм**](https://github.com/i-galimov/golang/blob/main/projects/rss_habrbot.go)
---
### [Пример кода](https://github.com/i-galimov/golang/blob/main/projects/server_FIRST.go)
> *Это не баг, а фича*
```
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayhello)           // Устанавливаем роутер
	err := http.ListenAndServe(":8080", nil) // Устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет!")
}
```
********
![Golang](https://i0.wp.com/owlweb.ru/wp-content/uploads/2020/04/jazyk-programmirovanija-go-min.jpg?fit=600%2C400&ssl=1)
