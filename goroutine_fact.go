package main

import (
	"fmt"
	"math/big"
	"sync"
	"time"
)

func supervisor(buff chan *big.Int, wg *sync.WaitGroup, i, until int64) {
	f := fact(big.NewInt(i), big.NewInt(until))

	buff <- f

	wg.Done()
}

func fact(i, until *big.Int) *big.Int {
	if i.Cmp(until) == 0 {
		return big.NewInt(1)
	}
	b := big.NewInt(0)
	b.Sub(i, big.NewInt(1))
	c := big.NewInt(0)
	c.Mul(i, fact(b, until))

	return c
}

//Fact calculates Fact
func Fact(target int, ngr int) *big.Int {

	buff := make(chan *big.Int, ngr)
	result := make(chan *big.Int)
	wg := &sync.WaitGroup{}

	for i := 0; i < ngr; i++ {
		func(i int) {
			wg.Add(1)
			go supervisor(buff, wg, int64((target/ngr)*(i+1)), int64((target/ngr)*i))
		}(i)
	}
	go func() {

		wg.Wait()
		close(buff)
		var val = big.NewInt(1)

		for b := range buff {

			temp := big.NewInt(1)
			val = temp.Mul(val, b)
		}
		result <- val
	}()
	return <-result
}

func main() {
	s := time.Now()
	_ = Fact(1000000, 5000)
	fmt.Println(time.Since(s))
}
