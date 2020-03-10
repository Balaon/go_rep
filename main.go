package main

import "fmt"

func main() {

	arrey := []float64{17, 3, 1, 7, 4, 5, 8, 43, 9, 10, 2, 6, 22, 21, -23, 432, 0.23, 5.5, 23.1, 10, 10}
	fmt.Println(arrey)
	println("После сортировка_____________________________")
	fmt.Println(quickSort(arrey))

}

func quickSort(arr []float64) []float64 {
	sl := make([]float64, 0)
	sl = append(sl, arr...)

	if len(sl) <= 2 {
		return sl
	} else {
		less := make([]float64, 0)
		gross := make([]float64, 0)
		fn := make([]float64, 0)
		g := 0
		var numb float64 = sl[g]
		for _, value := range sl[1:] {
			if value > numb {
				gross = append(gross, value)
			} else {
				less = append(less, value)
			}
		}

		quickSort(less)
		quickSort(gross)
		fn = append(fn, quickSort(less)...)
		fn = append(fn, numb)
		fn = append(fn, quickSort(gross)...)

		return fn

	}

}
