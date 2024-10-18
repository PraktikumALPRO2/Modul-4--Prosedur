package main

import (
	"fmt"
	"math/big"
)

func factorial(n int) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	result := big.NewInt(1)
	for i := 1; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

func permutation(n, r int) *big.Int {
	return factorial(n).Div(factorial(n), factorial(n-r))
}

func combination(n, r int) *big.Int {
	return factorial(n).Div(factorial(n), factorial(r).Mul(factorial(r), factorial(n-r)))
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		p1 := permutation(a, c)
		c1 := combination(a, c)
		p2 := permutation(b, d)
		c2 := combination(b, d)

		fmt.Printf("P(%d,%d) %s\n", a, c, p1.String())
		fmt.Printf("C(%d,%d) %s\n", a, c, c1.String())
		fmt.Printf("P(%d,%d) %s\n", b, d, p2.String())
		fmt.Printf("C(%d,%d) %s\n", b, d, c2.String())
	}
}
