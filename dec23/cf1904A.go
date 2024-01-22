package main

import (
	"bufio"
	. "fmt"
	"os"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var a, b int
	Fscan(in, &a, &b)
	var xk, yk int
	Fscan(in, &xk, &yk)
	var xq, yq int
	Fscan(in, &xq, &yq)
	dirs := [][]int{
		{a, b},
		{a, -b},
		{-a, b},
		{-a, -b},
		{b, a},
		{b, -a},
		{-b, a},
		{-b, -a},
	}
	kings := make(map[[2]int]bool)
	queens := make(map[[2]int]bool)
	for _, dir := range dirs {
		kings[[2]int{xk+dir[0], yk+dir[1]}] = true
		queens[[2]int{xq+dir[0], yq+dir[1]}] = true
	}
	res := 0
	for king, _ := range kings {
		for queen, _ := range queens {
			if king == queen {
				res++
			}
		}
	}
	Fprint(out, res)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var T int
	for Fscan(in, &T); T > 0; T-- {
		solve(in, out)
		Fprint(out, "\n")
	}
	Fprintln(out)
}

func min(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if v < res {
			res = v
		}
	}
	return res
}

func max(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if v > res {
			res = v
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
