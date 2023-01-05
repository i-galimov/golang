// По данному трехзначному числу определите, все ли его цифры различны.

package main

import "fmt"

func main() {
  var num, a, b, c int
  fmt.Scan(&num)
  a = num % 10 
  b = (num / 10) % 10 
  c = num / 100
  if a != b && c != b && a != c {
    fmt.Println("YES")
  } else {
    fmt.Println("NO")
  }
}
