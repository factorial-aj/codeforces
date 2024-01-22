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
	if n == 1 {
		Fprint(out, 0)
		return
	}
	last := a[n-1]
	res := 0
	for i := n - 2; i >= 0; i-- {
		if a[i] <= last {
			last = a[i]
			continue
		}
		k := a[i]/last
		if a[i] % last != 0 {
			k++
		}
		last = a[i]/k
		res += k-1
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
