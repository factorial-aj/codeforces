package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
		prefix[i+1] = prefix[i] + a[i]
	}
	b := make([]int, n)
	mx := make([]int, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &b[i])
		mx[i+1] = max(mx[i], b[i])
	}
	res := 0
	for i := 0; i < min(n, k); i++ {
		res = max(res, prefix[i+1]+(k-i-1)*mx[i+1])
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

func lowerBound(nums []int, target int) int {
	return sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
}
