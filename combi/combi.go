package combi

import (
	"math/big"

	"github.com/muitsfriday/gomath/core"
)

// Combination find combination value nCr = n! / r!(n-r)!
func Combination(n int, r int) *big.Int {
	s := big.NewInt(1)
	nFac := core.Factorial(n)
	rFac := core.Factorial(r)
	nrFac := core.Factorial(n - r)
	denominator := big.NewInt(1)

	return s.Div(nFac, denominator.Mul(rFac, nrFac))
}

// Permutation find permutation value nPr = n! / (n-r)!
func Permutation(n int, r int) *big.Int {
	nFac := core.Factorial(n)
	nrFac := core.Factorial(n - r)
	return big.NewInt(1).Div(nFac, nrFac)
}
