package main

import "testing"

func TestNum(t *testing.T) {
	// Arrange
	Num := []int{1, 2, -1, 0, 4, 11, 4}
	expected_average := 3
	expected_mediana := 2
	expected_moda := 4
	// Act
	res1 := aver(Num)
	res2 := med(Num)
	res3 := moda(Num)
	// Assert
	if res1 != float64(expected_average) {
		t.Errorf("Ошибка среднего. Expect %v, got %v", expected_average, res1)
	}
	if res2 != float64(expected_mediana) {
		t.Errorf("Ошибка среднего. Expect %v, got %v", expected_mediana, res2)
	}
	if res3 != expected_moda {
		t.Errorf("Ошибка среднего. Expect %d, got %d", expected_moda, res3)
	}
}
