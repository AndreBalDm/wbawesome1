package wbawesome1

//Дана последовательность температурных колебаний:
//-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов.
//Последовательность в подмножноствах не важна.
//Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

import (
	"fmt"
)

func Wb10() {
	result := make(map[int][]float32)
	t := []float32{-9, 9, -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	for _, c := range t {
		var index int
		cInt := int(c)

		if c < 0 {
			index = cInt - cInt%10 - 10
		} else {
			index = cInt - cInt%10
		}
		fmt.Println(index, cInt, c)
		result[index] = append(result[index], c)

	}
	fmt.Printf("%+v\n", result)
}
