package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var s string
	Fscan(in, &s)
	n := len(s)
	zeroes := 0
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			zeroes++
		}
	}
	ones := n - zeroes
	i := 0
	for i = 0; i < n; i++ {
		if s[i] == '0' {
			if ones == 0 {
				break
			}
			ones--
		} else {
			if zeroes == 0 {
				break
			}
			zeroes--
		}
	}
	Fprint(out, n-i)
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
