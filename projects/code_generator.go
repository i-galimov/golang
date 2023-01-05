package main

import ("fmt"; "math/rand"; "time")

func init() {
    fmt.Println("Привет, эта программа сгенерирует уникальный 10 значный код!")
    fmt.Println()
}

func random_digit(n int) int {
rand.Seed(time.Now().UTC().UnixNano()) 
var digit int = rand.Intn(n)
return digit
}

func random_str() string {
    var value int
    value = random_digit(89) + 33
    return string(value)
}

func main() {
    var a string
    for i:= 0; i < 10; i++ {
        a += random_str()
    }
    fmt.Println(a)
}
