package main

import "math/big"

func simpleFact(i *big.Int) *big.Int {
	if i.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(1)
	}
	b := big.NewInt(0).Sub(i, big.NewInt(1))
	c := big.NewInt(0).Mul(i, simpleFact(b))
	return c
}
