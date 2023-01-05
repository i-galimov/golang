// Дано неотрицательное целое число. Найдите и выведите первую цифру числа. 

package main

import (
  "fmt"
  "strconv"
)  

func main() {
  var n int
  fmt.Scan(&n)
  a := strconv.Itoa(n)
  fmt.Printf(string(a[0]))
}
