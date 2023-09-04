package wbawesome1

// Поменять местами два числа без создания временной переменной.

import "fmt"

func Wb13() {
	a := 1
	b := 2
	a += b
	b = a - b
	a = a - b
	//a, b = b, a
	fmt.Println(a, b)
}
