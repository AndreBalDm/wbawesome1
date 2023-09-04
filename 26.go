package wbawesome1

//Разработать программу, которая проверяет, что все символы в строке уникальные
//(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
//Например: abcd — true, abCdefAaf — false,	aabcd — false

import (
	"fmt"
	"strings"
)

func check(s string) bool {
	fr := make(map[rune]struct{})
	s = strings.ToLower(s)

	for _, r := range []rune(s) {
		//r = unicode.ToLower(r)
		if _, ok := fr[r]; ok {
			return false
		}
		fr[r] = struct{}{}
	}
	return true
}

func Wb26() {
	fmt.Println(check("abCdefAaf"))
	fmt.Println(check("abcd"))
	fmt.Println(check("aabcd"))
}
