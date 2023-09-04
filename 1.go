package wbawesome1

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

import "fmt"

type Human1 struct {
	age int
}

type Action struct {
	Human1
}

func (h *Human1) SetAge(age int) {
	h.age = age
}

func Wb1() {
	action := &Action{}
	action.SetAge(12)
	fmt.Printf("%+v\n", action)
}
