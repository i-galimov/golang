package main

import (
"fmt"
)

func main() {
    var name string
    defer goodbye()
    fmt.Println("Привет, меня зовут Goша, а тебя как?")
    fmt.Scan(&name)
    Hi(name)
    phrase2, len_name := digname(name)
    fmt.Println(phrase2)
    for i := 0; i < len_name; i++ {
        fmt.Println(name[i:])
    }
    
}

func Hi (name string) {
    var phrase string = "Я очень рад тебя видеть, " + name + "!"
    for i := 0; i < len(phrase); i++ {
    slice_phrase := phrase[:i]
    fmt.Println(slice_phrase)
    }
}

func digname (name string) (string , int) {
    var len_name int = len(name)
    var phrase2 string = name + "- красивое имя!"
    return phrase2, len_name
}

func goodbye() {
  fmt.Println("Заходи ещё!")
}
