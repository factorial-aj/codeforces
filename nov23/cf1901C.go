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
	mn, mx := 1<<31-1, 0
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
		mx = max(mx, a[i])
		mn = min(mn, a[i])
	}
	if mn == mx {
		Fprint(out, 0)
		return
	}
	res := make([]int, 0)
	for mn != mx {
		t := mn & 1
		res = append(res, t)
		mn += t
		mn >>= 1
		mx += t
		mx >>= 1
	}
	Fprint(out, len(res))
	if len(res) <= n {
		Fprint(out, "\n")
		for i := 0; i < len(res); i++ {
			Fprint(out, res[i], " ")
		}
	}
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
