package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sorter(arr []float64) []float64 {
	newArr := make([]float64, 0)
	newArr = append(newArr, arr...)
	sort.Float64s(newArr)
	return newArr

}

func redrawer(arr []float64) {
	newArr := make([]string, 0)
	for _, val := range arr {
		checker := val - math.Ceil(val)
		if checker != 0 {
			numb := strings.Split(fmt.Sprintf("%v", val), ".")
			numb = append(numb[:1], ",", numb[1])
			newArr = append(newArr, strings.Join(numb, ""))
		} else {
			newArr = append(newArr, fmt.Sprintf("%v", val))

		}

	}

	fmt.Println("Отображение с зяпятой: ", newArr)
}

func fractionalNumbers(arr []float64) {
	newArr := make([]float64, 0)
	for _, val := range arr {
		numb := val - math.Ceil(val)
		if numb != 0 {
			newArr = append(newArr, val)
		}
	}
	fmt.Println("Дробные числа: ", newArr)
}

func reader() []float64 {
	slice := make([]float64, 0)
	var name string
chik:
	for true {
		fmt.Print("Введите элемент массива: ")
		fmt.Fscan(os.Stdin, &name)
		if name == "stop" {
			break chik
		}
		if s, err := strconv.ParseFloat(name, 5); err == nil {
			slice = append(slice, s)
		}
	}
	return slice
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

func negativer(arr []float64) {

	newArr := make([]float64, 0)
	newArr = append(newArr, sorter(arr)...)
	idx := int(math.Ceil(float64(len(newArr)/2) - 1))

	if newArr[idx] > 0 {
	flags1:
		for i, val := range newArr[:idx] {
			if val <= 0 {
				continue

			} else {
				idx = i - 1
				break flags1

			}
		}

		fmt.Println("Отричательные элементы: ", newArr[:idx])

	} else {
	flags2:
		for i, val := range newArr[idx:] {
			if val < 0 {
				continue
			} else {
				idx = i
				break flags2
			}
		}
		fmt.Println("Отричательные элементы: ", newArr[idx:])

	}

}

func main() {

	array := make([]float64, 0)
	array = append(array, reader()...)
	fmt.Println("Исходыный массив: ", array)

	fmt.Println("отсортированный массив: ", sorter(array))
	fmt.Println("отсортированный через рекурсию массив: ", quickSort(array))

	negativer(array)
	fractionalNumbers(array)
	redrawer(array)

}
