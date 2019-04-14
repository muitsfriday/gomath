package main

import (
	"fmt"

	"github.com/muitsfriday/gomath/combi"
	"github.com/muitsfriday/gomath/core"
	"github.com/muitsfriday/gomath/prob"
)

func main() {
	fac := core.NewFac(10)
	val := core.Factorial(7)
	fmt.Println(fac.String())
	fmt.Println(val.String())
	fmt.Println(combi.Combination(10, 5))
	fmt.Println(combi.Permutation(10, 1))
	fmt.Println(combi.Permutation(10, 1))
	fmt.Println(prob.Bin(10, float64(0.5))(1))
	fmt.Println(prob.Bin(10, float64(0.5))(3))
	fmt.Println(prob.Bin(10, float64(0.5))(5))
	fmt.Println(prob.Bin(10, float64(0.5))(7))
	fmt.Println(prob.Pois(float64(1.0 / 2.0))(4))

}
