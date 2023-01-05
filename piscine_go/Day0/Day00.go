package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	_ "strings"
	_ "unicode"
)

func main() {
	//медиана и среднее
	var num string
	var count int
	var sum float64
	var mediana float64
	var average float64
	//стандартное отклонение
	var st_dev float64
	var sum_dev float64
	// мода
	var count_repeat int
	var moda int
	repeat := 1
	count_repeat = 1
	//слайс для ввода данных
	a := []int{}
	for {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic occurred:", err)
			}
		}()
		fmt.Scan(&num)
		if num == "0" {
			a = append(a, 0)
			count++
		}
		num_b, _ := strconv.Atoi(num)
		if num_b != 0 {
			a = append(a, num_b)
			sort.Ints(a)
			count++
			sum += float64(num_b)
			average = sum / float64(count)
			if len(a) == 1 {
				moda = a[0]
			}
			if len(a) != 0 {
				if count%2 == 0 {
					mediana = float64(a[count/2-1])/2 + float64(a[count/2])/2
				} else if count%2 != 0 {
					mediana = float64(a[count/2])
				}
				for i, numbers := range a {
					dev := math.Pow((float64(numbers) - average), 2)
					sum_dev += dev
					if i != 0 && a[i] == a[i-1] {
						repeat++
						if repeat >= count_repeat {
							count_repeat = repeat
							moda = numbers
						}
					} else if i != 0 && a[i] != a[i-1] {
						repeat = 1
					}
				}
				st_dev = math.Sqrt(sum_dev)
				fmt.Println(a)
				fmt.Println("Среднее равно", average)
				fmt.Println("Медиана равна", mediana)
				fmt.Println("Стандартное отклонение равно", st_dev)
				fmt.Println("Мода равна", moda)
			}
		}
	}
}
