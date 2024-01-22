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
	b := make([]int, n)
	for i:=0; i<n; i++ {
		Fscan(in, &b[i])
	}
	lMax, rMin := 0, 1 << 31 - 1
	res := 0
	for i := 0; i < n; i++ {
		if a[i] > b[i] {
			a[i], b[i] = b[i], a[i]
		}
		lMax = max(lMax, a[i])
		rMin = min(rMin, b[i])
		res += (b[i] - a[i])
	}
	if lMax > rMin {
		res += 2 * (lMax - rMin)
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
