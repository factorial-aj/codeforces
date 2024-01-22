package main

import (
	"bufio"
	. "fmt"
	"os"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + a[i]
	}
	res := 0
	for k:=1; k<=n/2; k++ {
		if n%k != 0 {
			continue
		}
		mx, mn := 0, 1<<63-1
		for i:=0; i+k<=n; i+=k {
			curr := prefix[i+k] - prefix[i]
			mx = max(mx, curr)
			mn = min(mn, curr)
		}
		res = max(res, mx-mn)
	}
	Fprint(out, res)
}

func sum(a []int) int {
	res := 0
	for _, v := range a {
		res += v
	}
	return res
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

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
