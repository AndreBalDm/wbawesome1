package wbawesome1

// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
//и выведет их квадраты в stdout.

import (
	"fmt"
	"sync"
)

func Wb2() {
	wg := sync.WaitGroup{}
	for _, v := range []int{2, 4, 6, 8, 10} {
		wg.Add(1)
		go func(v int) {
			fmt.Printf("%d ", v*v)
			wg.Done()
		}(v)
	}
	wg.Wait()
	fmt.Printf("\n")
}
