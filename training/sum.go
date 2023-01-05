package main

import "fmt"

func main() {
    var one = 5
    var two int = 4
    
    fmt.Println(sum(one, two))

}

func sum (a int, b int) int {
    return a + b
}
