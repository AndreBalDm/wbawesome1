package wbawesome1

//Реализовать бинарный поиск встроенными методами языка.

import (
	"fmt"
	"sort"
)

func Wb17() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(sort.SearchInts(a, 5))
}
