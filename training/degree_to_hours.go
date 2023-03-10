/*Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m.

Входные данные

На вход программе подается целое число d (0 < d < 360).

Выходные данные

Выведите на экран фразу:

It is ... hours ... minutes.

Вместо многоточий программа должна выводить значения h и m, отделяя их от слов ровно одним пробелом. */

package main
import "fmt"

func main () {
  var d, h, m, time int
  fmt.Scan(&d)
  time = d * 2
  h = time / 60
  m = time - 60 * h
  fmt.Println("It is", h, "hours", m, "minutes.") 
}
