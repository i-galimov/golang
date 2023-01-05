// Последовательность состоит из натуральных чисел и завершается числом 0. Определите количество элементов этой последовательности, которые равны ее наибольшему элементу.

package main

import "fmt"

func main() {
  count := 1
  var n int
  fmt.Scan(&n)
  bigger := n
  for n != 0 {
  if bigger < n {
    count = 1
    bigger = n
    fmt.Scan(&n)
  } else if bigger == n {
    count++
    fmt.Scan(&n)
  } else {
    fmt.Scan(&n) }
  }
  fmt.Println(count)
}
