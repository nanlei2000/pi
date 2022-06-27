package main

import (
	"math/big"
)

func Calc(n uint64) *big.Int {
	t := n + 10

	b := new(big.Int)
	b.Exp(big.NewInt(10), big.NewInt(int64(t)), nil)

	x1 := new(big.Int)
	x1.Mul(b, big.NewInt(4))
	FloorDiv(x1, x1, big.NewInt(5))

	x2 := new(big.Int)
	FloorDiv(x2, b, big.NewInt(-239))

	s := new(big.Int)
	s.Add(x1, x2)

	n *= 2

	for i := uint64(3); i < n; i += 2 {
		FloorDiv(x1, x1, big.NewInt(-25))
		FloorDiv(x2, x2, big.NewInt(-57121))
		x := new(big.Int)

		x.Add(x1, x2)
		FloorDiv(x, x, big.NewInt(int64(i)))
		s.Add(s, x)
	}

	s.Mul(s, big.NewInt(4))
	FloorDiv(s, s, new(big.Int).Exp(big.NewInt(10), big.NewInt(10), nil))

	return s
}

func FloorDiv(z, x, y *big.Int) {
	m := new(big.Int)
	z.DivMod(x, y, m)
	if z.Cmp(big.NewInt(0)) == -1 && m.Cmp(big.NewInt(0)) == +1 {
		z.Add(z, big.NewInt(-1))
	}
}
