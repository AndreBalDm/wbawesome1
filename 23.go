package wbawesome1

// Удалить i-ый элемент из слайса.

import "fmt"

func Wb23() {
	var i int = 5
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("Было:  %+v\n", slice)
	slice = append(slice[:i], slice[i+1:]...)
	fmt.Printf("Стало: %+v\n", slice)
}
