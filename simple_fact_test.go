package main

import (
	"math/big"
	"testing"
)

func TestSimpleFact(t *testing.T) {
	tt := []struct {
		Up       int
		Expected *big.Int
	}{
		{
			5, big.NewInt(120),
		},
		{
			10, big.NewInt(3628800),
		},
	}
	for _, tc := range tt {
		f := simpleFact(big.NewInt(int64(tc.Up)))
		if f.Cmp(tc.Expected) != 0 {
			t.Errorf("Expected: %v Got:%v", tc.Expected, f)
		}
	}

}
