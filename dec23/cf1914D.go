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
	for i:=0; i<n; i++ {
		Fscan(in, &a[i][0])
		a[i][1] = i
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] > a[j][0]
	})
	b := make([][2]int, n)
	for i:=0; i<n; i++ {
		Fscan(in, &b[i][0])
		b[i][1] = i
	}
	sort.Slice(b, func(i, j int) bool {
		return b[i][0] > b[j][0]
	})
	c := make([][2]int, n)
	for i:=0; i<n; i++ {
		Fscan(in, &c[i][0])
		c[i][1] = i
	}
	sort.Slice(c, func(i, j int) bool {
		return c[i][0] > c[j][0]
	})
	res := 0
	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			for k:=0; k<3; k++ {
				if a[i][1] != b[j][1] && b[j][1] != c[k][1] && c[k][1] != a[i][1] {
					res = max(res, a[i][0]+b[j][0]+c[k][0])
				}
			}
		}
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

func lowerBound(nums []int, target int) int {
	return sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
}