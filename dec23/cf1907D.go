package main

import (
	"bufio"
	. "fmt"
	"os"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	Fscan(in, &n)
	seg := make([][]int, n)
	start, end := 0, 0
	for i := 0; i < n; i++ {
		seg[i] = make([]int, 2)
		Fscan(in, &seg[i][0], &seg[i][1])
		end = max(end, seg[i][1])
	}
	feasible := func(k int) bool {
		l, r := 0, 0
		for i := 0; i < n; i++ {
			l = max(l-k, seg[i][0])
			r = min(r+k, seg[i][1])
			if l > r {
				return false
			}
		}
		return true
	}
	res := 0
	for start <= end {
		mid := start + (end-start)/2
		if feasible(mid) {
			res = mid
			end = mid - 1
		} else {
			start = mid + 1
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
