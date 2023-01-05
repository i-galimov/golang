package main

import ("fmt"; "math/rand"; "time")

func main() {
rand.Seed(time.Now().UTC().UnixNano()) 
var digit int = rand.Intn(10)
fmt.Println(digit)
}
