package main

import (
	"fmt"
	"math/big"
)

var (
	n25    = big.NewInt(-25)
	n57121 = big.NewInt(-57121)
	p10    = big.NewInt(10)
	zero   = big.NewInt(0)
)

// https://zhuanlan.zhihu.com/p/114320417, see part 3
func Calc(n uint64) *big.Int {
	t := n + 10

	b := new(big.Int)
	b.Exp(p10, big.NewInt(int64(t)), nil)

	x1 := new(big.Int)
	x1.Mul(b, big.NewInt(4))
	FloorDiv(x1, x1, big.NewInt(5))

	x2 := new(big.Int)
	FloorDiv(x2, b, big.NewInt(-239))

	s := new(big.Int)
	s.Add(x1, x2)

	n *= 2
	x := new(big.Int)
	for i := uint64(3); i < n; i += 2 {
		FloorDiv(x1, x1, n25)
		FloorDiv(x2, x2, n57121)

		x.Add(x1, x2)
		FloorDiv(x, x, big.NewInt(int64(i)))
		s.Add(s, x)
	}

	s.Mul(s, big.NewInt(4))
	FloorDiv(s, s, new(big.Int).Exp(p10, p10, nil))

	return s
}

func FloorDiv(z, x, y *big.Int) {
	m := new(big.Int)
	z.DivMod(x, y, m)
	if z.Cmp(zero) == -1 && m.Cmp(zero) == +1 {
		z.Add(z, big.NewInt(-1))
	}
}

func Fmt(pi *big.Int) string {
	piStr := pi.String()
	return fmt.Sprintf("%s.%s", string(piStr[0]), string(piStr[1:]))
}
