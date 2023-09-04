package wbawesome1

// Реализовать конкурентную запись данных в map.

import (
	"fmt"
	"sync"
)

func Wb7() {
	m := make(map[int]int)
	var wg sync.WaitGroup
	var mx sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			mx.Lock()
			m[1] = 1
			mx.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(m)
}
