package wbawesome1

//Разработать программу, которая переворачивает подаваемую на ход строку
//(например: «главрыба — абырвалг»). Символы могут быть unicode.

import "fmt"

func Wb19() {
	str1 := "главрыба — абырвалг"
	runes := []rune(str1)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	str2 := string(runes)

	fmt.Printf("Было: %s\n", str1)
	fmt.Printf("Стало: %s\n", str2)
}
