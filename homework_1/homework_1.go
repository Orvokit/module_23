//Напишите функцию, сортирующую массив длины 10 вставками.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 10

func create() [size]int {
	var arr [size]int
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(100)
	}

	return arr
}

func insertSort(arr [size]int) [size]int {
	for i := 1; i < size; i++ {
		key := arr[i]
		j := i
		for (j >= 1) && (arr[j-1] > key) {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
		arr[j] = key
	}
	return arr
}

func main() {
	arr := create()
	fmt.Println("Сгенерированный массив", arr)
	sortedArr := insertSort(arr)
	fmt.Println("Отсортированный массив", sortedArr)

}
