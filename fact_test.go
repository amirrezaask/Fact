package main

import (
	"math/big"
	"testing"
)

func TestFact(t *testing.T) {
	tt := []struct {
		Up       int
		Down     int
		Expected *big.Int
	}{
		{
			5, 1, big.NewInt(120),
		},
		{
			10, 1, big.NewInt(3628800),
		},
	}
	for _, tc := range tt {
		f := fact(big.NewInt(int64(tc.Up)), big.NewInt(int64(tc.Down)))
		if f.Cmp(tc.Expected) != 0 {
			t.Errorf("Expected: %v Got:%v", tc.Expected, f)
		}
	}

}
