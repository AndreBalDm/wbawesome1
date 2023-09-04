package wbawesome1

//Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Wb3() {
	wg := sync.WaitGroup{}
	var result int64
	var resultWrong int64

	for _, v := range []int{2, 4, 6, 8, 10} {
		wg.Add(1)
		go func(v int) {
			atomic.AddInt64(&result, int64(v*v))
			resultWrong += int64(v * v)
			wg.Done()
		}(v)
	}
	wg.Wait()
	fmt.Println("result :      ", atomic.LoadInt64(&result))
	fmt.Println("resultWrong : ", resultWrong)
}
