package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	Fscan(in, &n)
	a := make([][2]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i][0])
		a[i][1] = i
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})
	res := make([]int, n)
	j, curr := 0, 0
	for i := 0; i < n; i++ {
		if j == i {
			j++
			curr += a[i][0]
		}
		for j < n && curr >= a[j][0] {
			curr += a[j][0]
			j++
		}
		res[a[i][1]] = j - 1
	}
	for i := 0; i < n; i++ {
		Fprint(out, res[i], " ")
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
