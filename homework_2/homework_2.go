//Напишите анонимную функцию, которая на вход получает массив типа integer,
//сортирует его пузырьком и переворачивает (либо сразу сортирует в обратном порядке, как посчитаете нужным).

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 10

func main() {
	var arr [size]int
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("Сгенерированный массив", arr)

	FUNC := func(arr [size]int) [size]int {
		for i := 0; i < size; i++ {
			for j := 0; j < size-i-1; j++ {
				if arr[j] < arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
		return arr
	}

	fmt.Println("Отсортированный и перевернутый массив", FUNC(arr))

}
