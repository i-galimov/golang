// Дано натуральное число, выведите его последнюю цифру.

package main

import "fmt"

func main() {
  var N int
  fmt.Scan(&N)
  N %= 10
  fmt.Println(N)
}
