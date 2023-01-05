// Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных числа A и B (каждое не более 100, A < B). Вывести сумму всех чисел от A до B  включительно.

package main

import "fmt"

func main() {
  var A, B, sum int
  fmt.Scan(&A)
  fmt.Scan(&B)
  for i := A; i <= B; i++ {
  sum += i
  }
  fmt.Println(sum)
}
