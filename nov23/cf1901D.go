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
	suffix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = max(prefix[i], a[i]+n-i-1)
	}
	for i := n - 1; i >= 0; i-- {
		suffix[i] = max(suffix[i+1], a[i]+i)
	}
	res := 1<<31 - 1
	for i := 0; i < n; i++ {
		res = min(res, max(prefix[i], a[i], suffix[i+1]))
	}
	Fprint(out, res)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	//var T int
	//for Fscan(in, &T); T > 0; T-- {
	solve(in, out)
	Fprint(out, "\n")
	//}
	Fprintln(out)
}

func min(x ...int) int {
	res := x[0]
	for _, v := range x {
		if v < res {
			res = v
		}
	}
	return res
}

func max(x ...int) int {
	res := x[0]
	for _, v := range x {
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
