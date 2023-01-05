package main

import (
	"fmt"
	_ "math"
	"sort"
	_ "strconv"
	_ "strings"
	_ "unicode"
)

func main() {
	fmt.Println("Hi")
}

func aver(Num []int) (average float64) {
	sum := 0
	count := 0
	for _, value := range Num {
		sum += value
		count++
	}
	average = float64(sum) / float64(count)
	return
}

func med(Num []int) (mediana float64) {
	sort.Ints(Num)
	if len(Num)%2 == 0 {
		mediana = float64(Num[len(Num)/2-1])/2 + float64(Num[len(Num)/2])/2
	} else if len(Num)%2 != 0 {
		mediana = float64(Num[len(Num)/2])
	}
	return
}

func moda(Num []int) (moda int) {
	var repeat int
	var count_repeat int
	for i, numbers := range Num {
		if i != 0 && Num[i] == Num[i-1] {
			repeat++
			if repeat >= count_repeat {
				count_repeat = repeat
				moda = numbers
			}
		} else if i != 0 && Num[i] != Num[i-1] {
			repeat = 1
		}
	}
	return
}
