package wbawesome1

//Разработать конвейер чисел.
//Даны два канала: в первый пишутся числа (x) из массива,
//во второй — результат операции x*2,
//после чего данные из второго канала должны выводиться в stdout.

import (
	"fmt"
	"sync"
)

func Wb9() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup

	// множетель
	wg.Add(1)
	go func(ch1 chan int, ch2 chan int) {
		for v := range ch1 {
			ch2 <- v * v
		}
		// закрываем канал сообщая, что больше сообщений не будет
		close(ch2)
		wg.Done()
	}(ch1, ch2)
	// вывод
	wg.Add(1)
	go func(ch2 chan int) {
		for v := range ch2 {
			fmt.Printf("%d ", v)
		}
		wg.Done()
	}(ch2)
	// заприсываем в 1 канал числа
	for _, v := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
		ch1 <- v
	}
	// закрываем канал сообщая, что больше сообщений не будет
	close(ch1)
	wg.Wait()
	fmt.Printf("\n")
}
