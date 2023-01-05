// Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.

package main

import "fmt"

func main() {
  var num, a, b, c, d, e, f int
  fmt.Scan(&num)
  a = num % 10
  num /= 10
  b = num % 10
  num /= 10
  c = num % 10
  num /= 10
  d = num % 10
  num /= 10
  e = num % 10
  num /= 10
  f = num % 10
  if a + b + c == d + e + f {
  fmt.Println("YES")
  } else {
  fmt.Println("NO")
  }
}
