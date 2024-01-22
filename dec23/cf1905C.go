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
	var s string
	Fscan(in, &s)
	subseq := []int{}
	for i := 0; i < n; i++ {
		for len(subseq) > 0 && s[subseq[len(subseq)-1]] < s[i] {
			subseq = subseq[:len(subseq)-1]
		}
		subseq = append(subseq, i)
	}
	j := 0
	for j < len(subseq) && s[subseq[0]] == s[subseq[j]] {
		j++
	}
	res := len(subseq) - j
	sBytes := []byte(s)
	for i:=0; i<len(subseq)/2; i++ {
		sBytes[subseq[i]], sBytes[subseq[len(subseq)-1-i]] = sBytes[subseq[len(subseq)-1-i]], sBytes[subseq[i]]
	}
	for i:=1; i<n; i++ {
		if sBytes[i] < sBytes[i-1] {
			res = -1
			break
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
