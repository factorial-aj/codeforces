package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var s1, s2, s3 string
	Fscan(in, &s1)
	Fscan(in, &s2)
	Fscan(in, &s3)
	mp := map[byte]bool{
		'A': true,
		'B': true,
		'C': true,
	}
	for i := 0; i < 3; i++ {
		if _, ok := mp[s1[i]]; ok {
			delete(mp, s1[i])
		}
	}
	for key, _ := range mp {
		Fprint(out, string(key))
		return
	}
	mp = map[byte]bool{
		'A': true,
		'B': true,
		'C': true,
	}
	for i := 0; i < 3; i++ {
		if _, ok := mp[s2[i]]; ok {
			delete(mp, s2[i])
		}
	}
	for key, _ := range mp {
		Fprint(out, string(key))
		return
	}
	mp = map[byte]bool{
		'A': true,
		'B': true,
		'C': true,
	}
	for i := 0; i < 3; i++ {
		if _, ok := mp[s3[i]]; ok {
			delete(mp, s3[i])
		}
	}
	for key, _ := range mp {
		Fprint(out, string(key))
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
