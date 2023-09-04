package wbawesome1

//Разработать программу, которая перемножает, делит, складывает, вычитает
//две числовых переменных a,b, значение которых > 2^20.

import (
	"fmt"
	"math/big"
)

func Wb22() {
	x := 2 << 22 // первое число
	y := 2 << 30 // второе число

	{
		fmt.Println("easy")
		a := float64(x)
		b := float64(y)

		fmt.Println("+ ", a+b)
		fmt.Println("- ", a-b)
		fmt.Println("* ", a*b)
		fmt.Println("/ ", a/b)
	}
	{
		fmt.Println("math/big")
		a := big.NewFloat(float64(x))
		b := big.NewFloat(float64(y))

		fmt.Println("+ ", new(big.Float).Add(a, b))
		fmt.Println("- ", new(big.Float).Sub(a, b))
		fmt.Println("* ", new(big.Float).Mul(a, b))
		fmt.Println("/ ", new(big.Float).Quo(a, b))
	}
}
