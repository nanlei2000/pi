package main

import (
	"fmt"
	"math/big"
)

func Calc(n int64) {
	t := n + 10

	b := new(big.Int)
	b.Exp(big.NewInt(10), big.NewInt(t), nil)

	x1 := new(big.Int)
	x1.Mul(b, big.NewInt(4))
	x1.Div(x1, big.NewInt(5))

	x2 := new(big.Int)
	x2.Div(b, big.NewInt(-239))

	// fmt.Println("x1", x1, "x2", x2)
	s := new(big.Int)
	s.Add(x1, x2)

	n *= 2

	for i := int64(3); i < n; i += 2 {
		x1.Div(x1, big.NewInt(-25))
		x2.Div(x2, big.NewInt(-57121))
		// fmt.Println("x1", x1, "x2", x2)
		x1.Add(x1, x2)
		x1.Div(x1, big.NewInt(i))
		s.Add(s, x1)
	}

	s.Mul(s, big.NewInt(4))
	s.Div(s, new(big.Int).Exp(big.NewInt(10), big.NewInt(10), nil))

	fmt.Println(s)
}
