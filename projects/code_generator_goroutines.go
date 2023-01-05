package main

import ("fmt"; "math/rand"; "time"; "sync")

var wg sync.WaitGroup
var num int

func init() {
    fmt.Println("Привет, эта программа сгенерирует уникальный код! Сколько нужно знаков?")
    fmt.Scan(&num)
    fmt.Println()
}

func random_digit(n int) int {
rand.Seed(time.Now().UTC().UnixNano()) 
var digit int = rand.Intn(n)
return digit
}

func random_str() {
    var value int
    defer wg.Done()
    value = random_digit(89) + 33
    fmt.Printf("%s", string(value))
}

func main() {
    for i:= 0; i < num; i++ {
        go random_str()
        wg.Add(1) 
    }
    wg.Wait()
    fmt.Println()
    fmt.Println()
    fmt.Println("Напечатано", num, "символов!")
}
