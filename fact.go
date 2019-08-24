package main

import (
	"fmt"
	"math/big"
	"sync"
	"time"
)

var buff chan *big.Int
var wg sync.WaitGroup

func supervisor(i, until, done int64) {
	f := fact(big.NewInt(i), big.NewInt(until))
	// log.Printf("Fact of %v till %v is %v\n", until, i, f)

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

	buff = make(chan *big.Int, ngr)
	result := make(chan *big.Int)

	for i := 0; i < ngr; i++ {
		// log.Printf("UP:%v DOWN:%v", (target/f.NumberOfGoroutines)*(i+1), (target/f.NumberOfGoroutines)*i)
		func(i int) {
			// log.Printf("Goroutine #%v", i)
			wg.Add(1)
			go supervisor(int64((target/ngr)*(i+1)), int64((target/ngr)*i), int64(target))
		}(i)
	}
	go func() {

		// log.Println("waiting for facts to be ready")
		wg.Wait()
		close(buff)
		var val = big.NewInt(1)

		for b := range buff {

			temp := big.NewInt(1)
			val = temp.Mul(val, b)
			// fmt.Printf("val %v\n", val)
		}
		result <- val
	}()
	return <-result
}

func main() {
	s := time.Now()
	_ = Fact(500000, 1000)
	fmt.Println(time.Since(s))
	// fmt.Println(res)
}
