//nolint
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func fillData(data []int64, rng int64) {
	max := big.NewInt(rng)
	l := len(data)

	for i := 0; i < l; i++ {
		n,_ := rand.Int(rand.Reader, max)
		data[i] = n.Int64()
	}
}

func haveRepeats(data []int64, n int) bool {
	l := len(data) - n
outer:
	for i := 0; i <= l; i++ {
		for j := 1; j < n; j++ {
			if data[i] != data[i+j] {
				continue outer
			}
		}

		return true
	}

	return false
}

const (
	seqLen = 1000
	rng = 10
	repeats = 5
	tries = 50000
)

func main() {
	_ = getExpectation(10, 1000, 5)

	n := 0
	data := make([]int64, seqLen)

	for i := 0; i < tries; i++ {
		fillData(data, rng)
		if haveRepeats(data, repeats) {
			n++
		}

		if i%100 == 99 {
			fmt.Print(".")
		}

		if i%10000 == 9999 {
			fmt.Printf("  %d\n", n)
		}
	}
	fmt.Printf("\nTotal: %d; Have %d repeats: %d; [%f]\n", tries, repeats, n, float64(n)/float64(tries))
}

func getExpectation(rng int, seqLen int, repeats int) *big.Float {
	r := big.NewFloat(float64(rng))

	sp := big.NewFloat(0)
	sp.Copy(r)

	for i := 2; i < repeats; i++ {
		sp.Mul(sp, r)
	}

	sp.Quo(big.NewFloat(1), sp)
	fmt.Printf("'Sequnce of %d numbers is equal' probability: %s\n", repeats, sp.String())

	sp.Sub(big.NewFloat(1), sp)
	fmt.Printf("'Sequnce of %d numbers isn't equal probability': %s\n", repeats, sp.String())


	tp := big.NewFloat(0)
	tp.Copy(sp)

	for i := repeats; i < seqLen; i++ {
		tp.Mul(tp,sp)
	}

	sub := tp.Sub(big.NewFloat(1), tp)
	fmt.Printf("At least one sequence of %d equal numbers in the row of %d numbers: %s\n\n", repeats, seqLen, sub.String())

	return sub
}
