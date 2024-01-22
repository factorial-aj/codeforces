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
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	if k >= 3 {
		Fprint(out, 0)
		return
	}
	sort.Ints(a)
	mn := a[0]
	for i := 0; i < n-1; i++ {
		mn = min(mn, a[i+1]-a[i])
	}
	if k == 1 {
		Fprint(out, mn)
		return
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			v := a[j] - a[i]
			p := lowerBound(a, v)
			if p < n {
				mn = min(mn, a[p]-v)
			}
			if p > 0 {
				mn = min(mn, v-a[p-1])
			}
		}
	}
	Fprint(out, mn)
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

func lowerBound(arr []int, target int) int {
	return sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})
}
