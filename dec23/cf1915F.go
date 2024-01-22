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
		Fscan(in, &a[i][0], &a[i][1])
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = a[i][1]
	}
	res := 0
	// count of smaller nums after self
	var mergeSort func(nums []int) []int
	mergeSort = func(nums []int) []int {
		if len(nums) < 2 {
			return nums
		}
		mid := len(nums)/2
		left := mergeSort(nums[:mid])
		right := mergeSort(nums[mid:])
		sortedNums := []int{}
		i, j := 0, 0
		for i < len(left) || j < len(right) {
			if j >= len(right) || (i < len(left) && left[i] <= right[j]) {
				sortedNums = append(sortedNums, left[i])
				res += j
				i++
			} else {
				sortedNums = append(sortedNums, right[j])
				j++
			}
		}
		return sortedNums
	}
	mergeSort(b)
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
