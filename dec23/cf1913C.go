package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var m int
	Fscan(in, &m)
	cnt := make([]int, 30)
	for ; m > 0; m-- {
		var t, v int
		Fscan(in, &t, &v)
		if t == 1 {
			cnt[v]++
		} else {
			for i := 29; i >= 0; i-- {
				take := min(v>>i, cnt[i])
				v -= take << i
			}
			if v == 0 {
				Fprint(out, "YES", "\n")
			} else {
				Fprint(out, "NO", "\n")
			}
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	//var T int
	//for Fscan(in, &T); T > 0; T-- {
	solve(in, out)
	//Fprint(out, "\n")
	//}
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
