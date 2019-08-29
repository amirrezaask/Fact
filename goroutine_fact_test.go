package main

import (
	"math/big"
	"testing"
)

func TestInternalFactFunction(t *testing.T) {
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

func TestPublicFact(t *testing.T) {
	tt := []struct {
		target   int
		ngr      int
		Expected *big.Int
	}{
		{
			5, 1, big.NewInt(120),
		},
		{
			5, 5, big.NewInt(120),
		},
		{
			10, 1, big.NewInt(3628800),
		},
		{
			10, 2, big.NewInt(3628800),
		},
		{
			10, 5, big.NewInt(3628800),
		},
		{
			10, 10, big.NewInt(3628800),
		},
	}
	for _, tc := range tt {
		f := Fact(tc.target, tc.ngr)
		if f.Cmp(tc.Expected) != 0 {
			t.Errorf("Expected: %v Got:%v", tc.Expected, f)
		}
	}
}
