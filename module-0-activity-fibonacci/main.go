package main

import (
	"fmt"
	"math/big"
)

func f(n int) *big.Int {
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(99), nil)
	dp := make(map[int]*big.Int, n)

	return fib(n, dp)

}

func fib(n int, dp map[int]*big.Int) *big.Int {

	if n == 0 {
		return big.NewInt(0)
	} else if n == 1 {
		return big.NewInt(1)
	}

	var ans big.Int
	if _, ok := dp[n]; !ok {
		ans.Add(fib(n-1, dp), fib(n-2, dp))
		dp[n] = &ans
	}
	return dp[n]
}

func main() {
	n := 8181
	fmt.Println(f(n))
}
