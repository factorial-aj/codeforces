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
		a[i] = 1<<30 - 1
	}
	M := make([][]int, n)
	for i := 0; i < n; i++ {
		M[i] = make([]int, n)
		for j := 0; j < n; j++ {
			Fscan(in, &M[i][j])
			if i != j {
				a[i] &= M[i][j]
				a[j] &= M[i][j]
			}
		}
	}
	ok := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && (a[i]|a[j]) != M[i][j] {
				ok = false
			}
		}
	}
	if !ok {
		Fprint(out, "NO")
		return
	}
	Fprint(out, "YES", "\n")
	for i := 0; i < n; i++ {
		Fprint(out, a[i], " ")
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
