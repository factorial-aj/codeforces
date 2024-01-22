package main

import (
	"bufio"
	. "fmt"
	"os"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	var s string
	Fscan(in, &s)
	countB := 0
	for _, ch := range s {
		if ch == 'B' {
			countB++
		}
	}
	if countB == k {
		Fprint(out, 0)
		return
	}
	if countB < k {
		Fprint(out, 1, "\n")
		res := 0
		i := 0
		for countB < k {
			if s[i] == 'A' {
				res++
				countB++
				i++
			} else {
				res++
				i++
			}
		}		
		Fprint(out, res, " ", "B")
		return
	} else {
		Fprint(out, 1, "\n")
		res := 0
		i := 0
		for countB > k {
			if s[i] == 'B' {
				res++
				countB--
				i++
			} else {
				res++
				i++
			}
		}
		Fprint(out, res, " ", "A")
		return
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
