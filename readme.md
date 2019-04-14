# gomath


## Core

```
import "github.com/muitsfriday/gomath/core"
```

### Factorial

Find the factorial number.

```
core.Factorial(10)
```


## Combinatoric

```
import "github.com/muitsfriday/gomath/combi"
```

### nCr

Find the combination number of C(n, r)

```
combi.Combination(10, 5)
```

### nPr

Find the permutation number of P(n, r)


```
combi.Permutation(10, 1)
```


## Probability

Import the package
```
import	"github.com/muitsfriday/gomath/prob"
```

### Binomial Distibution

Create binomial distribution function with `prob.Bin(n int, p float64)`.

Example 

A fair coin tossed 10 times. Find the probability of 5 coins become head. 

```
pmf := prob.Bin(10, float64(0.5))
pval := pmf(5) 
```

### Poisson Distibution

Create binomial poisson function with `prob.Pois(l float64)`.

Example

The mean of accident in a specified crossroad is 2 times per year.
Find the probability that accident will occur 3 time this year.

```
pmf := prob.Pois(float64(2.0))
pval := pmf(3) 
```