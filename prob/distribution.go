package prob

import (
	"fmt"
	"math"
	"math/big"

	"github.com/muitsfriday/gomath/combi"
	"github.com/muitsfriday/gomath/core"
)

// Bin get binomial distribution function
// n is the number of berneulli trials
// p is the probability that single binomial trial will success
// returns a function (x int) -> float
func Bin(n int, p float64) func(x int) float64 {
	return func(x int) float64 {
		if x < 0 || x > n {
			return float64(0)
		}
		ppow := big.NewFloat(math.Pow(p, float64(x)))
		qpow := big.NewFloat(math.Pow(1-p, float64(n-x)))
		c := new(big.Float).SetInt(combi.Combination(n, x))

		pterm := big.NewFloat(float64(1.0)).Mul(ppow, qpow)
		prob, _ := big.NewFloat(1).Mul(c, pterm).Float64()
		return prob
	}
}

// Pois get poisson distribution function
func Pois(l float64) func(x int) float64 {
	fmt.Println("bbbb", l)

	return func(x int) float64 {
		lbig := new(big.Float).SetFloat64(l)
		e := math.Exp(1)

		term1 := core.Bigpow(lbig, x)
		term2 := new(big.Float).SetFloat64(math.Pow(e, -1*l))
		fmt.Println(term1, term2)
		numer := new(big.Float).Mul(term1, term2)
		denot := new(big.Float).SetInt(core.Factorial(x))

		result, _ := new(big.Float).Quo(numer, denot).Float64()
		return result
	}
}
