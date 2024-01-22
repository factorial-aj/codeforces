package main

import (
	"bufio"
	. "fmt"
	"os"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, P, l, t int
	Fscan(in, &n, &P, &l, &t)
	c := (n+6)/7
	feasible := func(k int) bool {
		return k*l + min(c, 2*k) * t >= P
	}
	start, end := 0, n
	res := 0
	for start <= end {
		mid := start + (end - start) / 2
		if feasible(mid) {
			res = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	Fprint(out, n-res)
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
