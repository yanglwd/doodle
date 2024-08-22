package leetcode

import "math"

// LeetCode 279.完全平方数
type NumberSquares struct {
}

func (ns *NumberSquares) Calc(number int) int {
	f := make([]int, number+1)
	for i := 1; i <= number; i++ {
		res := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			res = min(res, f[i-j*j])
		}
		f[i] = res + 1
	}
	return f[number]
}

func (ns *NumberSquares) min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
