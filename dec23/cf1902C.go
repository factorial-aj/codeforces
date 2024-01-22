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
	set := make(map[int]bool)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
		set[a[i]] = true
	}
	if n == 1 {
		Fprint(out, 1)
		return
	}
	sort.Ints(a)
	x := a[1] - a[0]
	for i := 2; i < n; i++ {
		x = gcd(x, a[i]-a[i-1])
	}
	res := 0
	for i := 0; i < n; i++ {
		res += (a[n-1] - a[i]) / x
	}
	k := -1
	for set[a[n-1]+x*k] {
		k--
	}
	res -= k
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
