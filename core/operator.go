package core

import (
	"fmt"
	"math/big"
)

// Factorial find a factorial value for given non-negative integer n
func Factorial(n int) *big.Int {
	s := big.NewInt(1)
	for i := 1; i <= n; i++ {
		s.Mul(s, big.NewInt(int64(i)))
	}
	return s
}

// Bigpow power function for big.Float with exponent of integer
func Bigpow(b *big.Float, e int) *big.Float {
	result := new(big.Float).SetInt(big.NewInt(1))
	for i := 0; i < e; i++ {
		fmt.Println("aaa", result.String(), b.String())
		result = result.Mul(result, b)
	}
	return result
}
