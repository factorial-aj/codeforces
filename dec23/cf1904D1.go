package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	o := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
		o[i] = i
	}
	b := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &b[i])
	}
	sort.Slice(o, func(i, j int) bool {
		return a[o[i]] < a[o[j]]
	})
	for _, i := range o {
		for j := i; j >= 0; j-- {
			if a[j] > a[i] || b[j] < a[i] {
				break
			}
			a[j] = a[i]
		}
		for j := i; j < n; j++ {
			if a[j] > a[i] || b[j] < a[i] {
				break
			}
			a[j] = a[i]
		}
	}
	for i:=0; i<n; i++ {
		if a[i] != b[i] {
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
