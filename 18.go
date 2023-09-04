package wbawesome1

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

import (
	"fmt"
	"sync"
)

type counter struct {
	i int
	sync.Mutex
}

func (s *counter) show() {
	s.Lock()
	fmt.Printf("%d ", s.i)
	s.i++
	s.Unlock()
}

func Wb18() {
	wg := sync.WaitGroup{}
	item := &counter{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				item.show()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("\n")
}
