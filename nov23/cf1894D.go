package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

// https://codeforces.com/contest/1894/problem/D
func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	b := make([]int, m)
	for i := 0; i < m; i++ {
		Fscan(in, &b[i])
	}
	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})
	l, r := 0, 0
	for l < n || r < m {
		if r >= m || (l < n && a[l] > b[r]) {
			Fprint(out, a[l], " ")
			l++
		} else {
			Fprint(out, b[r], " ")
			r++
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
