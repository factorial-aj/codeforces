package main

import (
	"bufio"
	. "fmt"
	"os"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	Fscan(in, &n)
	var s string
	Fscan(in, &s)
	res := 0
	curr := 0
	for i := 0; i < n; i++ {
		if s[i] == '.' {
			curr++
		} else {
			res += curr
			curr = 0
		}
		if curr == 3 {
			Fprint(out, 2)
			return
		}
	}
	res += curr
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
