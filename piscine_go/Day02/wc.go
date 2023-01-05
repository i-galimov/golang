package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var nameFile string
	var textFile string
	fmt.Println("Введите название файла, который хотите считать:")
	fmt.Scan(&nameFile)

	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path) // for example /home/user

	file, err := os.Open(path + "/" + nameFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)
	fmt.Println("Вывожу содержимое файла:")
	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		textFile += string(data[:n])
	}
	fmt.Println(textFile)
	fmt.Println("\nКоличество слов:")
	words := strings.Split(textFile, " ")
	fmt.Println(len(words) + 1)
}
