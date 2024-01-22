package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, f, a, b int64
	Fscan(in, &n, &f, &a, &b)
	m := make([]int64, n+1)
	var i int64 = 1
	for i = 1; i < n+1; i++ {
		Fscan(in, &m[i])
	}
	var t int64 = 0
	for i = 1; i < n+1; i++ {
		if (m[i]-m[i-1]) * a >= b {
			t += b
		} else {
			t += (m[i] - m[i-1]) * a
		}
		if t >= f {
			Fprint(out, "NO")
			return
		}
	}
	Fprint(out, "YES")
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

func lowerBound(nums []int, target int) int {
	return sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
}
