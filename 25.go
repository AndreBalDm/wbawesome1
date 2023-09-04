package wbawesome1

//Реализовать собственную функцию sleep.

import (
	"fmt"
	"time"
)

func Sleep(i int) {
	<-time.After(time.Second * time.Duration(i))
}

func Wb25() {
	fmt.Println("start")
	Sleep(2)
	fmt.Println("stop")
}
