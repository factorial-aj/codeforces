package main

import (
	"bufio"
	. "fmt"
	"os"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	Fscan(in, &n)
	res := 1
	for n > 0 {
		d := n % 10
		n /= 10
		mul := 0
		for i := 0; i <= d; i++ {
			for j := 0; j <= d; j++ {
				if d-i-j >= 0 {
					mul++
				}
			}
		}
		res *= mul
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
