package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://codeforces.com/contest/1894/problem/C
func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &b[i])
	}
	k = min(n, k)
	last := n - 1
	for k > 0 {
		if b[last] > n {
			Fprint(out, "NO", "\n")
			return
		}
		last = (last + n - b[last]) % n
		k--
	}
	Fprint(out, "YES", "\n")
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var T int
	for Fscan(in, &T); T > 0; T-- {
		solve(in, out)
	}
	Fprintln(out)
}
